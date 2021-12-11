package oden

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/asaskevich/EventBus"
	"golang.org/x/net/websocket"
)

type WidgetID int

func (id WidgetID) String() string {
	return fmt.Sprintf("oden-%d", id)
}

var widgetID WidgetID

var idMutex sync.Mutex

func NewWidgetID() WidgetID {
	idMutex.Lock()
	defer idMutex.Unlock()
	widgetID++
	return widgetID
}

type Widget interface {
	ID() WidgetID
	View() string
	Attach(app *App)
}

type rawEvent struct {
	Target    string                 `json:"target"`
	EventName string                 `json:"event"`
	Props     map[string]interface{} `json:"props"`
}

func (e *rawEvent) ID() string {
	return fmt.Sprintf("%s.%s", e.Target, e.EventName)
}

type Event interface {
	ID() string
	Target() Widget
	EventName() string
	Props() map[string]interface{}
}

type actualEvent struct {
	target    Widget
	eventName string
	props     map[string]interface{}
}

func (e *actualEvent) ID() string {
	return fmt.Sprintf("%s.%s", e.target.ID(), e.eventName)
}

func (e *actualEvent) Target() Widget {
	return e.target
}

func (e *actualEvent) EventName() string {
	return e.eventName
}

func (e *actualEvent) Props() map[string]interface{} {
	return e.props
}

type App struct {
	ctx      context.Context
	name     string
	cancel   context.CancelFunc
	server   *http.Server
	listener net.Listener
	width    int
	height   int
	view     Widget
	widgetID int
	msgs     chan string
}

func NewApp(name string, width, height int, view Widget) *App {
	ctx, cancel := context.WithCancel(context.Background())
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	app := &App{
		name:     name,
		ctx:      ctx,
		cancel:   cancel,
		server:   &http.Server{Addr: listener.Addr().String(), Handler: mux},
		listener: listener,
		width:    width,
		height:   height,
		view:     view,
		widgetID: 0,
		msgs:     make(chan string),
	}

	topWidget := view

	tmpl, err := template.New("index").Parse(indexTmpl)
	if err != nil {
		panic(err)
	}

	assetHandler := http.FileServer(assetsFS)
	mux.Handle("/assets/", assetHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ipAddr := getIPAddr(r)
		if !(ipAddr == "[::1]" || ipAddr == "127.0.0.1" || ipAddr == "localhost" || ipAddr == "::1") {
			return
		}
		err := tmpl.Execute(w, &templateParams{
			Name:         app.name,
			HeadElements: headElements,
			Events:       targetEvents,
			Widget:       topWidget.View(),
			Port:         app.port(),
		})
		if err != nil {
			log.Printf("failed to execute the index template: %v", err)
		}
	})
	mux.Handle("/ws", websocket.Handler(app.handleWebSocket))
	mux.HandleFunc("/turbo.es2017-umd.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(turbo))
	})
	return app
}

func getIPAddr(r *http.Request) string {
	parts := strings.Split(r.RemoteAddr, ":")
	if len(parts) < 2 {
		return ""
	}

	nameParts := parts[:len(parts)-1]
	return strings.Join(nameParts, ":")
}

func (app *App) handleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	u, _ := url.Parse(ws.RemoteAddr().String())
	hostname := u.Hostname()
	if !(hostname == "localhost" || hostname == "127.0.0.1") {
		return
	}

	go func() {
		for {
			var r string
			err := websocket.Message.Receive(ws, &r)
			if err != nil {
				if err == io.EOF {
					app.cancel()
					return
				}
				log.Fatal(err)
			}

			var ev rawEvent
			err = json.Unmarshal([]byte(r), &ev)
			if err != nil {
				log.Fatal(err)
			}

			bus.Publish(ev.ID(), &ev)
		}
	}()

	for {
		msg := <-app.msgs
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			log.Fatal(err)
		}
	}

}

//go:embed index.html
var indexTmpl string

//go:embed turbo.es2017-umd.js
var turbo string

type templateParams struct {
	Name         string
	HeadElements string
	Events       []TargetEvent
	Widget       string
	Port         int
}

func (app *App) serve() {
	app.view.Attach(app)
	app.server.Serve(app.listener)
}

func (app *App) port() int {
	return app.listener.Addr().(*net.TCPAddr).Port
}

func (app *App) Run() {
	go app.serve()
	browser := detectBrowser()
	if browser == nil {
		log.Fatal("any supported browser not found")
		return
	}
	browser.open(app.name, app.port(), app.width, app.height)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
	case <-app.ctx.Done():
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func (app *App) PostUpdate(w Widget) {
	app.msgs <- fmt.Sprintf(
		`<turbo-stream action="replace" target="oden-%d"><template>%s</template></turbo-stream>`,
		w.ID(),
		w.View(),
	)
}

var bus = EventBus.New()

func AddEventHandler(w Widget, event string, handler func(ev Event)) {
	bus.Subscribe(
		fmt.Sprintf("%d.%s", w.ID(), event),
		func(rev *rawEvent) {
			ev := &actualEvent{
				target:    w,
				eventName: rev.EventName,
				props:     rev.Props,
			}
			ev.target = w
			handler(ev)
		},
	)
}

var headElements string

func SetHeadElements(s string) {
	headElements = s
}

type TargetEvent struct {
	Name     string
	PropName string
}

var targetEvents []TargetEvent

func SetTargetEvents(events []TargetEvent) {
	targetEvents = events
}

var assetsFS http.FileSystem

func MountAssets(assets embed.FS) {
	assetsLFS, err := fs.Sub(assets, ".")
	if err != nil {
		panic(err)
	}

	assetsFS = http.FS(assetsLFS)
}
