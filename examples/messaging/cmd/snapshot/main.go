// Command snapshot renders the messaging example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/messaging/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/bubble"
	"github.com/smallnest/uix/examples/messaging/view"
)

func main() {
	// The plain shot draws the conversation as it opens.
	view.Reset()
	save(ui.NewTester(view.MessagingView, view.Width, view.Height), "messaging.png")

	// The sent shot appends a send, so the newest turns show.
	view.Reset()
	view.State.Turns = append(view.State.Turns,
		view.Turn{Align: bubble.End, Name: "You", Time: "You · now", Text: "Good morning", Variant: bubble.Default},
		view.Turn{Name: "Ada", Time: "Ada · now", Text: "Got it. Anything else?", Variant: bubble.Secondary},
	)
	save(ui.NewTester(view.MessagingView, view.Width, view.Height), "messaging-sent.png")

	// The up shot draws the transcript scrolled up, with the button
	// back to the newest over the bottom.
	view.Reset()
	tt := ui.NewTester(view.MessagingView, view.Width, view.Height)
	view.State.Scroll.Offset.Y = 0
	tt.Frame()
	save(tt, "messaging-up.png")

	// The dark shot draws the conversation under the dark appearance.
	view.Reset()
	tt = ui.NewTester(view.MessagingView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "messaging-dark.png")
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
