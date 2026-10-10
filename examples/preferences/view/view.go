// Package view builds the preferences example: a window of MyGo native
// UI that shows the item, native select and alert dialog components of
// uix together, as a settings screen for an account. The item group
// lists the sections, the native select chooses a language, and the
// destructive item opens the alert dialog that confirms deleting the
// account. The main package shows it in a window; the tests and the
// snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/alertdialog"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/field"
	"github.com/smallnest/uix/components/item"
	"github.com/smallnest/uix/components/nativeselect"
)

// Width and Height are the size of the example window.
const Width, Height = 440, 500

// user is the media of the Account item.
var user = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>`))

// bell is the media of the Notifications item.
var bell = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"/><path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"/></svg>`))

// shield is the media of the Privacy item.
var shield = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/></svg>`))

// trash is the media of the Delete account item.
var trash = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"/><path d="M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>`))

// State is the state of the example; the controls edit it in place.
var State = PreferencesState{Language: "en"}

// PreferencesState holds the example: the language chosen, whether the
// confirm dialog shows, and the message of the last deletion.
type PreferencesState struct {
	// Language is the value chosen in the native select.
	Language string
	// ConfirmOpen shows the alert dialog.
	ConfirmOpen bool
	// Deleted is the message of the last completion of the dialog.
	Deleted string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = PreferencesState{Language: "en"} }

// PreferencesView draws the example: the heading, the sections of the
// account, the language select and the danger zone.
func PreferencesView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Preferences").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Items, a native select and an alert dialog, from uix.").TextColor(t.TextMuted)
		})
		item.ItemGroup(c, []item.Props{
			{Media: user, Title: "Account", Description: "Your name and photo", OnClick: func() {}},
			{Media: bell, Title: "Notifications", Description: "What pings you"},
			{Media: shield, Title: "Privacy", Description: "What apps may see"},
		})
		field.Field(c, field.Props{Label: "Language"}, func() {
			nativeselect.NativeSelect(c, nativeselect.Props{
				Selected: &State.Language,
				Options: []nativeselect.Option{
					{Label: "English", Value: "en"},
					{Label: "简体中文", Value: "zh"},
					{Label: "日本語", Value: "ja"},
				},
			})
		})
		item.Item(c, item.Props{
			Media:   trash,
			Title:   "Delete account…",
			Variant: item.Muted,
			OnClick: func() { State.ConfirmOpen = true },
		})
		if State.Deleted != "" {
			ui.Text(c, State.Deleted).FontSize(t.FontSize * 0.85).TextColor(t.Danger)
		}
		alertdialog.AlertDialog(c, alertdialog.Props{
			Open:        &State.ConfirmOpen,
			Title:       "Delete this account?",
			Description: "This cannot be undone. Your files will be removed, and you will lose access forever.",
			Confirm:     "Delete",
			Destructive: true,
			OnConfirm: func() {
				State.Deleted = "Account deleted"
			},
		})
	})
}
