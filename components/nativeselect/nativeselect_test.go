package nativeselect

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			NativeSelect(c, p)
		})
	}
}

// TestRenders draws the trigger showing the choice.
func TestRenders(t *testing.T) {
	chosen := "en"
	tt := ui.NewTester(frame(Props{Selected: &chosen, Options: []Option{
		{Label: "English", Value: "en"},
		{Label: "简体中文", Value: "zh"},
	}}), 320, 100)
	if !tt.HasText("English") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestPlaceholder shows the placeholder while nothing is chosen.
func TestPlaceholder(t *testing.T) {
	chosen := ""
	tt := ui.NewTester(frame(Props{Selected: &chosen, Placeholder: "Choose a language", Options: []Option{
		{Label: "English", Value: "en"},
	}}), 320, 100)
	if !tt.HasText("Choose a language") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// open presses the trigger, as the pointer goes down on it, to show the
// native menu.
func open(t *testing.T, tt *ui.Tester, label string) {
	t.Helper()
	r, ok := tt.Find(label)
	if !ok {
		t.Fatalf("no trigger showing %q in %q", label, tt.Texts())
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
}

// TestOpens shows the options when the trigger goes down.
func TestOpens(t *testing.T) {
	chosen := "en"
	tt := ui.NewTester(frame(Props{Selected: &chosen, Options: []Option{
		{Label: "English", Value: "en"},
		{Label: "简体中文", Value: "zh"},
	}}), 320, 100)
	open(t, tt, "English")
	if got := tt.Menu(); !slices.Equal(got, []string{"English", "简体中文"}) {
		t.Fatalf("menu %q, want the options", got)
	}
}

// TestChooses takes the choice of the item chosen and runs OnChange.
func TestChooses(t *testing.T) {
	chosen := "en"
	changed := false
	tt := ui.NewTester(frame(Props{Selected: &chosen, Options: []Option{
		{Label: "English", Value: "en"},
		{Label: "简体中文", Value: "zh"},
	}, OnChange: func() { changed = true }}), 320, 100)
	open(t, tt, "English")
	if err := tt.ChooseMenuItem("简体中文"); err != nil {
		t.Fatal(err)
	}
	if chosen != "zh" {
		t.Fatalf("chosen %q, want zh", chosen)
	}
	if !changed {
		t.Fatal("OnChange did not run")
	}
}

// TestGroups shows the label and the options of each group, the label
// as a disabled entry with a line under it.
func TestGroups(t *testing.T) {
	chosen := "en"
	tt := ui.NewTester(frame(Props{Selected: &chosen, Groups: []Group{
		{Label: "Spoken", Options: []Option{{Label: "English", Value: "en"}}},
		{Label: "Formats", Options: []Option{{Label: "US", Value: "us"}}},
	}}), 320, 100)
	open(t, tt, "English")
	if got := tt.Menu(); !slices.Equal(got, []string{"Spoken", "-", "English", "Formats", "-", "US"}) {
		t.Fatalf("menu %q, want the groups with their lines", got)
	}
}

// TestDisabled keeps the menu from opening.
func TestDisabled(t *testing.T) {
	chosen := "en"
	tt := ui.NewTester(frame(Props{Selected: &chosen, Options: []Option{
		{Label: "English", Value: "en"},
	}, Disabled: true}), 320, 100)
	if err := tt.Click("English"); err != nil {
		t.Fatal(err)
	}
	if tt.Menu() != nil {
		t.Fatal("a disabled select opened its menu")
	}
}

// TestDarkMode draws the trigger under the dark appearance too.
func TestDarkMode(t *testing.T) {
	chosen := "en"
	tt := ui.NewTester(frame(Props{Selected: &chosen, Options: []Option{
		{Label: "English", Value: "en"},
	}}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("English") {
		t.Fatal("select missing under dark mode")
	}
}
