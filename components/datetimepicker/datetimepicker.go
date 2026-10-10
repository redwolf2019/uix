// Package datetimepicker provides the DateTime Picker expansion of
// shadcn/ui: a date field and a time field editing one value together.
//
//	datetimepicker.DateTimePicker(c, datetimepicker.Props{Value: &when})
package datetimepicker

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the picker to draw.
type Props struct {
	// Value is the moment being edited: the date field keeps the time of
	// day, the time field the date.
	Value *time.Time
	// Disabled keeps the fields from being edited.
	Disabled bool
}

// DateTimePicker draws the date and the time side by side, editing
// *Value, and returns the row, so a view can chain more calls on it.
func DateTimePicker(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	return ui.Row(c).FillWidth().Gap(t.Space(2)).Children(func() {
		date := ui.DateInput(c, p.Value).Grow(1)
		time := ui.TimeInput(c, p.Value)
		if p.Disabled {
			date.Disabled(true)
			time.Disabled(true)
		}
	})
}
