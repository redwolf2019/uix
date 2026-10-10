// The messaging example shows the message family of uix together: a
// conversation in the message scroller, whose rows are the message
// turns with their bubbles and avatars, and the composer below. A send
// appends a turn of each side, and the scroller follows the new
// messages while the reader is at the end; scrolling back shows the
// button back to the newest.
//
//	go run ./examples/messaging
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/messaging/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Messaging",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.MessagingView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
