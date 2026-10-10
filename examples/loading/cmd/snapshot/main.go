// Command snapshot renders the loading example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/loading/cmd/snapshot
package main

import (
	"image/png"
	"math"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/loading/view"
)

func main() {
	// The plain shot draws the checkout as it opens, on the catalog
	// step.
	view.Reset()
	save(ui.NewTester(view.LoadingView, view.Width, view.Height), "loading.png")

	// The paying shot draws the Pay step with the button spinning.
	view.Reset()
	view.State.Step = 2
	view.State.Paying = true
	save(ui.NewTester(view.LoadingView, view.Width, view.Height), "loading-paying.png")

	// The order shot draws the order modal over the checkout, and the
	// catalog loaded to its end.
	view.Reset()
	view.State.Step = 2
	view.State.Open = true
	tt := ui.NewTester(view.LoadingView, view.Width, view.Height)
	view.State.Catalog.Scroll.Y = math.MaxFloat32
	tt.Frame()
	save(tt, "loading-order.png")

	// The dark shot draws the checkout under the dark appearance.
	view.Reset()
	tt = ui.NewTester(view.LoadingView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "loading-dark.png")
}

// save writes the frame of the tester into a PNG file.
func save(tt *ui.Tester, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, tt.Image()); err != nil {
		panic(err)
	}
}
