package infinitescroll

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestLoadsMore loads the next page as the list nears its end, once per
// visit, and not while the rows fit.
func TestLoadsMore(t *testing.T) {
	st := &State{}
	count := 0
	rows := 50
	view := func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			InfiniteScroll(c, Props{
				State: st,
				OnMore: func() {
					count++
					rows += 20
				},
				Children: func() {
					for i := 0; i < rows; i++ {
						ui.Box(c).FillWidth().Height(32)
					}
				},
			})
		})
	}
	tt := ui.NewTester(view, 300, 200)
	if count != 0 {
		t.Fatalf("loaded %d pages before the end", count)
	}
	st.Scroll.Y = math.MaxFloat32
	tt.Frame()
	if count != 1 {
		t.Fatalf("the end loaded %d pages, want 1", count)
	}
	tt.Frame()
	if count != 1 {
		t.Fatalf("the end loaded %d pages again, want 1", count)
	}
	st.Scroll.Y = 0
	tt.Frame()
	st.Scroll.Y = math.MaxFloat32
	tt.Frame()
	if count != 2 {
		t.Fatalf("the end loaded %d pages in all, want 2", count)
	}
}
