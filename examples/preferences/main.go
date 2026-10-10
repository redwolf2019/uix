// The preferences example shows the item, native select and alert
// dialog components of uix together, as a settings screen for an
// account: the item group lists the sections, the native select
// chooses a language, and the destructive item opens the alert dialog
// that confirms deleting the account.
//
//	go run ./examples/preferences
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/preferences/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Preferences",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.PreferencesView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
