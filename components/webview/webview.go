// Package webview shows a web page inside the native UI, as a component
// for the WebView of MyGo: a page the window made with
// mygo.Window.NewWebView lives in, and this component places it in the
// view. The page works as a window's does: it loads the app's frontend
// or a URL, and can call Go methods, hear events and evaluate scripts.
//
//	docs, err := win.NewWebView(mygo.WebViewOptions{URL: "/docs"})
//	// ...
//	webview.WebView(c, webview.Props{Web: docs}).Grow(1)
package webview

import "github.com/egoist/mygo/ui"

// Props describes the web view to draw.
type Props struct {
	// Web is the web view to show, made with mygo.Window.NewWebView of
	// the window drawing this view. A window without a web view must not
	// draw one: the page keeps its state while the element does not
	// show, until Destroy or until its window closes.
	Web ui.NativeWebView
}

// WebView draws the web view of props and returns it, so a view can
// chain more calls on it. A web view has no size of its own: it
// stretches across its container, and takes the room along it with
// Grow or a size. Its children, popovers, menus, dialogs, tooltips and
// toasts paint over the page, which its background and border frame.
func WebView(c *ui.Context, p Props) ui.Element {
	return ui.WebView(c, p.Web)
}
