// Package view builds the fields example: a form of MyGo native UI that
// shows the field components of the shadcn/ui expansions together — the
// floating label inputs, the auto-size text area, the date-time picker
// and the multiple selector. The fields edit the state in place, and
// Save counts the saves; the page re-types nothing.
package view

import (
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/autosizetextarea"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/datetimepicker"
	"github.com/smallnest/uix/components/floatinglabelinput"
	"github.com/smallnest/uix/components/multiselect"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 540

// State is the state of the example; the controls edit it in place.
var State = FieldsState{}

// FieldsState holds the example: the fields of the form and where the
// multi-select's popup is.
type FieldsState struct {
	// Title and Email are the floating label fields.
	Title string
	Email string
	// Notes is the auto-size text area.
	Notes string
	// When is the moment the date and time fields edit.
	When time.Time
	// Guests are the chosen options of the multi-select.
	Guests []string
	// Open is whether the multi-select's popup shows.
	Open bool
	// Saved is how many saves the session made, for the tests.
	Saved int
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = FieldsState{
		When: time.Date(2026, 10, 18, 15, 0, 0, 0, time.Local),
	}
}

// FieldsView draws the example: the heading and the fields of the form.
func FieldsView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(3)).Children(func() {
		ui.Text(c, "New event").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
		floatinglabelinput.FloatingLabelInput(c, floatinglabelinput.Props{
			Label: "Title",
			Value: &State.Title,
		})
		floatinglabelinput.FloatingLabelInput(c, floatinglabelinput.Props{
			Label: "Email",
			Value: &State.Email,
		})
		ui.Column(c).FillWidth().Gap(2).Children(func() {
			ui.Text(c, "Notes").FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			autosizetextarea.AutosizeTextarea(c, autosizetextarea.Props{
				Value:       &State.Notes,
				Placeholder: "Agenda…",
				Label:       "notes",
				MinRows:     2,
				MaxRows:     5,
			})
		})
		ui.Column(c).FillWidth().Gap(2).Children(func() {
			ui.Text(c, "When").FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			datetimepicker.DateTimePicker(c, datetimepicker.Props{Value: &State.When})
		})
		multiselect.Multiselect(c, multiselect.Props{
			Values:      &State.Guests,
			Options:     []string{"Ada", "Grace", "Linus"},
			Placeholder: "Invite guests",
			Open:        &State.Open,
		})
		ui.Row(c).FillWidth().Justify(ui.End).Children(func() {
			button.Button(c, button.Props{Label: "Save", OnClick: func() { State.Saved++ }})
		})
	})
}
