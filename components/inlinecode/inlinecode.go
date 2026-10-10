// Package inlinecode provides the Inline Code expansion of shadcn/ui: a
// code chip in the prose, as `uix add` is written inside a sentence,
// with a muted face and the monospace font.
//
//	inlinecode.InlineCode(c, inlinecode.Props{Text: "uix add button"})
package inlinecode

import "github.com/egoist/mygo/ui"

// Props describes the code chip to draw.
type Props struct {
	// Text is the code, as written in the sentence around it.
	Text string
}

// InlineCode draws the code chip and returns it, so a view can chain
// more calls on it.
func InlineCode(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	e := ui.Box(c).Padding(t.Space(0.75), t.Space(1.5)).Radius(t.Radius - 2).
		Background(t.SurfaceHover)
	e.Children(func() {
		ui.Text(c, p.Text).Font("monospace").FontSize(t.FontSize * 0.875).FontWeight(600).SingleLine()
	})
	return e
}
