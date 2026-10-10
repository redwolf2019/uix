package autosizetextarea

import (
	"strings"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			AutosizeTextarea(c, p)
		})
	}
}

// TestRenders draws the field, named by its Label.
func TestRenders(t *testing.T) {
	v := "hello"
	tt := ui.NewTester(frame(Props{Value: &v, Label: "Notes"}), 320, 200)
	if _, ok := tt.Find("Notes"); !ok {
		t.Fatal("field not found")
	}
}

// TestGrows makes the field taller as the text gains lines, and stops
// past MaxRows, which scrolls instead.
func TestGrows(t *testing.T) {
	one := "a"
	two := strings.Repeat("line\n", 5) + "x" // six lines
	tt := ui.NewTester(frame(Props{Value: &one, Label: "Notes"}), 320, 300)
	_, ok := tt.Find("Notes")
	if !ok {
		t.Fatal("field not found")
	}
	h1 := rectH(tt)
	one = two
	tt.Frame()
	h2 := rectH(tt)
	if h2 <= h1+20 {
		t.Fatalf("the field did not grow: %f to %f", h1, h2)
	}
	one = strings.Repeat("line\n", 12) + "x" // thirteen lines
	tt.Frame()
	h3 := rectH(tt)
	if h3 > h2+2 {
		t.Fatalf("the field grew past its limit: %f to %f", h2, h3)
	}
}

// rectH returns the height of the field with the Notes label.
func rectH(tt *ui.Tester) float32 {
	r, ok := tt.Find("Notes")
	if !ok {
		panic("field not found")
	}
	return r.H
}

// TestDisabled keeps the field from being edited.
func TestDisabled(t *testing.T) {
	v := "hello"
	tt := ui.NewTester(frame(Props{Value: &v, Label: "Notes", Disabled: true}), 320, 200)
	r, _ := tt.Find("Notes")
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	tt.Type("!")
	if v != "hello" {
		t.Fatalf("a disabled field changed to %q", v)
	}
}

// TestDarkMode draws the field under the dark appearance too.
func TestDarkMode(t *testing.T) {
	v := "hello"
	tt := ui.NewTester(frame(Props{Value: &v, Label: "Notes"}), 320, 200)
	tt.SetDark(true)
	if _, ok := tt.Find("Notes"); !ok {
		t.Fatal("field missing under dark mode")
	}
}
