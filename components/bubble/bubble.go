// Package bubble provides a shadcn/ui-style Bubble: the framed surface
// of one message in a conversation, as chat apps draw the current
// user's turn in the accent color. The variant sets the look, from the
// strong default bubble to an unframed ghost; the alignment puts the
// bubble at the start or the end of its row. The message component
// builds the row around it; a Bubble alone is the surface.
//
//	bubble.Bubble(c, bubble.Props{
//		Text:    "Done!",
//		Variant: bubble.Default,
//		Align:   bubble.End,
//	})
package bubble

import "github.com/egoist/mygo/ui"

// Variant is the look of a bubble.
type Variant int

const (
	// Default is the strong primary bubble, as the current user's
	// message.
	Default Variant = iota
	// Secondary is the standard neutral bubble, as the other side's.
	Secondary
	// Muted is a bubble of lower emphasis.
	Muted
	// Tinted is a soft accent face.
	Tinted
	// Outline is a bubble with a border and no face.
	Outline
	// Ghost is unframed and fills the row, as assistant text does.
	Ghost
	// Destructive signals an error or a failed action.
	Destructive
)

// Align is the side of the row a bubble sits on.
type Align int

const (
	// Start puts the bubble at the start of the row. The default.
	Start Align = iota
	// End puts the bubble at the end of the row.
	End
)

// Props describes the bubble to draw.
type Props struct {
	// Text is the message inside the bubble.
	Text string
	// Children draw inside the bubble instead of Text, as rich content
	// of their own.
	Children func()
	// Variant is the look of the bubble; Default unless set.
	Variant Variant
	// Align is the side of the row the bubble sits on; Start unless set.
	Align Align
}

// Bubble draws the bubble on its side of a row and returns the row, so
// a view can chain more calls on it. The bubble sizes to its content,
// capped at 80% of the row, except Ghost, which fills the row.
func Bubble(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	// The look of each variant: the face and the text, and a border for
	// the Outline look.
	face, fg, border := ui.Color{}, t.Text, ui.Color{}
	switch p.Variant {
	case Default:
		face, fg = t.Accent, t.AccentText
	case Secondary:
		face = t.Surface
	case Muted:
		face, fg = t.SurfaceHover, t.TextMuted
	case Tinted:
		face, fg = t.Accent.Alpha(0.15), t.Accent
	case Outline:
		face, border = t.Background, t.Border
	case Destructive:
		face, fg = t.Danger, ui.Hex("#ffffff")
	}
	justify := ui.Start
	if p.Align == End {
		justify = ui.End
	}
	// The bubble is built among the row's children, so it takes part in
	// the row's layout.
	return ui.Row(c).FillWidth().Justify(justify).Children(func() {
		b := ui.Box(c)
		if p.Variant == Ghost {
			b.FillWidth()
		} else {
			b.Padding(t.Space(2), t.Space(3)).Radius(t.Radius + 2).MaxWidthPercent(0.8).
				Background(face).TextColor(fg)
			if border.A > 0 {
				b.Border(1, border)
			}
		}
		b.Children(func() {
			if p.Children != nil {
				p.Children()
			} else if p.Text != "" {
				ui.Text(c, p.Text)
			}
		})
	})
}
