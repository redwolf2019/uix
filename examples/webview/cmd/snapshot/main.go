// Command snapshot renders the webview example headless into PNG files,
// so the toolbar of the example can be seen without opening a window:
//
//	go run ./examples/webview/cmd/snapshot
//
// The page itself shows only in a window; the snapshot shows the native
// chrome the window draws over it.
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/webview/view"
)

func main() {
	for _, dark := range []bool{false, true} {
		view.Reset()
		tt := ui.NewTester(view.WebviewView, view.Width, view.Height)
		tt.SetDark(dark)
		name := "light.png"
		if dark {
			name = "dark.png"
		}
		save(tt, name)
	}
}

// save renders one more frame of tt and writes it to name.
func save(tt *ui.Tester, name string) {
	tt.Frame()
	img := tt.Image()
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
