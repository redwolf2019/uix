// Package item provides a shadcn/ui-style Item: a row of a list or a
// settings screen, with a media tile, a title and a description, and
// actions at the far end. A click on the row runs its OnClick, as a
// settings entry opens its page. ItemGroup draws the items of one list
// in a bordered box, with a hairline between them.
//
//	item.Item(c, item.Props{
//		Title:       "Storage",
//		Description: "12 GB of 100 GB used",
//		Media:       drive,
//		OnClick:     app.showStorage,
//	})
package item

import "github.com/egoist/mygo/ui"

// Variant is the look of an item.
type Variant int

const (
	// Default is a transparent item that fills on hover. The default.
	Default Variant = iota
	// Outline draws a border around the item.
	Outline
	// Muted draws the item on a muted face.
	Muted
)

// Size is the density of an item.
type Size int

const (
	// SizeDefault is the standard padding.
	SizeDefault Size = iota
	// SizeSm is the compact padding.
	SizeSm
	// SizeXs is the most compact padding.
	SizeXs
)

// Props describes the item to draw.
type Props struct {
	// Media is the shape of the tile at the start, an SVG in the muted
	// color; the Image takes its place when set.
	Media *ui.SVG
	// Image is a photo filling the tile, cropped to it.
	Image *ui.Bitmap
	// Title is the bold line of the item.
	Title string
	// Description is the muted line under the title.
	Description string
	// Actions draw at the far end of the item, as buttons or a badge;
	// nil draws none.
	Actions func()
	// OnClick runs for a click on the item; nil makes it inert.
	OnClick func()
	// Variant is the look of the item; Default unless set.
	Variant Variant
	// Size is the density of the item; SizeDefault unless set.
	Size Size
}

// Item draws the row and returns it, so a view can chain more calls on
// it. The row fills the width of its container.
func Item(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	padX, padY := t.Space(3), t.Space(2)
	if p.Size == SizeSm {
		padX, padY = t.Space(2.5), t.Space(1.5)
	}
	if p.Size == SizeXs {
		padX, padY = t.Space(2), t.Space(1)
	}
	b := ui.ButtonBase(c).FillWidth().Justify(ui.Start).Gap(t.Space(2.5)).
		Padding(padY, padX).Radius(t.Radius)
	if p.Variant == Outline {
		b.Border(1, t.Border)
	}
	if p.Variant == Muted {
		b.Background(t.Surface)
	}
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.Hovered() || b.Pressed() {
			pp.Fill(r, t.SurfaceHover, t.Radius)
		}
	})
	b.Children(func() {
		ui.Row(c).FillWidth().Gap(t.Space(2.5)).AlignItems(ui.Center).Children(func() {
			tile(c, t, p)
			ui.Column(c).Grow(1).Gap(t.Space(0.25)).Children(func() {
				ui.Text(c, p.Title).SingleLine()
				if p.Description != "" {
					ui.Text(c, p.Description).SingleLine().FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
				}
			})
			if p.Actions != nil {
				ui.Row(c).Gap(t.Space(1)).Shrink(0).Children(func() { p.Actions() })
			}
		})
	})
	if b.Clicked() && p.OnClick != nil {
		p.OnClick()
	}
	return b
}

// tile draws the media of the item: a rounded square with the photo, or
// the icon, or nothing.
func tile(c *ui.Context, t *ui.Theme, p Props) {
	if p.Media == nil && p.Image == nil {
		return
	}
	ui.Box(c).Size(t.Space(9), t.Space(9)).Shrink(0).Radius(t.Radius - 1).Clip().Center().
		Background(t.SurfaceHover).Children(func() {
		if p.Image != nil {
			ui.Image(c, p.Image).Fit(ui.Cover).Size(t.Space(9), t.Space(9))
		} else {
			ui.Icon(c, p.Media).Size(t.Space(4.5), t.Space(4.5)).TextColor(t.TextMuted)
		}
	})
}

// ItemGroup draws the items of a list in a bordered box, a hairline
// between them, and returns the box, so a view can chain more calls on
// it. The box fills the width of its container.
func ItemGroup(c *ui.Context, items []Props) ui.Element {
	t := c.Theme()
	return ui.Column(c).FillWidth().Border(1, t.Border).Radius(t.Radius).Clip().Children(func() {
		for i, it := range items {
			if i > 0 {
				ui.Box(c).FillWidth().Height(1).Background(t.Border)
			}
			Item(c, it)
		}
	})
}
