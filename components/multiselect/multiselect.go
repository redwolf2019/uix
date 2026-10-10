// Package multiselect provides the Multiple Selector expansion of
// shadcn/ui: a field choosing several of its options, whose trigger
// lists them, and a popup with a check box for each.
//
//	multiselect.Multiselect(c, multiselect.Props{
//		Values:      &roles,
//		Options:     []string{"Editor", "Reviewer", "Admin"},
//		Placeholder: "Choose roles",
//	})
package multiselect

import (
	"slices"
	"strings"

	"github.com/egoist/mygo/ui"
)

// Props describes the multi-select to draw.
type Props struct {
	// Values are the chosen options; the trigger lists them, and the
	// popup toggles them.
	Values *[]string
	// Options are the choices in the popup.
	Options []string
	// Placeholder shows in the trigger while nothing is chosen, and
	// names the field to assistive technology when Label is empty.
	Placeholder string
	// Label names the field to assistive technology, and the tests and
	// apps that find it by it; Placeholder when empty.
	Label string
	// Open is whether the popup shows; the trigger toggles it, and
	// clicking outside or pressing Escape closes it.
	Open *bool
	// Disabled keeps the select from opening.
	Disabled bool
}

// Multiselect draws the trigger and the popup for props and returns the
// trigger, so a view can chain more calls on it.
func Multiselect(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	name := p.Label
	if name == "" {
		name = p.Placeholder
	}
	anchor := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Gap(t.Space(1.5)).Radius(t.Radius).
		Justify(ui.SpaceBetween).MinWidth(t.Space(35)).Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	if name != "" {
		anchor.Label(name)
	}
	if p.Disabled {
		anchor.Disabled(true)
		anchor.TextColor(t.TextMuted)
	} else if anchor.Clicked() {
		*p.Open = !*p.Open
	}
	anchor.Children(func() {
		if label := strings.Join(*p.Values, ", "); label != "" {
			ui.Text(c, label).SingleLine()
		} else if p.Placeholder != "" {
			ui.Text(c, p.Placeholder).SingleLine().TextColor(t.TextMuted)
		}
		chevron(c, t)
	})
	ui.Popover(c, anchor, p.Open, func() {
		ui.Column(c).Gap(2).Children(func() {
			for _, opt := range p.Options {
				on := slices.Contains(*p.Values, opt)
				cb := ui.Checkbox(c, &on, opt)
				if cb.Changed() {
					if on {
						*p.Values = append(*p.Values, opt)
					} else {
						*p.Values = slices.DeleteFunc(*p.Values, func(s string) bool { return s == opt })
					}
				}
			}
		})
	})
	return anchor
}

// chevron draws the arrow of the trigger.
func chevron(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(t.Space(2.5), t.Space(2.5)).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		var path ui.Path
		path.MoveTo(r.X+r.W*0.1, r.Y+r.H*0.3).LineTo(r.X+r.W*0.5, r.Y+r.H*0.7).LineTo(r.X+r.W*0.9, r.Y+r.H*0.3)
		p.StrokePath(&path, 1.5, t.TextMuted)
	})
}
