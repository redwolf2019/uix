// The fields example shows the field components of the shadcn/ui
// expansions together in a form: the floating label inputs, the
// auto-size text area, the date-time picker and the multiple selector.
// The fields edit the form in place, and Save counts the saves.
//
//	go run ./examples/fields
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/fields/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Fields",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.FieldsView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
