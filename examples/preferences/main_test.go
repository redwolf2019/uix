package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/preferences/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.PreferencesView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading, the items, the
// language select and the danger zone.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Preferences",
		"Account", "Notifications", "Privacy",
		"Language", "English",
		"Delete account…",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSelects opens the native select, chooses a language, and the
// trigger shows it.
func TestSelects(t *testing.T) {
	tt := newTester(t)
	// The native menu opens as the trigger goes down.
	r, ok := tt.Find("English")
	if !ok {
		t.Fatal("no language trigger")
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
	if err := tt.ChooseMenuItem("日本語"); err != nil {
		t.Fatal(err)
	}
	if view.State.Language != "ja" {
		t.Fatalf("language = %q, want ja", view.State.Language)
	}
	if !tt.HasText("日本語") {
		t.Fatalf("the trigger does not show the choice in %q", tt.Texts())
	}
}

// TestDeletes opens the alert dialog from the danger item, confirms it,
// and reports the deletion.
func TestDeletes(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Delete account…"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Delete this account?") {
		t.Fatalf("the confirm dialog is missing in %q", tt.Texts())
	}
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	if view.State.ConfirmOpen {
		t.Fatal("the dialog stayed open after the confirm")
	}
	if !tt.HasText("Account deleted") {
		t.Fatalf("missing the deletion message in %q", tt.Texts())
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Preferences") {
		t.Fatal("preferences missing under dark mode")
	}
}
