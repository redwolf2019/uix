// Command snapshot renders the fields example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/fields/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/fields/view"
)

func main() {
	// The plain shot draws the form as it opens.
	view.Reset()
	save(ui.NewTester(view.FieldsView, view.Width, view.Height), "fields.png")

	// The filled shot draws the form with the fields filled and a
	// guest chosen.
	view.Reset()
	view.State.Title = "Team lunch"
	view.State.Email = "ada@example.com"
	view.State.Notes = "Agenda:\n- Roadmap\n- Design review\n- Retro"
	view.State.Guests = []string{"Ada", "Grace"}
	tt := ui.NewTester(view.FieldsView, view.Width, view.Height)
	save(tt, "fields-filled.png")

	// The popup shot draws the multi-select's popup open.
	view.Reset()
	tt = ui.NewTester(view.FieldsView, view.Width, view.Height)
	view.State.Open = true
	tt.Frame()
	save(tt, "fields-popup.png")

	// The dark shot draws the form under the dark appearance.
	view.Reset()
	tt = ui.NewTester(view.FieldsView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "fields-dark.png")
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
