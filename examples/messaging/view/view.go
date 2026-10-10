// Package view builds the messaging example: a window of MyGo native UI
// that shows the message family of uix together — the message scroller,
// the message rows with their bubbles and avatars, and the composer
// below. A send appends a turn of each side, and the scroller follows
// the new messages while the reader is at the end; scrolling back shows
// the button back to the newest. The main package shows it in a window;
// the tests and the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/bubble"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/message"
	"github.com/smallnest/uix/components/messagescroller"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 500

// Turn is one message of the conversation.
type Turn struct {
	// Align is the side of the row the turn sits on.
	Align bubble.Align
	// Name is the sender, whose initials fill the avatar.
	Name string
	// Time is the line above the bubble.
	Time string
	// Text is the message.
	Text string
	// Variant is the look of the bubble.
	Variant bubble.Variant
}

// State is the state of the example; the controls edit it in place.
var State = MessagingState{}

// MessagingState holds the example: the turns, where the scroller is,
// and the draft of the composer.
type MessagingState struct {
	// Turns are the messages of the conversation, in order.
	Turns []Turn
	// Scroll is where the transcript scroller is.
	Scroll messagescroller.State
	// Draft is the text of the composer.
	Draft string
	// Sent is how many sends the session made, for the tests.
	Sent int
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = MessagingState{Turns: []Turn{
		{Name: "Ada", Time: "Ada · yesterday", Text: "Welcome! I can help you set up the project.", Variant: bubble.Secondary},
		{Align: bubble.End, Name: "You", Time: "You · now", Text: "Let's begin.", Variant: bubble.Default},
		{Name: "Ada", Time: "Ada · now", Text: "Great. I sent you the outline in the docs.", Variant: bubble.Secondary},
		{Align: bubble.End, Name: "You", Time: "You · now", Text: "Perfect, I read it.", Variant: bubble.Default},
		{Name: "Ada", Time: "Ada · now", Text: "Shall we walk through the first step together?", Variant: bubble.Secondary},
		{Align: bubble.End, Name: "You", Time: "You · now", Text: "Yes, go ahead.", Variant: bubble.Default},
		{Name: "Ada", Time: "Ada · now", Text: "It is the part where the app asks for your name.", Variant: bubble.Secondary},
	}}
}

// MessagingView draws the example: the heading, the transcript in its
// scroller, and the composer at the bottom.
func MessagingView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Children(func() {
		ui.Column(c).FillWidth().Padding(24, 24, 12, 24).Gap(1).Children(func() {
			ui.Text(c, "Messaging").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "The message family of uix.").TextColor(t.TextMuted)
		})
		messagescroller.MessageScroller(c, messagescroller.Props{
			State:    &State.Scroll,
			NewLabel: "New messages",
			Children: func() {
				ui.Column(c).FillWidth().Gap(t.Space(2)).Padding(24, 12).Children(func() {
					for _, turn := range State.Turns {
						message.Message(c, message.Props{
							Align:   turn.Align,
							Name:    turn.Name,
							Header:  turn.Time,
							Text:    turn.Text,
							Variant: turn.Variant,
						})
					}
				})
			},
		})
		composer(c, t)
	})
}

// composer draws the field and the Send button at the bottom of the
// window; a send appends a turn of each side and clears the field.
func composer(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Gap(t.Space(2)).Padding(24, 16).Children(func() {
		ui.Column(c).Grow(1).Children(func() {
			in := input.Input(c, input.Props{Value: &State.Draft, Placeholder: "Message…"})
			if in.Submitted() {
				send()
			}
		})
		button.Button(c, button.Props{Label: "Send", OnClick: send})
	})
}

// send appends the draft as a turn of the current user, and the reply
// of the assistant after it, and clears the draft.
func send() {
	if State.Draft == "" {
		return
	}
	text := State.Draft
	State.Draft = ""
	State.Turns = append(State.Turns, Turn{Align: bubble.End, Name: "You", Time: "You · now", Text: text, Variant: bubble.Default})
	State.Turns = append(State.Turns, Turn{Name: "Ada", Time: "Ada · now", Text: "Got it. Anything else?", Variant: bubble.Secondary})
	State.Sent++
}
