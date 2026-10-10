// Package blockquote provides a shadcn/ui-style Block Quote, as the
// expansion of the same name: a quote on a muted face with an accent
// bar on the start, and the author line at the far end.
//
//	blockquote.Blockquote(c, blockquote.Props{
//		Text:   "The only thing we have to fear is fear itself.",
//		Author: "Franklin D. Roosevelt",
//	})
package blockquote

import "github.com/egoist/mygo/ui"

// Props describes the quote to draw.
type Props struct {
	// Text is the quote.
	Text string
	// Children draw inside the quote instead of Text, as rich content
	// of their own.
	Children func()
	// Author is the attribution at the far end; empty draws none.
	Author string
}

// Blockquote draws the quote and returns it, so a view can chain more
// calls on it. The quote fills the width of its container.
func Blockquote(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	return ui.Box(c).FillWidth().Radius(t.Radius).Clip().Background(t.Surface).Children(func() {
		ui.Row(c).FillWidth().AlignItems(ui.Stretch).Children(func() {
			// The accent bar on the start, as a quotation marks it.
			ui.Box(c).Width(3).Shrink(0).Background(t.Accent)
			ui.Column(c).Grow(1).Padding(t.Space(3), t.Space(4)).Gap(t.Space(0.5)).Children(func() {
				ui.Row(c).Gap(t.Space(2)).Children(func() {
					// The decorative quotation mark, large and muted.
					ui.Text(c, "“").FontSize(t.FontSize * 2).FontWeight(700).TextColor(t.TextMuted)
					ui.Column(c).Grow(1).Children(func() {
						if p.Children != nil {
							p.Children()
						} else if p.Text != "" {
							ui.Text(c, p.Text).TextAlign(ui.Start)
						}
					})
				})
				if p.Author != "" {
					ui.Row(c).FillWidth().Justify(ui.End).Children(func() {
						ui.Text(c, p.Author).FontWeight(600)
					})
				}
			})
		})
	})
}
