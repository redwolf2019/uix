package main

import (
	"math"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/loading/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.LoadingView, view.Width, view.Height)
}

// TestRenders draws the header, the stepper and the first product.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{"Checkout", "Catalog", "Product 1", "Next"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestLoadsMore loads the next page of products as the catalog nears
// its end.
func TestLoadsMore(t *testing.T) {
	tt := newTester(t)
	if len(view.State.Products) != 40 {
		t.Fatalf("started with %d products, want 40", len(view.State.Products))
	}
	view.State.Catalog.Scroll.Y = math.MaxFloat32
	tt.Frame()
	if len(view.State.Products) != 50 {
		t.Fatalf("the end loaded %d products, want 50", len(view.State.Products))
	}
}

// TestPays places the order: the Pay button spins, then the modal
// opens.
func TestPays(t *testing.T) {
	tt := newTester(t)
	// Step to the Pay step.
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Pay $160.00"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Paying {
		t.Fatal("the button did not start loading")
	}
	time.Sleep(500 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Order placed") {
		t.Fatalf("the order modal is missing in %q", tt.Texts())
	}
	if view.State.Paid != 1 {
		t.Fatalf("paid %d, want 1", view.State.Paid)
	}
}

// TestGoesHome shows the home message of the back button while there is
// nowhere to go back to.
func TestGoesHome(t *testing.T) {
	tt := newTester(t)
	tt.ClickAt(40, 40)
	if !tt.HasText("At the shop") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Checkout") {
		t.Fatal("loading missing under dark mode")
	}
}
