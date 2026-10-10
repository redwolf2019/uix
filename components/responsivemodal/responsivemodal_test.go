package responsivemodal

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			ResponsiveModal(c, p)
		})
	}
}

// TestCentered shows the modal centered on a wide window.
func TestCentered(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Confirm"}), 800, 400)
	r, ok := tt.Find("Confirm")
	if !ok {
		t.Fatal("the modal does not show")
	}
	if r.Y > 300 {
		t.Fatalf("the modal did not center: %+v", r)
	}
}

// TestSheet rises from the bottom on a narrow window.
func TestSheet(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Confirm"}), 360, 480)
	r, ok := tt.Find("Confirm")
	if !ok {
		t.Fatal("the modal does not show")
	}
	if r.Y < 300 {
		t.Fatalf("the sheet did not rise from the bottom: %+v", r)
	}
}

// TestClosed draws nothing while the modal is closed.
func TestClosed(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Confirm"}), 800, 400)
	if tt.HasText("Confirm") {
		t.Fatal("closed, it shows its content")
	}
}

// TestDarkMode draws the modal under the dark appearance too.
func TestDarkMode(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Confirm"}), 800, 400)
	tt.SetDark(true)
	if !tt.HasText("Confirm") {
		t.Fatal("modal missing under dark mode")
	}
}
