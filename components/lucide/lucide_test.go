// Package lucide tests the icon set: the shared cache each accessor
// keeps, and drawing one through the icon component.
package lucide

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/icon"
)

// TestShares parses each icon once and returns the same pointer.
func TestShares(t *testing.T) {
	if Search() != Search() {
		t.Fatal("an accessor did not share its parsed SVG")
	}
	if Search() == nil {
		t.Fatal("Search returned no SVG")
	}
	if Search() == Settings() {
		t.Fatal("distinct icons shared a pointer")
	}
}

// TestRenders draws an icon through the uix icon component, named by its
// label for assistive technology.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			icon.Icon(c, icon.Props{SVG: Search(), Label: "Search"})
		})
	}, 200, 120)
	if !tt.HasText("Search") {
		t.Fatalf("texts %q", tt.Texts())
	}
}
