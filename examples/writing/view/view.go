// Package view builds the writing example: a page of prose of MyGo
// native UI that shows the writing components of the shadcn/ui
// expansions together — the heading with its anchor, the typewriter and
// the block quote. Clicking the anchor of the heading copies
// "uix#getting-started" to the clipboard, as a document's headings
// deep-link to their place; the typewriter types its line once, with
// the caret blinking until it finishes.
package view

import (
	"time"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/blockquote"
	"github.com/smallnest/uix/components/headingwithanchor"
	"github.com/smallnest/uix/components/inlinecode"
	"github.com/smallnest/uix/components/typewriter"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 420

// State is the state of the example; the controls edit it in place.
var State = WritingState{}

// WritingState holds the example: the typewriter's progress and the
// last anchor copied.
type WritingState struct {
	// Typing is where the typewriter keeps its progress; the view
	// resets it, so the line types again.
	Typing typewriter.State
	// Copied is the last anchor the heading copied, for the tests.
	Copied string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = WritingState{}
}

// WritingView draws the example: the heading with its anchor, the
// typewriter line and the block quote.
func WritingView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(3)).Children(func() {
		headingwithanchor.Heading(c, headingwithanchor.Props{
			Text:     "Getting Started",
			Anchor:   "getting-started",
			OnAnchor: func() { State.Copied = "uix#getting-started" },
		})
		ui.Text(c, "The uix registry brings the shadcn/ui components to MyGo, a native toolkit: the same roles, drawn without HTML.")
		ui.Row(c).FillWidth().Wrap().AlignItems(ui.Center).Gap(4).Children(func() {
			ui.Text(c, "Run")
			inlinecode.InlineCode(c, inlinecode.Props{Text: "uix add button"})
			ui.Text(c, "to add a component to your app.")
		})
		ui.Column(c).FillWidth().Padding(t.Space(2), 0).Children(func() {
			typewriter.Typewriter(c, typewriter.Props{
				Text:  "Build native apps with uix.",
				Speed: 30 * time.Millisecond,
				State: &State.Typing,
			})
		})
		blockquote.Blockquote(c, blockquote.Props{
			Text:   "Simplicity is the soul of efficiency.",
			Author: "Austin Freeman",
		})
	})
}
