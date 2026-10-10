package multiselect

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Multiselect(c, p)
		})
	}
}

// TestRenders draws the placeholder while nothing is chosen, and the
// chosen options once there are.
func TestRenders(t *testing.T) {
	open := false
	vals := []string{}
	tt := ui.NewTester(frame(Props{
		Values:      &vals,
		Options:     []string{"Editor", "Reviewer", "Admin"},
		Placeholder: "Choose roles",
		Open:        &open,
	}), 360, 160)
	if !tt.HasText("Choose roles") {
		t.Fatalf("texts %q", tt.Texts())
	}
	vals = []string{"Editor", "Admin"}
	tt.Frame()
	if !tt.HasText("Editor, Admin") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestChecks adds a choice in the popup, which the trigger then lists.
func TestChecks(t *testing.T) {
	open := false
	vals := []string{}
	tt := ui.NewTester(frame(Props{
		Values:      &vals,
		Options:     []string{"Editor", "Reviewer", "Admin"},
		Placeholder: "Choose roles",
		Open:        &open,
	}), 360, 220)
	if err := tt.Click("Choose roles"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("a click did not open the popup")
	}
	r, ok := tt.Find("Editor")
	if !ok {
		t.Fatalf("the popup shows %q", tt.Texts())
	}
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	if !slices.Equal(vals, []string{"Editor"}) {
		t.Fatalf("values %q, want [Editor]", vals)
	}
}

// TestUnchecks removes a choice, leaving the others.
func TestUnchecks(t *testing.T) {
	open := false
	vals := []string{"Editor", "Reviewer"}
	tt := ui.NewTester(frame(Props{
		Values:      &vals,
		Options:     []string{"Editor", "Reviewer", "Admin"},
		Placeholder: "Choose roles",
		Open:        &open,
	}), 360, 220)
	if err := tt.Click("Choose roles"); err != nil {
		t.Fatal(err)
	}
	r, ok := tt.Find("Reviewer")
	if !ok {
		t.Fatalf("the popup shows %q", tt.Texts())
	}
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	if !slices.Equal(vals, []string{"Editor"}) {
		t.Fatalf("values %q, want [Editor]", vals)
	}
}

// TestDisabled keeps the popup from opening.
func TestDisabled(t *testing.T) {
	open := false
	vals := []string{}
	tt := ui.NewTester(frame(Props{
		Values:      &vals,
		Options:     []string{"Editor"},
		Placeholder: "Choose roles",
		Open:        &open,
		Disabled:    true,
	}), 360, 160)
	if err := tt.Click("Choose roles"); err != nil {
		t.Fatal(err)
	}
	if open || tt.HasText("Editor") {
		t.Fatal("a disabled select opened its popup")
	}
}
