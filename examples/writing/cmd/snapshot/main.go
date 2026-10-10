// Command snapshot renders the writing example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/writing/cmd/snapshot
package main

import (
	"image/png"
	"os"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/writing/view"
)

func main() {
	// The plain shot draws the page as it opens, with the typewriter
	// part-way through its line.
	view.Reset()
	tt := ui.NewTester(view.WritingView, view.Width, view.Height)
	time.Sleep(300 * time.Millisecond)
	tt.Frame()
	save(tt, "writing.png")

	// The typed shot draws the typewriter line complete.
	view.Reset()
	tt = ui.NewTester(view.WritingView, view.Width, view.Height)
	time.Sleep(1100 * time.Millisecond)
	tt.Frame()
	save(tt, "writing-typed.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt = ui.NewTester(view.WritingView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "writing-dark.png")
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
