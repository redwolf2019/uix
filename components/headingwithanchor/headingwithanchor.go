// Package headingwithanchor provides the Heading With Anchor expansion
// of shadcn/ui: a heading with a link button beside it, which copies
// the anchor of the section to the clipboard, as a document's headings
// deep-link to their place. A native app has no URL; the anchor is the
// name the app gives the section, and the copy goes to the clipboard,
// or to the OnAnchor callback instead.
//
//	headingwithanchor.Heading(c, headingwithanchor.Props{
//		Text:   "Installation",
//		Anchor: "installation",
//	})
package headingwithanchor

import "github.com/egoist/mygo/ui"

// Level is the size of a heading.
type Level int

const (
	// H1 is the largest heading, H6 the smallest.
	H1 Level = iota
	H2
	H3
	H4
	H5
	H6
)

// link is the shape of the anchor button, a chain of two links.
var link = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>`))

// Props describes the heading to draw.
type Props struct {
	// Text is the heading.
	Text string
	// Level is the size of the heading; H1 unless set.
	Level Level
	// Anchor is the name of the section, copied to the clipboard as
	// "uix#anchor"; empty draws no anchor button.
	Anchor string
	// OnAnchor runs for a click on the anchor button, after the copy;
	// nil copies to the clipboard alone.
	OnAnchor func()
	// NoCopy keeps the anchor button from copying to the clipboard.
	NoCopy bool
}

// sizes of the heading levels, as the expansion scales them.
var sizes = []float32{2.25, 1.875, 1.5, 1.25, 1.125, 1}

// Heading draws the text with the anchor button at its end, and returns
// the row. The heading fills the width of its container.
func Heading(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	level := int(p.Level)
	if level < 0 || level >= len(sizes) {
		level = 0
	}
	return ui.Row(c).FillWidth().Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
		ui.Text(c, p.Text).FontSize(t.FontSize * sizes[level]).FontWeight(700)
		if p.Anchor != "" {
			b := ui.ButtonBase(c).Size(t.Space(6), t.Space(6)).Shrink(0).Radius(t.Radius)
			b.Draw(func(pp *ui.Painter, r ui.Rect) {
				if b.Hovered() || b.Pressed() {
					pp.Fill(r, t.SurfaceHover, t.Radius)
				}
			})
			b.Children(func() {
				ui.Icon(c, link).Size(t.Space(3.5), t.Space(3.5)).TextColor(t.TextMuted)
			})
			if b.Clicked() {
				if !p.NoCopy {
					c.WriteClipboard("uix#" + p.Anchor)
				}
				if p.OnAnchor != nil {
					p.OnAnchor()
				}
			}
		}
	})
}
