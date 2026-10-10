// The writing example shows the writing components of the shadcn/ui
// expansions together in a page of prose: the heading with its anchor,
// the typewriter and the block quote. Clicking the anchor copies
// "uix#getting-started" to the clipboard, and the typewriter types its
// line once.
//
//	go run ./examples/writing
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/writing/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Writing",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.WritingView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
