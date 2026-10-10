// Command snapshot renders the preferences example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/preferences/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/preferences/view"
)

func main() {
	// The plain shot draws the screen as it opens.
	view.Reset()
	save(ui.NewTester(view.PreferencesView, view.Width, view.Height), "preferences.png")

	// The menu shot opens the native select of the language field.
	view.Reset()
	tt := ui.NewTester(view.PreferencesView, view.Width, view.Height)
	r, ok := tt.Find("English")
	if !ok {
		panic("no language trigger")
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
	save(tt, "preferences-menu.png")

	// The confirm shot opens the alert dialog of the danger zone.
	view.Reset()
	view.State.ConfirmOpen = true
	save(ui.NewTester(view.PreferencesView, view.Width, view.Height), "preferences-confirm.png")

	// The dark shot draws the screen under the dark appearance.
	view.Reset()
	tt = ui.NewTester(view.PreferencesView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "preferences-dark.png")
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
