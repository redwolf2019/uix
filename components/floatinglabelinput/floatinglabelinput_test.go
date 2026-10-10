package floatinglabelinput

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			FloatingLabelInput(c, p)
		})
	}
}

// TestIdle hides the floating label while the field is empty, and shows
// the placeholder in its place.
func TestIdle(t *testing.T) {
	v := ""
	tt := ui.NewTester(frame(Props{Label: "Email", Value: &v}), 320, 100)
	if tt.HasText("Email") {
		t.Fatalf("texts %q, want no floating label", tt.Texts())
	}
}

// TestFloatsOnFocus floats the label when the field takes the focus.
func TestFloatsOnFocus(t *testing.T) {
	v := ""
	tt := ui.NewTester(frame(Props{Label: "Email", Value: &v}), 320, 100)
	tt.ClickAt(50, 30)
	if !tt.HasText("Email") {
		t.Fatalf("texts %q, want the label floated", tt.Texts())
	}
}

// TestFloatsOnValue floats the label once the field is filled.
func TestFloatsOnValue(t *testing.T) {
	v := "me@example.com"
	tt := ui.NewTester(frame(Props{Label: "Email", Value: &v}), 320, 100)
	if !tt.HasText("Email") {
		t.Fatalf("texts %q, want the label floated", tt.Texts())
	}
}

// TestDisabled keeps the field from being edited.
func TestDisabled(t *testing.T) {
	v := "me@example.com"
	tt := ui.NewTester(frame(Props{Label: "Email", Value: &v, Disabled: true}), 320, 100)
	tt.ClickAt(50, 30)
	tt.Type("!")
	if v != "me@example.com" {
		t.Fatalf("a disabled field changed to %q", v)
	}
}

// TestDarkMode draws the field under the dark appearance too.
func TestDarkMode(t *testing.T) {
	v := "me@example.com"
	tt := ui.NewTester(frame(Props{Label: "Email", Value: &v}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("Email") {
		t.Fatal("label missing under dark mode")
	}
}
