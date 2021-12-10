// +build windows

package oden

import (
	"fmt"
	"syscall"

	"github.com/jchv/go-webview2"
)

type WebView2 struct {
	wv webview2.WebView
}

func (wv *WebView2) Open(title string, port, width, height int) {
	wv.wv.SetTitle(title)
	wv.wv.SetSize(width, height, webview2.HintNone)
	wv.wv.Navigate(fmt.Sprintf("http://localhost:%d", port))
	defer wv.wv.Destroy()
	wv.wv.Run()
	return
}

func detectWebview2() *WebView2 {
	dll := syscall.MustLoadDLL("user32")
	if proc, err := dll.FindProc("SetProcessDpiAwarenessContext"); err == nil {
		aware := -4
		proc.Call(uintptr(aware))
	}
	wv := webview2.New(false)
	if wv == nil {
		return nil
	}
	return &WebView2{wv: wv}
}
