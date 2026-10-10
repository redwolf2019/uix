package blockquote

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Blockquote(c, p)
		})
	}
}

// TestRenders draws the quote and its author.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Text:   "The only thing we have to fear is fear itself.",
		Author: "Franklin D. Roosevelt",
	}), 360, 160)
	for _, want := range []string{"The only thing we have to fear is fear itself.", "Franklin D. Roosevelt"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestNoAuthor draws the quote without an attribution.
func TestNoAuthor(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "A quiet thought."}), 360, 140)
	if !tt.HasText("A quiet thought.") {
		t.Fatalf("missing the quote in %q", tt.Texts())
	}
}

// TestDarkMode draws the quote under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "A quiet thought."}), 360, 140)
	tt.SetDark(true)
	if !tt.HasText("A quiet thought.") {
		t.Fatal("blockquote missing under dark mode")
	}
}
