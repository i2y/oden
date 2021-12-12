//go:build !windows
// +build !windows

package oden

type WebView2 struct{}

func (wv *WebView2) open(title string, port, width, height int) {
}

func detectWebview2() *WebView2 {
	return nil
}
