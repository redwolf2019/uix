package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/messaging/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.MessagingView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading and the turns of
// the initial conversation.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Messaging",
		"Welcome! I can help you set up the project.",
		"Let's begin.",
		"Yes, go ahead.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSends types a message and sends it: the turn of each side
// appends, and the scroller stays at the end.
func TestSends(t *testing.T) {
	tt := newTester(t)
	// The composer is at the bottom: the field spans from the left edge
	// to the Send button.
	send, ok := tt.Find("Send")
	if !ok {
		t.Fatal("no Send button")
	}
	tt.ClickAt(60, send.Y+send.H/2)
	tt.Type("Good morning")
	if err := tt.Click("Send"); err != nil {
		t.Fatal(err)
	}
	if view.State.Sent != 1 {
		t.Fatalf("sent %d, want 1", view.State.Sent)
	}
	if !tt.HasText("Good morning") {
		t.Fatalf("the sent message is missing in %q", tt.Texts())
	}
	if !tt.HasText("Got it. Anything else?") {
		t.Fatalf("the reply is missing in %q", tt.Texts())
	}
	if !view.State.Scroll.AtEnd {
		t.Fatal("the scroller left the end after the send")
	}
}

// TestJumps scrolls the transcript up and returns to the newest with
// the button.
func TestJumps(t *testing.T) {
	tt := newTester(t)
	// The reader scrolled up: the next frame shows the jump button.
	view.State.Scroll.Offset.Y = 0
	tt.Frame()
	if !tt.HasText("New messages") {
		t.Fatalf("the jump button is missing in %q", tt.Texts())
	}
	if err := tt.Click("New messages"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Scroll.AtEnd {
		t.Fatal("the jump did not return to the end")
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Messaging") {
		t.Fatal("messaging missing under dark mode")
	}
}
