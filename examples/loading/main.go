// The loading example shows the remaining components of the shadcn/ui
// expansions together in a checkout: the magic back button, the
// interactive stepper with the infinite-scroll catalog and the loading
// button, and the responsive modal that confirms the order. The catalog
// loads more products as it nears its end; the Pay button turns into a
// spinner for a moment and then opens the order as a modal.
//
//	go run ./examples/loading
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/loading/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Loading",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.LoadingView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
