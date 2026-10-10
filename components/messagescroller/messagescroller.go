// Package messagescroller provides a shadcn/ui-style Message Scroller:
// the scroll viewport of a conversation, which follows the new messages
// while the reader is at the end and lets the reader scroll back to
// read, then offers a button back to the newest. It scrolls the rows
// the view draws; it does not own the messages.
//
//	messagescroller.MessageScroller(c, messagescroller.Props{
//		State:    &st,
//		NewLabel: "New messages",
//		Children: func() { transcript(c) },
//	})
package messagescroller

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// State is where the scroller is, kept by the app.
type State struct {
	// Offset is where the viewport is, as the scroll container tracks
	// it: setting Offset.Y scrolls it, and math.MaxFloat32 shows the
	// end.
	Offset ui.ScrollState
	// AtEnd says the last frame ended at the newest row, so the next
	// frame follows one more. The app reads it.
	AtEnd bool
}

// Props describes the scroller to draw.
type Props struct {
	// State is where the scroller is, for the app to keep.
	State *State
	// Children draw the rows of the conversation: messages, markers,
	// separators, in order.
	Children func()
	// FollowEnd scrolls to the newest row as one appends, while the
	// reader is at the end; on by default. Scrolling up releases it.
	FollowEnd *bool
	// NewLabel labels the button that jumps to the newest row, shown
	// over the bottom while the reader is away from the end; empty
	// hides the button.
	NewLabel string
	// OnNew runs when the jump button is clicked.
	OnNew func()
}

// MessageScroller draws the viewport and returns it, so a view can
// chain more calls on it. The viewport grows to the space its container
// leaves, as a pane of a column beside a composer does.
func MessageScroller(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	follow := p.FollowEnd == nil || *p.FollowEnd
	// Follow the end while the reader is at it, as the scroll state of
	// MyGo tracks: the container clamps the request to the content, and
	// an empty scroller ends at the end from the first frame.
	if follow && st.Offset.Y >= st.Offset.MaxY {
		st.Offset.Y = math.MaxFloat32
	}
	return ui.Box(c).FillWidth().Grow(1).Children(func() {
		ui.Scroll(c).Fill().TrackScroll(&st.Offset).Children(func() {
			if p.Children != nil {
				p.Children()
			}
		})
		st.AtEnd = st.Offset.Y >= st.Offset.MaxY
		if !st.AtEnd && p.NewLabel != "" {
			// The button back to the newest row, over the bottom of the
			// viewport, where the reader looks when scrolled up.
			ui.Row(c).Absolute().Left(0).Right(0).Bottom(t.Space(2)).Justify(ui.Center).Children(func() {
				b := ui.ButtonBase(c).Padding(t.Space(1), t.Space(2.5)).Radius(999).
					Background(t.Background).Border(1, t.Border).TextColor(t.Text).
					Shadow(0, 4, 12, 0, ui.RGBA(0, 0, 0, 0.15))
				b.Children(func() { ui.Text(c, p.NewLabel).FontSize(t.FontSize * 0.85).SingleLine() })
				if b.Clicked() {
					st.Offset.Y = math.MaxFloat32
					if p.OnNew != nil {
						p.OnNew()
					}
				}
			})
		}
	})
}
