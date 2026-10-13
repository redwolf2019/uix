// The webview example shows the webview component: a window of MyGo
// native UI with a toolbar of native controls over a web page, the page
// of a web view made with mygo.Window.NewWebView. The reload button
// reloads the page, and the native chrome of the window draws over it.
//
//	go run ./examples/webview
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/webview/view"
)

func main() {
	view.Reset()
	mygo.App.WhenReady(func() {
		win := mygo.NewWindow(mygo.WindowOptions{
			Title:   view.Title(),
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.WebviewView),
		})
		v, err := win.NewWebView(mygo.WebViewOptions{URL: view.HomeURL})
		if err != nil {
			panic(err)
		}
		view.Web = v
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
