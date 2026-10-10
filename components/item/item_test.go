package item

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Item(c, p)
		})
	}
}

// TestRenders draws the title, the description and the action.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Item(c, Props{
				Title:       "Storage",
				Description: "12 GB of 100 GB used",
				Actions:     func() { ui.Text(c, "Manage") },
			})
		})
	}, 320, 100)
	for _, want := range []string{"Storage", "12 GB of 100 GB used", "Manage"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestClick runs the action of the item for a click on it.
func TestClick(t *testing.T) {
	opened := false
	tt := ui.NewTester(frame(Props{Title: "Storage", OnClick: func() { opened = true }}), 320, 80)
	if err := tt.Click("Storage"); err != nil {
		t.Fatal(err)
	}
	if !opened {
		t.Fatal("the click did not run the action")
	}
}

// TestGroup draws the items of a list in a bordered box, each of the
// rows showing its line.
func TestGroup(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			ItemGroup(c, []Props{
				{Title: "Profile", Description: "Your name and photo"},
				{Title: "Billing", Description: "Your plan and invoices"},
				{Title: "Privacy", Description: "What apps may see"},
			})
		})
	}, 320, 240)
	for _, want := range []string{"Profile", "Billing", "Privacy"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestVariants draw every look of an item.
func TestVariants(t *testing.T) {
	for _, v := range []Variant{Default, Outline, Muted} {
		tt := ui.NewTester(frame(Props{Title: "Storage", Variant: v}), 320, 80)
		if !tt.HasText("Storage") {
			t.Fatalf("variant %d lost the title: %q", v, tt.Texts())
		}
	}
}

// TestDarkMode draws an item under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Title: "Storage"}), 320, 80)
	tt.SetDark(true)
	if !tt.HasText("Storage") {
		t.Fatal("item missing under dark mode")
	}
}
