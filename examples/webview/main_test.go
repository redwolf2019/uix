package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/webview/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.Reset() }

// TestRenders draws the webview view headless and checks that the
// toolbar came through. Headless there is no page, only the chrome of
// the example, which is what the window draws over the page.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.WebviewView, view.Width, view.Height)
	for _, s := range []string{"WebView", view.HomeURL, "Reload"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestReload clicks the reload button; with no web view made headless,
// the click must still be safe.
func TestReload(t *testing.T) {
	reset()
	tt := ui.NewTester(view.WebviewView, view.Width, view.Height)
	if err := tt.Click("Reload"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("WebView") {
		t.Fatal("toolbar gone after reload")
	}
}

// TestDarkMode draws the webview view under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.WebviewView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("WebView") {
		t.Fatal("dark view missing the title")
	}
}
