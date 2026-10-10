package typewriter

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Typewriter(c, p)
		})
	}
}

// TestRenders draws the text once the typing, short enough to finish
// quickly, is done.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hi", Speed: 5 * time.Millisecond}), 320, 100)
	time.Sleep(50 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Hi") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestTypesOut reveals the text over time, and keeps it once typed.
func TestTypesOut(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hello World", Speed: 150 * time.Millisecond}), 360, 120)
	if tt.HasText("Hello World") {
		t.Fatal("the text is already typed")
	}
	time.Sleep(400 * time.Millisecond)
	tt.Frame()
	if tt.HasText("Hello World") {
		t.Fatal("the text typed too early")
	}
	time.Sleep(1500 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Hello World") {
		t.Fatal("the text never finished")
	}
	time.Sleep(300 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Hello World") {
		t.Fatal("the text did not stay")
	}
}

// TestOnComplete runs the callback once, when the typing finishes.
func TestOnComplete(t *testing.T) {
	calls := 0
	st := &State{}
	tt := ui.NewTester(frame(Props{
		Text:       "Hi",
		Speed:      10 * time.Millisecond,
		State:      st,
		OnComplete: func() { calls++ },
	}), 320, 100)
	time.Sleep(80 * time.Millisecond)
	tt.Frame()
	time.Sleep(80 * time.Millisecond)
	tt.Frame()
	if !st.Done {
		t.Fatal("the typewriter did not finish")
	}
	if calls != 1 {
		t.Fatalf("OnComplete ran %d times, want 1", calls)
	}
}

// TestDarkMode draws the typewriter under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Hi", Speed: 5 * time.Millisecond}), 320, 100)
	time.Sleep(50 * time.Millisecond)
	tt.Frame()
	tt.SetDark(true)
	if !tt.HasText("Hi") {
		t.Fatal("typewriter missing under dark mode")
	}
}
