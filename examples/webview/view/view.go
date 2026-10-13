// Package view builds the webview example: a window of MyGo native UI
// that shows the webview component of uix — a toolbar of native controls
// over a web page. The main package makes the web view and sets Web; the
// tests and the snapshot command draw the view headless, where there is
// no page and the toolbar alone shows.
package view

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/webview"
)

// Width and Height are the size of the example window.
const Width, Height = 760, 520

// HomeURL is the page the web view opens.
const HomeURL = "https://example.com"

// Web is the web view of the window; nil while the window has not made
// it, so the headless tests and the snapshot draw the toolbar alone.
var Web *mygo.WebView

// Reset restores the initial state of the example.
func Reset() { Web = nil }

// Title is the title of the example window.
func Title() string { return "WebView" }

// WebviewView draws the example: a toolbar with the title, the page's
// URL and a reload button, and the web view growing below it. The reload
// button reloads the page; the native chrome around it paints over it.
func WebviewView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).FillWidth().MinWidth(0).AlignItems(ui.Center).Gap(t.Space(3)).
			Padding(t.Space(3), t.Space(4)).Children(func() {
			ui.Text(c, "WebView").Bold()
			ui.Text(c, HomeURL).Grow(1).MinWidth(0).SingleLine().TextColor(t.TextMuted)
			button.Button(c, button.Props{
				Label:   "Reload",
				Variant: button.Ghost,
				OnClick: func() {
					if Web != nil {
						Web.Page().Reload()
					}
				},
			})
		})
		if Web != nil {
			webview.WebView(c, webview.Props{Web: Web}).Grow(1)
		}
	})
}
