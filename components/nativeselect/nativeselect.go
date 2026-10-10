// Package nativeselect provides a shadcn/ui-style Native Select: a
// drop-down that uses the platform's own menu, as a native HTML select
// does on the web, for the behavior and the look of the system. The
// select component of uix builds its own popup; use this one for the
// native menu, or when options come in groups. The options are entries
// of the menu, the chosen one checked; a group draws its label and a
// line before its options.
//
//	nativeselect.NativeSelect(c, nativeselect.Props{
//		Selected:    &language,
//		Placeholder: "Choose a language",
//		Options: []nativeselect.Option{
//			{Label: "English", Value: "en"},
//			{Label: "简体中文", Value: "zh"},
//		},
//	})
package nativeselect

import "github.com/egoist/mygo/ui"

// Option is one choice of the select.
type Option struct {
	// Label is the text of the choice.
	Label string
	// Value is what the app sees as the choice; the Label when empty.
	Value string
	// Disabled keeps the choice from being taken.
	Disabled bool
}

// Group is a labeled run of options of the select.
type Group struct {
	// Label is the heading of the group, drawn as a disabled entry
	// above its options, as an optgroup of a native select.
	Label string
	// Options are the choices of the group.
	Options []Option
}

// Props describes the select to draw.
type Props struct {
	// Selected is the value chosen, which the select changes in place.
	Selected *string
	// Options are the choices of the select, when it is not grouped.
	Options []Option
	// Groups label runs of options; when set, they replace Options.
	Groups []Group
	// Placeholder shows in the trigger while nothing is chosen, in the
	// muted color.
	Placeholder string
	// Label names the select for assistive technology.
	Label string
	// Disabled keeps the select from opening.
	Disabled bool
	// OnChange runs when a choice is taken.
	OnChange func()
}

// NativeSelect draws the trigger and its menu and returns it, so a view
// can chain more calls on it. The trigger shows the label of the chosen
// value, or the placeholder while none is chosen; the platform shows
// the menu on a press.
func NativeSelect(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Gap(t.Space(1.5)).
		Radius(t.Radius).Justify(ui.SpaceBetween).MinWidth(t.Space(35)).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	if p.Label != "" {
		b.Label(p.Label)
	}
	if p.Disabled {
		b.Disabled(true).TextColor(t.TextMuted)
	}
	b.Children(func() {
		switch chosen := *p.Selected; {
		case chosen != "":
			ui.Text(c, labelOf(p, chosen)).SingleLine()
		case p.Placeholder != "":
			ui.Text(c, p.Placeholder).SingleLine().TextColor(t.TextMuted)
		}
		chevron(c, t)
	})
	b.Menu(func(m *ui.Menu) {
		if len(p.Groups) > 0 {
			for _, g := range p.Groups {
				if g.Label != "" {
					m.Item(g.Label).Disabled(true)
					m.Separator()
				}
				options(c, m, p, g.Options)
			}
			return
		}
		options(c, m, p, p.Options)
	})
	return b
}

// labelOf returns the label of the chosen value, for the trigger to
// show; the value itself when no option carries it.
func labelOf(p Props, chosen string) string {
	best := chosen
	for _, o := range p.Options {
		value := o.Value
		if value == "" {
			value = o.Label
		}
		if value == chosen {
			return o.Label
		}
	}
	for _, g := range p.Groups {
		for _, o := range g.Options {
			value := o.Value
			if value == "" {
				value = o.Label
			}
			if value == chosen {
				return o.Label
			}
		}
	}
	return best
}

// options adds the options of a run to the menu, the chosen one
// checked, and takes the choice of the one chosen.
func options(c *ui.Context, m *ui.Menu, p Props, opts []Option) {
	for _, o := range opts {
		value := o.Value
		if value == "" {
			value = o.Label
		}
		item := m.Item(o.Label)
		item.Checked(value == *p.Selected).Disabled(o.Disabled)
		if item.Chosen() {
			*p.Selected = value
			if p.OnChange != nil {
				p.OnChange()
			}
		}
	}
}

// chevron draws the arrow of the trigger.
func chevron(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(t.Space(2.5), t.Space(2.5)).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		var path ui.Path
		path.MoveTo(r.X+r.W*0.1, r.Y+r.H*0.3).LineTo(r.X+r.W*0.5, r.Y+r.H*0.7).LineTo(r.X+r.W*0.9, r.Y+r.H*0.3)
		p.StrokePath(&path, 1.5, t.TextMuted)
	})
}
