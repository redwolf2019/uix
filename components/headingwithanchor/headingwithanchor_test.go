package headingwithanchor

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Heading(c, p)
		})
	}
}

// TestRenders draws the heading text.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Installation"}), 320, 100)
	if !tt.HasText("Installation") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestAnchor copies the anchor to the clipboard for a click on the
// button.
func TestAnchor(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Installation", Anchor: "installation"}), 320, 100)
	// The button has no text; find it by its rect past the heading.
	r, ok := tt.Find("Installation")
	if !ok {
		t.Fatal("no heading")
	}
	tt.ClickAt(r.X+r.W+20, r.Y+r.H/2)
	if got := tt.Clipboard(); got != "uix#installation" {
		t.Fatalf("clipboard %q, want uix#installation", got)
	}
}

// TestAnchorCallback runs the callback of the anchor instead.
func TestAnchorCallback(t *testing.T) {
	copied := ""
	tt := ui.NewTester(frame(Props{
		Text:     "Installation",
		Anchor:   "installation",
		OnAnchor: func() { copied = "yes" },
	}), 320, 100)
	r, ok := tt.Find("Installation")
	if !ok {
		t.Fatal("no heading")
	}
	tt.ClickAt(r.X+r.W+20, r.Y+r.H/2)
	if copied != "yes" {
		t.Fatal("the anchor callback did not run")
	}
}

// TestNoAnchor draws the heading without the button.
func TestNoAnchor(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Installation"}), 320, 100)
	r, _ := tt.Find("Installation")
	if r.W > 300 {
		t.Fatalf("the heading took a wide box %+v, want the text alone", r)
	}
}

// TestDarkMode draws the heading under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Installation", Anchor: "installation"}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("Installation") {
		t.Fatal("heading missing under dark mode")
	}
}
