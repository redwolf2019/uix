// Package magicback provides the Magic Back Button expansion of
// shadcn/ui: a round back button that returns the app to where it came
// from, or to its home when it has no history to go back to, as the web
// button sends a first-time visitor to the homepage.
//
//	magicback.MagicBack(c, magicback.Props{
//		Back:  app.canGoBack,
//		OnBack: goBack,
//		OnHome: goHome,
//	})
package magicback

import "github.com/egoist/mygo/ui"

// chevronLeft is the arrow of the button.
var chevronLeft = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>`))

// Props describes the back button to draw.
type Props struct {
	// Back is whether the app has a place to go back to; the button
	// goes back when true, and home when not, as the web button does
	// for a first-time visitor.
	Back bool
	// OnBack runs for a click while Back is true.
	OnBack func()
	// OnHome runs for a click while Back is false.
	OnHome func()
	// Label draws text beside the arrow; empty draws the arrow alone.
	Label string
}

// MagicBack draws the round back button and returns it, so a view can
// chain more calls on it.
func MagicBack(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	b := ui.ButtonBase(c).Padding(t.Space(2)).Radius(t.Space(3)).Gap(t.Space(1.5)).
		Background(t.Surface).Border(1, t.Border).TextColor(t.Text)
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.Hovered() || b.Pressed() {
			pp.Fill(r, t.SurfaceHover, t.Space(3))
		}
	})
	b.Children(func() {
		ui.Icon(c, chevronLeft).Size(t.Space(3.5), t.Space(3.5))
		if p.Label != "" {
			ui.Text(c, p.Label).SingleLine()
		}
	})
	if b.Clicked() {
		if p.Back {
			if p.OnBack != nil {
				p.OnBack()
			}
		} else if p.OnHome != nil {
			p.OnHome()
		}
	}
	return b
}
