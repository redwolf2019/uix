// Package loadingbutton provides the Loading Button expansion of
// shadcn/ui: a button that turns into a spinner while its work runs, and
// will not be clicked again until it is done.
//
//	loadingbutton.LoadingButton(c, loadingbutton.Props{
//		Label:   "Save",
//		Loading: app.saving,
//		OnClick: save,
//	})
package loadingbutton

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the button to draw.
type Props struct {
	// Label is what the button says; it keeps showing while Loading,
	// beside the spinner, so the button keeps its width.
	Label string
	// Loading turns the label into a spinner, and keeps the button from
	// being clicked, until the work it started finishes.
	Loading bool
	// OnClick runs for a click while the button is not loading.
	OnClick func()
}

// LoadingButton draws the button for props and returns it, so a view can
// chain more calls on it.
func LoadingButton(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).Gap(t.Space(2)).
		Background(t.Accent).TextColor(t.AccentText)
	if p.Loading {
		b.Disabled(true)
	}
	b.Children(func() {
		if p.Loading {
			// A spoke of its own, so the spinner matches the button's
			// text color, which the native Spinner's muted spokes do not.
			spin := b.Loop("spin", 800*time.Millisecond, ui.Linear)
			ui.Box(c).Size(t.Space(3.5), t.Space(3.5)).Shrink(0).Rotate(spin * 360).Draw(func(p *ui.Painter, r ui.Rect) {
				var path ui.Path
				path.MoveTo(r.X+r.W/2, r.Y+r.H*0.1).LineTo(r.X+r.W/2, r.Y+r.H*0.35)
				p.StrokePath(&path, 2, t.AccentText)
			})
		}
		ui.Text(c, p.Label).SingleLine()
	})
	if !p.Loading && b.Clicked() && p.OnClick != nil {
		p.OnClick()
	}
	return b
}
