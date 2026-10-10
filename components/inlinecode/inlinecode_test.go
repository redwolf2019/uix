package inlinecode

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			InlineCode(c, p)
		})
	}
}

// TestRenders draws the code chip.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "uix add button"}), 320, 100)
	if !tt.HasText("uix add button") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestDarkMode draws the chip under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "uix add button"}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("uix add button") {
		t.Fatal("chip missing under dark mode")
	}
}
