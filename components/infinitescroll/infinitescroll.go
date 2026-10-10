// Package infinitescroll provides the Infinite Scroll expansion of
// shadcn/ui: a list that loads the next page of rows as its end comes
// near, so scrolling never runs out, as a feed or a log of a long
// stream does.
//
//	infinitescroll.InfiniteScroll(c, infinitescroll.Props{
//		State:  &app.feed,
//		OnMore: loadMore,
//		Children: app.rows,
//	})
package infinitescroll

import "github.com/egoist/mygo/ui"

// State tracks the list of an InfiniteScroll; the view keeps one and
// passes it in, as a feed keeps where it is scrolled to.
type State struct {
	// Scroll is the position of the list; the view can set Y to scroll
	// it, as for any scroll container.
	Scroll ui.ScrollState
	// Fired latches the trigger, so OnMore runs once per visit to the
	// end; the component resets it as the list scrolls away.
	Fired bool
}

// Props describes the list to draw.
type Props struct {
	// State is where the list keeps its scroll and its trigger; the view
	// owns it, as a feed owns where it is.
	State *State
	// OnMore loads the next page of rows; it runs when the list nears
	// its end and the trigger has not fired since.
	OnMore func()
	// Children draws the rows; the view appends the next page of rows
	// in OnMore, and the list shows them as they are added.
	Children func()
}

// InfiniteScroll draws the scrolling list and returns it, so a view can
// chain more calls on it. It fills its container.
func InfiniteScroll(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	s := ui.Scroll(c).Fill().TrackScroll(&p.State.Scroll)
	s.Children(p.Children)
	// Near the end, load the next page, once until the list scrolls away
	// and back. MaxY is the height of the last frame; zero means the
	// rows fit the list, and nothing is loaded yet.
	if p.State.Scroll.MaxY > 0 {
		near := p.State.Scroll.Y >= p.State.Scroll.MaxY-t.Space(8)
		switch {
		case near && !p.State.Fired:
			p.State.Fired = true
			if p.OnMore != nil {
				p.OnMore()
			}
		case !near:
			p.State.Fired = false
		}
	}
	return s
}
