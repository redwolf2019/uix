// Package webview tests the component: the web view itself lives in the
// window and shows no page headless, so the tests draw the layout it
// sits in and check that the component fills its room without breaking
// it.
package webview

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestRenders draws the component with a text beside it and checks the
// layout still comes through and the frame paints.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			ui.Text(c, "Toolbar")
			WebView(c, Props{}).Grow(1)
		})
	}, 200, 150)
	if !tt.HasText("Toolbar") {
		t.Fatalf("texts %q", tt.Texts())
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 || img.Bounds().Dy() == 0 {
		t.Fatal("no frame rendered")
	}
}
