// Package view builds the loading example: a checkout of MyGo native UI
// that shows the remaining components of the shadcn/ui expansions
// together — the magic back button, the interactive stepper with the
// infinite-scroll catalog and the loading button, and the responsive
// modal that confirms the order. The catalog loads more products as it
// nears its end; the Pay button turns into a spinner for a moment and
// then opens the order as a modal.
package view

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/infinitescroll"
	"github.com/smallnest/uix/components/interactivestepper"
	"github.com/smallnest/uix/components/loadingbutton"
	"github.com/smallnest/uix/components/magicback"
	"github.com/smallnest/uix/components/responsivemodal"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 560

// State is the state of the example; the controls edit it in place.
var State = LoadingState{}

// LoadingState holds the example: the catalog and its scroll, the step
// of the stepper, the payment and the order modal.
type LoadingState struct {
	// Products are the rows of the catalog, loaded page by page.
	Products []string
	// Catalog is where the infinite scroll keeps its place and its
	// trigger.
	Catalog infinitescroll.State
	// Step is the current step of the stepper.
	Step int
	// Paying and PayStart make the Pay button a spinner for a moment,
	// and then open the order modal.
	Paying   bool
	PayStart time.Time
	// Open is whether the order modal shows.
	Open bool
	// BackLabel is what the header says under the title, for the back
	// button.
	BackLabel string
	// Paid is how many orders the session placed, for the tests.
	Paid int
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = LoadingState{Products: products(40)}
}

// products names n products of the catalog.
func products(n int) []string {
	names := make([]string, n)
	for i := range names {
		names[i] = fmt.Sprintf("Product %d", i+1)
	}
	return names
}

// LoadingView draws the example: the header with the back button, the
// stepper with the catalog and the pay button, and the order modal.
func LoadingView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	// The payment finishes a moment after Pay is clicked, opening the
	// order modal. Why the zero check: PayStart is zero while Paying is
	// set by hand, as the snapshot command does, and then the payment
	// must wait for a click, not finish at once.
	if State.Paying && !State.PayStart.IsZero() && time.Since(State.PayStart) >= 400*time.Millisecond {
		State.Paying = false
		State.Open = true
		State.Paid++
	}
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).FillWidth().AlignItems(ui.Center).Gap(t.Space(2)).Padding(24, 24, 24, 12).Children(func() {
			magicback.MagicBack(c, magicback.Props{
				Back:   State.Step > 0,
				OnBack: func() { State.Step = 0 },
				OnHome: func() { State.BackLabel = "At the shop" },
			})
			ui.Column(c).Grow(1).Children(func() {
				ui.Text(c, "Checkout").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
				if State.BackLabel != "" {
					ui.Text(c, State.BackLabel).TextColor(t.TextMuted)
				}
			})
		})
		ui.Column(c).FillWidth().Padding(24, 12, 24, 24).Children(func() {
			interactivestepper.InteractiveStepper(c, interactivestepper.Props{
				Step: &State.Step,
				Steps: []interactivestepper.Step{
					{Title: "Catalog", Body: func(c *ui.Context) { catalog(c, t) }},
					{Title: "Review", Body: func(c *ui.Context) { review(c, t) }},
					{Title: "Pay", Body: func(c *ui.Context) { pay(c, t) }},
				},
			})
		})
	})
	responsivemodal.ResponsiveModal(c, responsivemodal.Props{
		Open:  &State.Open,
		Title: "Order placed",
		Body: func(c *ui.Context) {
			ui.Text(c, "Your order is on its way.")
		},
		Actions: func(c *ui.Context) {
			ui.Row(c).FillWidth().Justify(ui.End).Children(func() {
				button.Button(c, button.Props{Label: "Done", OnClick: func() { State.Open = false }})
			})
		},
	})
}

// catalog draws the step that browses the catalog: a list that loads
// the next page as its end comes near.
func catalog(c *ui.Context, t *ui.Theme) {
	ui.Box(c).FillWidth().Height(300).Children(func() {
		infinitescroll.InfiniteScroll(c, infinitescroll.Props{
			State: &State.Catalog,
			OnMore: func() {
				State.Products = append(State.Products, products(10)...)
			},
			Children: func() {
				ui.Column(c).FillWidth().Children(func() {
					for _, name := range State.Products {
						ui.Row(c).FillWidth().Height(34).AlignItems(ui.Center).Padding(0, 12).Children(func() {
							ui.Text(c, name).Grow(1).TextColor(t.Text)
							ui.Text(c, "$4.00").TextColor(t.TextMuted)
						})
					}
				})
			},
		})
	})
}

// review draws the step that sums the catalog.
func review(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(1.5)).Children(func() {
		ui.Text(c, fmt.Sprintf("%d products in the catalog.", len(State.Products))).TextColor(t.Text)
		ui.Text(c, "Next, choose how to pay.").TextColor(t.TextMuted)
	})
}

// pay draws the step with the Pay button, which turns into a spinner
// while the payment goes through.
func pay(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		ui.Text(c, fmt.Sprintf("Total: $%d.00", len(State.Products)*4)).TextColor(t.TextMuted)
		loadingbutton.LoadingButton(c, loadingbutton.Props{
			Label:   fmt.Sprintf("Pay $%d.00", len(State.Products)*4),
			Loading: State.Paying,
			OnClick: func() {
				State.Paying = true
				State.PayStart = time.Now()
			},
		})
	})
}
