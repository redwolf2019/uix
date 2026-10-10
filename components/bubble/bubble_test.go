package bubble

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Bubble(c, p)
		})
	}
}

// TestRenders draws the text of a bubble.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hello there"}), 320, 120)
	if !tt.HasText("Hello there") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestAlignEnd puts the bubble at the end of its row: its text starts
// past the middle of the row.
func TestAlignEnd(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Done", Align: End}), 320, 120)
	r, ok := tt.Find("Done")
	if !ok {
		t.Fatal("the bubble text is missing")
	}
	if r.X < 160 {
		t.Fatalf("an end-aligned bubble starts at x=%g, want past the middle", r.X)
	}
}

// TestAlignStart leaves the bubble at the start of its row.
func TestAlignStart(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hello", Align: Start}), 320, 120)
	r, ok := tt.Find("Hello")
	if !ok {
		t.Fatal("the bubble text is missing")
	}
	if r.X > 160 {
		t.Fatalf("a start-aligned bubble starts at x=%g, want before the middle", r.X)
	}
}

// TestVariants draws every look of a bubble.
func TestVariants(t *testing.T) {
	for _, v := range []Variant{Default, Secondary, Muted, Tinted, Outline, Ghost, Destructive} {
		tt := ui.NewTester(frame(Props{Text: "A message", Variant: v}), 320, 120)
		if !tt.HasText("A message") {
			t.Fatalf("variant %d lost the text: %q", v, tt.Texts())
		}
	}
}

// TestDarkMode draws a bubble under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hello"}), 320, 120)
	tt.SetDark(true)
	if !tt.HasText("Hello") {
		t.Fatal("bubble missing under dark mode")
	}
}
