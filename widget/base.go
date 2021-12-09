package widget

import (
	"embed"
	"fmt"

	"github.com/asaskevich/EventBus"

	core "github.com/i2y/oden/core"
)

type Widget interface {
	core.Widget
	Detach()
	Update()
	SetSize(w, h int) Widget
	Width() int
	SetWidth(w int) Widget
	Height() int
	SetHeight(h int) Widget
	SizePolicy() SizePolicy
	SetSizePolicy(sp SizePolicy) Widget
	FixedWidth(n int) Widget
	FixedRatioWidth(n int) Widget
	FixedHeight(n int) Widget
	FixedRatioHeight(n int) Widget
	FixedSize(w, h int) Widget
	SizeStyle() string
	SetSizeStyle(s string)
	TextStyle() *TextStyle
	OtherStyle() *OtherStyle
	SetTextStyle(style *TextStyle) Widget
	Align(ta TextAlign) Widget
	VerticalAlign(va VerticalAlign) Widget
	FgColor(fg *Color) Widget
	BgColor(bg *Color) Widget
	BorderColor(c *Color) Widget
	BorderRadius(r int) Widget
	FontSize(size *FontSize) Widget
	Padding(n int) Widget
	Margin(n int) Widget
	OnClick(func(ev core.Event)) Widget
	OnChange(func(ev core.Event)) Widget
}

type SizePolicy int

const (
	Fixed SizePolicy = iota
	FixedWidth
	FixedRatioWidth
	FixedHeight
	FixedRatioHeight
	Expanding
)

type TextAlign int

const (
	Center TextAlign = iota
	Start
	End
	Justify
)

func (ta TextAlign) String() string {
	switch ta {
	case Center:
		return "center"
	case Start:
		return "start"
	case End:
		return "end"
	case Justify:
		return "justify"
	}
	return ""
}

type VerticalAlign int

const (
	Middle VerticalAlign = iota
	Baseline
	Top
	Bottom
	Sub
	TextTop
)

func (va VerticalAlign) String() string {
	switch va {
	case Middle:
		return "middle"
	case Baseline:
		return "baseline"
	case Top:
		return "top"
	case Bottom:
		return "buttom"
	case Sub:
		return "sub"
	case TextTop:
		return "text-top"
	}
	return ""
}

type Color struct {
	name      string
	swatchNum int
}

func (c *Color) String() string {
	return fmt.Sprintf("--sl-color-%s-%d", c.name, c.swatchNum)
}

func NewColor(name string, swatchNum int) *Color {
	return &Color{
		name:      name,
		swatchNum: swatchNum,
	}
}

var (
	PrimaryColor = NewColor("primary", 500)
	SuccessColor = NewColor("success", 500)
	WarningColor = NewColor("warning", 500)
	DangerColor  = NewColor("danger", 500)
	NeutralColor = NewColor("neutral", 500)
	Black        = NewColor("neutral", 0)
	White        = NewColor("neutral", 1000)
	Gray         = NewColor("gray", 500)
	Red          = NewColor("red", 500)
	Orange       = NewColor("orange", 500)
	Amber        = NewColor("amber", 500)
	Yellow       = NewColor("yellow", 500)
	Lime         = NewColor("lime", 500)
	Green        = NewColor("green", 500)
	Emerald      = NewColor("emerald", 500)
	Teal         = NewColor("teal", 500)
	Cyan         = NewColor("cyan", 500)
	Sky          = NewColor("sky", 500)
	Blue         = NewColor("blue", 500)
	Indigo       = NewColor("indigo", 500)
	Violet       = NewColor("violet", 500)
	Purple       = NewColor("purple", 500)
	Fuchsia      = NewColor("fuchsia", 500)
	Pink         = NewColor("pink", 500)
	Rose         = NewColor("rose", 500)
)

type FontSize struct {
	name string
}

func (fs *FontSize) String() string {
	return fmt.Sprintf("--sl-font-size-%s", fs.name)
}

func NewFontSize(name string) *FontSize {
	return &FontSize{
		name: name,
	}
}

var (
	TwoXSmall   = NewFontSize("2x-small")
	XSmall      = NewFontSize("x-small")
	Small       = NewFontSize("small")
	Medium      = NewFontSize("medium")
	Large       = NewFontSize("large")
	XLarge      = NewFontSize("x-large")
	TwoXLarge   = NewFontSize("2x-large")
	ThreeXLarge = NewFontSize("3x-large")
	FourXLarge  = NewFontSize("4x-large")
)

type TextStyle struct {
	align         TextAlign
	verticalAlign VerticalAlign
	fgColor       *Color
	bgColor       *Color
	borderColor   *Color
	borderRadius  int
	fontSize      *FontSize
	padding       int
	margin        int
}

func (s *TextStyle) String() string {
	style := fmt.Sprintf("text-align: %s; vertical-align: %s;", s.align, s.verticalAlign)
	if s.fgColor != nil {
		style += fmt.Sprintf(" color: var(%s);", s.fgColor)
	}
	if s.bgColor != nil {
		style += fmt.Sprintf(" background-color: var(%s);", s.bgColor)
	}
	if s.borderColor != nil {
		style += fmt.Sprintf(" border-color: var(%s);", s.borderColor)
	}
	style += fmt.Sprintf(" border-radius: %dpx;", s.borderRadius)
	if s.fontSize != nil {
		style += fmt.Sprintf(" font-size: var(%s);", s.fontSize)
	}
	style += fmt.Sprintf(" padding: %dpx", s.padding)
	return style
}

type OtherStyle struct {
	margin int
}

func (s *OtherStyle) String() string {
	style := fmt.Sprintf("padding: %dpx;", s.margin)
	return style
}

type Base struct {
	id           int
	attached     bool
	app          *core.App
	widget       Widget
	actualWidget Widget
	width        int
	height       int
	sizePolicy   SizePolicy
	sizeStyle    string
	textStyle    *TextStyle
	otherStyle   *OtherStyle
}

func NewBase() Base {
	return Base{
		id:         core.NewWidgetID(),
		attached:   false,
		sizePolicy: Expanding,
		textStyle: &TextStyle{
			align:         Center,
			verticalAlign: Middle,
		},
		otherStyle: &OtherStyle{
			margin: 0,
		},
	}
}

func (b *Base) ID() int {
	return b.id
}

func (b *Base) IDStr() string {
	return fmt.Sprintf("oden-%d", b.id)
}

func (b *Base) View() string {
	return ""
}

func (b *Base) Attach(a *core.App) {
	b.attached = true
	b.app = a
}

func (b *Base) Detach() {
	b.attached = false
}

func (b *Base) SetWidget(w Widget) {
	b.widget = w
}

func (b *Base) Update() {
	if b.attached {
		b.app.PostUpdate(b.widget)
	}
}

func (b *Base) OnClick(handler func(ev core.Event)) Widget {
	core.AddEventHandler(b.widget, "click", handler)
	return b.widget
}

func (b *Base) OnChange(handler func(ev core.Event)) Widget {
	core.AddEventHandler(b.widget, "sl-change", handler)
	return b.widget
}

type EventPublisher interface {
	AddListener(w Widget)
	Notify()
}

type Model struct {
	bus EventBus.Bus
}

func NewModel() Model {
	return Model{
		bus: EventBus.New(),
	}
}

func (m *Model) AddListener(b Widget) {
	m.bus.Subscribe("update", b.Update)
}

func (m *Model) Notify() {
	m.bus.Publish("update")
}

func (b *Base) SizePolicy() SizePolicy {
	return b.sizePolicy
}

func (b *Base) SetSizePolicy(sp SizePolicy) Widget {
	b.sizePolicy = sp
	return b.widget
}

func (b *Base) SetSize(width, height int) Widget {
	b.width = width
	b.height = height
	return b.widget
}

func (b *Base) Width() int {
	return b.width
}

func (b *Base) SetWidth(width int) Widget {
	b.width = width
	return b.widget
}

func (b *Base) FixedWidth(width int) Widget {
	return b.SetSizePolicy(FixedWidth).SetWidth(width)
}

func (b *Base) FixedRatioWidth(width int) Widget {
	return b.SetSizePolicy(FixedRatioWidth).SetWidth(width)
}

func (b *Base) Height() int {
	return b.height
}

func (b *Base) SetHeight(height int) Widget {
	b.height = height
	return b.widget
}

func (b *Base) FixedHeight(height int) Widget {
	return b.SetSizePolicy(FixedHeight).SetHeight(height)
}

func (b *Base) FixedRatioHeight(height int) Widget {
	return b.SetSizePolicy(FixedHeight).SetHeight(height)
}

func (b *Base) FixedSize(w, h int) Widget {
	return b.SetSizePolicy(Fixed).SetSize(w, h)
}

func (b *Base) SetSizeStyle(s string) {
	b.sizeStyle = s
}

func (b *Base) SizeStyle() string {
	return b.sizeStyle
}

func (b *Base) Build() Widget {
	return b.widget
}

func (b *Base) TextStyle() *TextStyle {
	return b.textStyle
}

func (b *Base) SetTextStyle(style *TextStyle) Widget {
	b.textStyle = style
	return b.widget
}

func (b *Base) Align(ta TextAlign) Widget {
	b.textStyle.align = ta
	return b.widget
}

func (b *Base) VerticalAlign(va VerticalAlign) Widget {
	b.textStyle.verticalAlign = va
	return b.widget
}

func (b *Base) FgColor(fg *Color) Widget {
	b.textStyle.fgColor = fg
	return b.widget
}

func (b *Base) BgColor(bg *Color) Widget {
	b.textStyle.bgColor = bg
	return b.widget
}

func (b *Base) BorderColor(c *Color) Widget {
	b.textStyle.borderColor = c
	return b.widget
}

func (b *Base) BorderRadius(r int) Widget {
	b.textStyle.borderRadius = r
	return b.widget
}

func (b *Base) FontSize(size *FontSize) Widget {
	b.textStyle.fontSize = size
	return b.widget
}

func (b *Base) Padding(n int) Widget {
	b.textStyle.padding = n
	return b.widget
}

func (b *Base) OtherStyle() *OtherStyle {
	return b.otherStyle
}

func (b *Base) Margin(n int) Widget {
	b.otherStyle.margin = n
	return b.widget
}

//go:embed assets
var assets embed.FS

func init() {
	core.SetHeadElements(`
  <link rel="stylesheet" media="(prefers-color-scheme:light)"
    href="assets/node_modules/@shoelace-style/shoelace/dist/themes/light.css">
  <link rel="stylesheet" media="(prefers-color-scheme:dark)"
    href="assets/node_modules/@shoelace-style/shoelace/dist/themes/dark.css"
    onload="document.documentElement.classList.add('sl-theme-dark');">
  <script type="module" src="assets/node_modules/@shoelace-style/shoelace/dist/shoelace.js"></script>

  <link rel="stylesheet" href="assets/style.css">`)
	core.SetTargetEvents([]core.TargetEvent{
		{Name: "click", PropName: ""},
		{Name: "sl-change", PropName: "value"}},
	)
	core.MountAssets(assets)
}
