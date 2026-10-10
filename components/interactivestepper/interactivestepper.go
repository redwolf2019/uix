// Package interactivestepper provides the Interactive Stepper expansion
// of shadcn/ui: a set of steps that show how far a multi-step process
// has gone, each reachable from the indicator, with Back and Next
// moving between them. The number stepper of uix is a different thing:
// it counts up and down by a step.
//
//	interactivestepper.InteractiveStepper(c, interactivestepper.Props{
//		Step:  &app.step,
//		Steps: []interactivestepper.Step{{Title: "Account", Body: accountForm}},
//	})
package interactivestepper

import (
	"strconv"

	"github.com/egoist/mygo/ui"
)

// Step describes one stage of the stepper.
type Step struct {
	// Title names the step, beside its indicator.
	Title string
	// Body draws the step's content; nil draws nothing for the step.
	Body func(c *ui.Context)
	// Disabled keeps the step from being reached.
	Disabled bool
}

// Props describes the steps to draw.
type Props struct {
	// Steps are the stages, in order.
	Steps []Step
	// Step is the current step; the component writes it as the user
	// moves. It is kept within the steps.
	Step *int
	// OnComplete runs when Next is chosen on the last step.
	OnComplete func()
}

// check is the mark of a finished step.
var check = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg>`))

// InteractiveStepper draws the indicators, the content of the current
// step and the Back and Next buttons, and returns the column, so a view
// can chain more calls on it.
func InteractiveStepper(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	if len(p.Steps) == 0 {
		return ui.Column(c)
	}
	cur := 0
	if p.Step != nil {
		cur = min(max(*p.Step, 0), len(p.Steps)-1)
	}
	return ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		// The indicators, with the line between them colored once the
		// step before it is done.
		ui.Row(c).FillWidth().AlignItems(ui.Center).Children(func() {
			for i := range p.Steps {
				if i > 0 {
					line := t.Border
					if i-1 < cur {
						line = t.Accent
					}
					ui.Box(c).Height(2).Grow(1).Radius(1).Background(line)
				}
				indicator(c, t, p, i, cur)
			}
		})
		if p.Steps[cur].Body != nil {
			p.Steps[cur].Body(c)
		}
		ui.Row(c).FillWidth().Justify(ui.SpaceBetween).Children(func() {
			backButton(c, t, p, cur)
			nextButton(c, t, p, cur)
		})
	})
}

// indicator draws one step's circle and title; a click on the circle
// moves the stepper to the step, unless it is disabled.
func indicator(c *ui.Context, t *ui.Theme, p Props, i, cur int) {
	ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1.5)).Children(func() {
		b := ui.ButtonBase(c).Size(t.Space(5), t.Space(5)).Radius(t.Space(2.5)).Shrink(0)
		switch {
		case i < cur || i == cur:
			b.Background(t.Accent)
		default:
			b.Background(t.Surface).Border(1, t.Border)
		}
		if p.Steps[i].Disabled {
			b.Opacity(0.4)
		} else {
			b.Children(func() {
				if i < cur {
					ui.Icon(c, check).Size(t.Space(3), t.Space(3)).TextColor(t.AccentText)
				} else {
					color := t.TextMuted
					if i == cur {
						color = t.AccentText
					}
					ui.Text(c, strconv.Itoa(i+1)).SingleLine().TextColor(color)
				}
			})
			if b.Clicked() {
				*p.Step = i
			}
		}
		ui.Text(c, p.Steps[i].Title).SingleLine()
	})
}

// backButton moves the stepper to the previous step that can be
// reached; there is none on the first step.
func backButton(c *ui.Context, t *ui.Theme, p Props, cur int) {
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	b.Children(func() { ui.Text(c, "Back").SingleLine() })
	if b.Clicked() && cur > 0 {
		for j := cur - 1; j >= 0; j-- {
			if !p.Steps[j].Disabled {
				*p.Step = j
				break
			}
		}
	}
}

// nextButton moves the stepper on, or finishes on the last step.
func nextButton(c *ui.Context, t *ui.Theme, p Props, cur int) {
	last := cur == len(p.Steps)-1
	label := "Next"
	if last {
		label = "Finish"
	}
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).
		Background(t.Accent).TextColor(t.AccentText)
	b.Children(func() { ui.Text(c, label).SingleLine() })
	if b.Clicked() {
		if last {
			if p.OnComplete != nil {
				p.OnComplete()
			}
			return
		}
		for j := cur + 1; j < len(p.Steps); j++ {
			if !p.Steps[j].Disabled {
				*p.Step = j
				break
			}
		}
	}
}
