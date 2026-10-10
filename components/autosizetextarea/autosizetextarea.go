// Package autosizetextarea provides the AutoSize Textarea expansion of
// shadcn/ui: a text area that grows with what is typed, up to a limit,
// and then scrolls instead of growing further.
//
//	autosizetextarea.AutosizeTextarea(c, autosizetextarea.Props{
//		Value:       &notes,
//		Placeholder: "Write your notes here…",
//	})
package autosizetextarea

import "github.com/egoist/mygo/ui"

// Props describes the text area to draw.
type Props struct {
	// Value is the text being edited.
	Value *string
	// Placeholder shows in the empty field, in the muted color.
	Placeholder string
	// MinRows and MaxRows bound the height in lines of text: the field
	// grows with its text from MinRows up to MaxRows, past which it
	// scrolls. 1 and 6 unless set.
	MinRows int
	MaxRows int
	// Label names the field to assistive technology; tests and apps can
	// find it by it.
	Label string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// AutosizeTextarea draws a multi-line field that grows with *Value and
// returns it, so a view can chain more calls on it. It fills the width
// of its container.
func AutosizeTextarea(c *ui.Context, p Props) ui.Element {
	min, max := p.MinRows, p.MaxRows
	if min <= 0 {
		min = 1
	}
	if max <= 0 {
		max = 6
	}
	e := ui.TextArea(c, p.Value).Lines(min, max).FillWidth()
	e.Placeholder(p.Placeholder)
	if p.Label != "" {
		e.Label(p.Label)
	}
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
