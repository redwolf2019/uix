package messagescroller

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// newTester draws a scroller with n rows tall enough to overflow the
// viewport, and the state the test drives.
func newTester(t *testing.T, st *State, n int) *ui.Tester {
	t.Helper()
	return ui.NewTester(func(c *ui.Context) {
		MessageScroller(c, Props{
			State:    st,
			NewLabel: "New messages",
			Children: func() {
				for i := 0; i < n; i++ {
					ui.Box(c).FillWidth().Height(48).Center().Children(func() {
						ui.Text(c, "Row")
					})
				}
			},
		})
	}, 240, 200)
}

// TestEndsAtEnd draws the scroller ending at the newest row: with a
// fresh state the first frames follow the end, so AtEnd comes true.
func TestEndsAtEnd(t *testing.T) {
	st := &State{}
	newTester(t, st, 10)
	if !st.AtEnd {
		t.Fatalf("a fresh scroller is not at the end (Y=%g MaxY=%g)", st.Offset.Y, st.Offset.MaxY)
	}
}

// TestFollows keeps the end as rows append.
func TestFollows(t *testing.T) {
	st := &State{}
	tt := newTester(t, st, 10)
	tt.Frame()
	if !st.AtEnd {
		t.Fatal("the scroller left the end as rows were drawn")
	}
}

// TestJumpsOffers the button back to the end after the reader scrolled
// up, and jumps to the end for a click on it.
func TestJumps(t *testing.T) {
	st := &State{}
	tt := newTester(t, st, 10)
	// The reader scrolled up: the next frame shows the button.
	st.Offset.Y = 0
	tt.Frame()
	if st.AtEnd {
		t.Fatal("the scroller still claims the end after scrolling up")
	}
	if !tt.HasText("New messages") {
		t.Fatalf("the jump button is missing: %q", tt.Texts())
	}
	// A click on the button jumps to the end again.
	if err := tt.Click("New messages"); err != nil {
		t.Fatal(err)
	}
	if !st.AtEnd {
		t.Fatal("the jump button did not return to the end")
	}
}

// TestNoFollow keeps the place of the reader when FollowEnd is off.
func TestNoFollow(t *testing.T) {
	st := &State{}
	follow := false
	tt := ui.NewTester(func(c *ui.Context) {
		MessageScroller(c, Props{
			State:     st,
			FollowEnd: &follow,
			Children: func() {
				for i := 0; i < 10; i++ {
					ui.Box(c).FillWidth().Height(48).Children(func() {
						ui.Text(c, "Row")
					})
				}
			},
		})
	}, 240, 200)
	tt.Frame()
	if st.Offset.Y > 0 {
		t.Fatalf("a scroller without follow moved to y=%g", st.Offset.Y)
	}
}
