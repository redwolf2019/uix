// Package message provides a shadcn/ui-style Message: one turn of a
// conversation, with the avatar, the sender line, the bubble and the
// status line around the message surface. The alignment sets which side
// of the row the turn sits on; the header stays at the start either way,
// and the status line follows the bubble. MessageGroup stacks the turns
// of one sender, the later ones without an avatar keeping the room of
// the first.
//
//	message.Message(c, message.Props{
//		Align:   message.End,
//		Name:    "You",
//		Header:  "You · now",
//		Text:    "Done!",
//		Variant: bubble.Default,
//	})
package message

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/avatar"
	"github.com/smallnest/uix/components/bubble"
)

// Props describes the message to draw.
type Props struct {
	// Align is the side of the row the turn sits on; Start unless set.
	Align bubble.Align
	// Name is the sender: its initials fill the avatar, and it labels
	// it for assistive technology. Empty draws an empty slot, as for a
	// group's earlier turns.
	Name string
	// Image is the photo of the sender, cropped to the avatar. Nil
	// draws the initials of the name.
	Image *ui.Bitmap
	// Header is the line above the bubble, such as the sender and the
	// time. It stays at the start of the row either way.
	Header string
	// Footer is the line below the bubble, such as a status. It follows
	// the side of the turn.
	Footer string
	// Variant is the look of the bubble; the default of the bubble
	// component unless set.
	Variant bubble.Variant
	// Text is the message inside the bubble.
	Text string
	// Children draw inside the bubble instead of Text, as rich content
	// of their own.
	Children func()
}

// Message draws the turn and returns it, so a view can chain more calls
// on it. The avatar sits at the bottom, at the side of the turn.
func Message(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	return ui.Row(c).FillWidth().Gap(t.Space(2)).AlignItems(ui.End).Children(func() {
		if p.Align == bubble.End {
			content(c, t, p)
			avatarSlot(c, t, p)
		} else {
			avatarSlot(c, t, p)
			content(c, t, p)
		}
	})
}

// content draws the header, the bubble and the footer of the turn: the
// header at the start of the row, the bubble and the footer at the side
// of the turn.
func content(c *ui.Context, t *ui.Theme, p Props) {
	side := ui.Start
	if p.Align == bubble.End {
		side = ui.End
	}
	ui.Column(c).Grow(1).AlignItems(side).Gap(t.Space(0.5)).Children(func() {
		if p.Header != "" {
			ui.Row(c).FillWidth().Justify(ui.Start).Children(func() {
				ui.Text(c, p.Header).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted)
			})
		}
		bubble.Bubble(c, bubble.Props{Text: p.Text, Children: p.Children, Variant: p.Variant, Align: p.Align})
		if p.Footer != "" {
			ui.Text(c, p.Footer).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted)
		}
	})
}

// avatarSlot draws the avatar of the turn at the bottom, or an empty
// slot of its size when the turn has no sender, as a group's earlier
// turns keep the alignment of the first.
func avatarSlot(c *ui.Context, t *ui.Theme, p Props) {
	if p.Name == "" && p.Image == nil {
		ui.Box(c).Size(t.FontSize*2.25, t.FontSize*2.25).Shrink(0)
		return
	}
	avatar.Avatar(c, avatar.Props{Name: p.Name, Image: p.Image}).Shrink(0)
}

// MessageGroup draws the turns of one sender stacked closely, and
// returns the column, so a view can chain more calls on it.
func MessageGroup(c *ui.Context, turns func()) ui.Element {
	t := c.Theme()
	return ui.Column(c).FillWidth().Gap(t.Space(0.5)).Children(turns)
}
