// Package responsivemodal provides the Responsive Modal expansion of
// shadcn/ui: a dialog that shows centered on a wide window, and as a
// bottom sheet on a narrow one, as a phone app's sheet rises from the
// bottom.
//
//	responsivemodal.ResponsiveModal(c, responsivemodal.Props{
//		Open:  &app.confirm,
//		Title: "Confirm",
//		Body:  body,
//	})
package responsivemodal

import "github.com/egoist/mygo/ui"

// Props describes the modal to draw.
type Props struct {
	// Open is the bool that shows and hides the modal; the backdrop and
	// Escape set it to false.
	Open *bool
	// Title is the heading of the modal.
	Title string
	// Body draws the content of the modal, under the title.
	Body func(c *ui.Context)
	// Actions draws the action row, at the bottom of the modal.
	Actions func(c *ui.Context)
	// Breakpoint is the width below which the modal becomes a bottom
	// sheet; 640 unless set.
	Breakpoint float32
}

// ResponsiveModal draws the modal for props and returns the panel, so a
// view can chain more calls on it. The backdrop dims the window;
// clicking it or pressing Escape closes the modal.
func ResponsiveModal(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	w, h := c.Size()
	bp := p.Breakpoint
	if bp <= 0 {
		bp = 640
	}
	return ui.DialogBase(c, p.Open, func(back, panel ui.Element) {
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		if w >= bp {
			panel.Padding(t.Space(5)).Gap(t.Space(3)).Radius(t.Radius+4).Background(t.Background).
				MinWidth(t.Space(48)).MaxWidth(t.Space(72)).MaxHeight(h-t.Space(10)).
				Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		} else {
			// The sheet rises from the bottom of the window; its top
			// corners round, its bottom meets the edge.
			panel.Absolute().Left(0).Right(0).Bottom(0).
				Padding(t.Space(5)).Gap(t.Space(3)).Radius(t.Radius+4, t.Radius+4, 0, 0).
				Background(t.Background).MaxHeight(h-t.Space(4)).
				Shadow(0, -6, 24, 0, ui.RGBA(0, 0, 0, 0.3))
		}
		panel.Children(func() {
			ui.Text(c, p.Title).Bold()
			if p.Body != nil {
				p.Body(c)
			}
			if p.Actions != nil {
				p.Actions(c)
			}
		})
	})
}
