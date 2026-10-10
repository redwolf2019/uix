// Package typewriter provides the Typewriter expansion of shadcn/ui: a
// text that types itself out, a character at a time, with a blinking
// caret. A native app uses it to make a headline arrive, or a message
// as if someone is typing it.
//
//	typewriter.Typewriter(c, typewriter.Props{Text: "Hello from MyGo"})
package typewriter

import (
	"time"
	"unicode/utf8"

	"github.com/egoist/mygo/ui"
)

// State tracks the typing of a typewriter, so the view can be told when
// it completes. Pass it in Props to receive the OnComplete call and to
// let the typewriter stop animating once it is done.
type State struct {
	// Done is true once the text is fully revealed; it stays true while
	// the typewriter is looping, until it starts over.
	Done bool
}

// Props describes the text to type.
type Props struct {
	// Text is what the typewriter types.
	Text string
	// Speed is the time for one character; 33ms unless set.
	Speed time.Duration
	// Loop replays the text from the start, once typed; false types it
	// once and keeps it.
	Loop bool
	// Cursor shows a blinking caret after the text; on unless set.
	Cursor *bool
	// State is where the typewriter keeps its progress. Without it, the
	// text still types once; pass it to be told when that completes
	// (OnComplete) and to let the typewriter rest when done.
	State *State
	// OnComplete runs once, when the text is fully revealed; it needs
	// State, so it runs once, not every frame.
	OnComplete func()
}

// Typewriter draws the revealed text with the caret and returns the row,
// so a view can chain more calls on it.
func Typewriter(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	speed := p.Speed
	if speed <= 0 {
		speed = 33 * time.Millisecond
	}
	cursor := p.Cursor == nil || *p.Cursor
	n := utf8.RuneCountInString(p.Text)
	total := time.Duration(max(n, 1)) * speed

	e := ui.Row(c).FillWidth().Gap(1).AlignItems(ui.Center)
	e.Children(func() {
		var progress float32
		switch {
		case p.Loop:
			// The loop re-types the text forever, as a marquee does.
			progress = e.Loop("type", total, ui.Linear)
		case p.State != nil && p.State.Done:
			progress = 1
		default:
			// Why a Loop that hands the animation to AnimateWith: the
			// Loop starts the reveal from 0 on the first frame, and
			// AnimateWith then moves it to 1 and keeps it there, where a
			// Loop alone would wrap around and type the text again.
			e.Loop("type", total, ui.Linear)
			progress = e.AnimateWith("type", 1, total, ui.Linear)
		}
		reveal := int(progress * float32(n))
		if reveal > n {
			reveal = n
		}
		shown := string([]rune(p.Text)[:reveal])
		ui.Text(c, shown)
		if cursor {
			// The caret blinks on its own Loop, a square on and off.
			blink := e.Loop("cursor", 600*time.Millisecond, ui.Linear)
			on := int(blink*2)%2 == 0
			ui.Box(c).Size(2, t.FontSize).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
				if on {
					p.Fill(r, t.Text, 0)
				}
			})
		}
		if reveal == n && p.State != nil && !p.State.Done {
			p.State.Done = true
			if p.OnComplete != nil {
				p.OnComplete()
			}
		}
	})
	return e
}
