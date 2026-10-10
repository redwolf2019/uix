package loadingbutton

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			LoadingButton(c, p)
		})
	}
}

// TestRenders draws the label.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Save"}), 320, 100)
	if !tt.HasText("Save") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestClicks runs OnClick while the button is not loading.
func TestClicks(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{Label: "Save", OnClick: func() { clicks++ }}), 320, 100)
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	if clicks != 1 {
		t.Fatalf("clicked %d times, want 1", clicks)
	}
}

// TestLoading keeps the label and stops the clicks while it loads.
func TestLoading(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{
		Label:   "Save",
		Loading: true,
		OnClick: func() { clicks++ },
	}), 320, 100)
	if !tt.HasText("Save") {
		t.Fatalf("texts %q", tt.Texts())
	}
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	if clicks != 0 {
		t.Fatalf("a loading button was clicked %d times", clicks)
	}
}

// TestDarkMode draws the button under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Save", Loading: true}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("Save") {
		t.Fatal("button missing under dark mode")
	}
}
