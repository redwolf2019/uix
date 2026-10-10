package alertdialog

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			AlertDialog(c, p)
		})
	}
}

// TestShows draws the dialog: the title, the description and the
// buttons.
func TestShows(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{
		Open:        &open,
		Title:       "Delete this account?",
		Description: "This cannot be undone.",
		Confirm:     "Delete",
	}), 360, 240)
	for _, want := range []string{"Delete this account?", "This cannot be undone.", "Delete", "Cancel"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestClosed draws nothing while the dialog is closed.
func TestClosed(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Delete this account?", Confirm: "Delete"}), 360, 240)
	if tt.HasText("Delete this account?") {
		t.Fatalf("a closed dialog drew its title: %q", tt.Texts())
	}
}

// TestConfirm runs OnConfirm for the action button, which closes the
// dialog.
func TestConfirm(t *testing.T) {
	open := true
	confirmed := false
	tt := ui.NewTester(frame(Props{
		Open: &open, Title: "Delete?", Confirm: "Delete",
		OnConfirm: func() { confirmed = true },
	}), 360, 240)
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	if !confirmed {
		t.Fatal("OnConfirm did not run")
	}
	if open {
		t.Fatal("the dialog stayed open after the action")
	}
}

// TestCancel runs OnCancel for the cancel button, which closes the
// dialog.
func TestCancel(t *testing.T) {
	open := true
	cancelled := false
	tt := ui.NewTester(frame(Props{
		Open: &open, Title: "Delete?", Confirm: "Delete",
		OnCancel: func() { cancelled = true },
	}), 360, 240)
	if err := tt.Click("Cancel"); err != nil {
		t.Fatal(err)
	}
	if !cancelled {
		t.Fatal("OnCancel did not run")
	}
	if open {
		t.Fatal("the dialog stayed open after the cancel")
	}
}

// TestMedia draws the icon of the dialog above the title.
func TestMedia(t *testing.T) {
	open := true
	icon := ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 6h18"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"/></svg>`))
	tt := ui.NewTester(frame(Props{
		Open: &open, Title: "Delete?", Confirm: "Delete", Media: icon,
	}), 360, 260)
	if !tt.HasText("Delete?") {
		t.Fatalf("missing the title in %q", tt.Texts())
	}
}

// TestDarkMode draws the dialog under the dark appearance too.
func TestDarkMode(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Delete?", Confirm: "Delete"}), 360, 240)
	tt.SetDark(true)
	if !tt.HasText("Delete?") {
		t.Fatal("dialog missing under dark mode")
	}
}
