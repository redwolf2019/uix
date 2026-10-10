package message

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/bubble"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Message(c, p)
		})
	}
}

// TestRenders draws the avatar, the header, the bubble and the footer
// of a turn.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Name:   "Ada Lovelace",
		Header: "Ada · now",
		Text:   "Hello there",
		Footer: "Read",
	}), 360, 140)
	for _, want := range []string{"AL", "Ada · now", "Hello there", "Read"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestAlignEnd puts the turn at the end of the row: the bubble text
// starts past the middle.
func TestAlignEnd(t *testing.T) {
	tt := ui.NewTester(frame(Props{Align: bubble.End, Name: "You", Text: "Done"}), 360, 120)
	r, ok := tt.Find("Done")
	if !ok {
		t.Fatal("the bubble text is missing")
	}
	if r.X < 180 {
		t.Fatalf("an end-aligned turn starts at x=%g, want past the middle", r.X)
	}
}

// TestEmptyAvatar keeps the room of the avatar for a turn without a
// sender, as a group's earlier turns do: the bubble starts where it
// does with the avatar of the first turn.
func TestEmptyAvatar(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Another turn"}), 360, 120)
	r, ok := tt.Find("Another turn")
	if !ok {
		t.Fatal("the bubble text is missing")
	}
	tt2 := ui.NewTester(frame(Props{Name: "Ada", Text: "With avatar"}), 360, 120)
	r2, ok := tt2.Find("With avatar")
	if !ok {
		t.Fatal("the avatar turn is missing")
	}
	if r.X != r2.X {
		t.Fatalf("a turn without a sender starts at x=%g, with one at x=%g; the room differs", r.X, r2.X)
	}
}

// TestGroup stacks the turns of one sender.
func TestGroup(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			MessageGroup(c, func() {
				Message(c, Props{Name: "Ada", Text: "First"})
				Message(c, Props{Text: "Second"})
			})
		})
	}, 360, 180)
	for _, want := range []string{"First", "Second"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestDarkMode draws a turn under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Name: "Ada", Text: "Hello"}), 360, 120)
	tt.SetDark(true)
	if !tt.HasText("Hello") {
		t.Fatal("message missing under dark mode")
	}
}
