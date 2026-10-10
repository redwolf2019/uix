// Package floatinglabelinput provides the Floating Label Input expansion
// of shadcn/ui, a Material-style field: the label sits inside the empty
// field as a placeholder, and floats to the top border on focus or once
// something is typed.
//
//	floatinglabelinput.FloatingLabelInput(c, floatinglabelinput.Props{
//		Label: "Email",
//		Value: &email,
//	})
package floatinglabelinput

import "github.com/egoist/mygo/ui"

// Props describes the field to draw.
type Props struct {
	// Label is what the field asks for; it shows inside the empty field,
	// and floats to the top border on focus or once the field is filled.
	Label string
	// Value is the text being edited.
	Value *string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// FloatingLabelInput draws the field for props and returns it, so a view
// can chain more calls on it. It fills the width of its container.
func FloatingLabelInput(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	return ui.Column(c).FillWidth().Children(func() {
		var field ui.Element
		floated := false
		box := ui.Box(c).FillWidth().Radius(t.Radius).Border(1, t.Border).Background(t.Surface)
		box.Children(func() {
			field = ui.TextInputBase(c, p.Value).FillWidth().TextColor(t.Text)
			floated = *p.Value != "" || field.Focused()
			if p.Disabled {
				field.Disabled(true)
			}
			if !floated {
				// While the field is empty and unfocused, the label is
				// its placeholder, inside it.
				field.Placeholder(p.Label)
			}
		})
		if floated {
			// The label floats to the top border, the background of the
			// field cutting the border behind it, as the expansion's
			// does.
			color := t.TextMuted
			if field.Focused() {
				color = t.Accent
			}
			ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).TextColor(color).
				TextBackground(t.Background).Absolute().Left(t.Space(3)).Top(-t.Space(2.5))
		}
	})
}
