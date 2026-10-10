// Package alertdialog provides a shadcn/ui-style Alert Dialog: a modal
// that interrupts the app for a confirmation, as deleting something
// permanent, and expects a response. A media icon may sit above the
// title, and the dialog comes in a compact size. The action button
// closes it, in the accent or the danger color for a destructive
// action; the backdrop and Escape close it without.
//
//	alertdialog.AlertDialog(c, alertdialog.Props{
//		Open:        &open,
//		Title:       "Delete this account?",
//		Description: "This cannot be undone.",
//		Confirm:     "Delete",
//		Destructive: true,
//		OnConfirm:   app.deleteAccount,
//	})
package alertdialog

import "github.com/egoist/mygo/ui"

// Props describes the alert dialog to draw.
type Props struct {
	// Open is the bool that shows and hides the dialog; the backdrop
	// and Escape set it to false.
	Open *bool
	// Title is the heading of the dialog.
	Title string
	// Description is the text under the title.
	Description string
	// Media is the shape above the title, an SVG in the accent color;
	// nil for none.
	Media *ui.SVG
	// Confirm labels the action button; an empty label hides it.
	Confirm string
	// Cancel labels the cancel button; "Cancel" when empty.
	Cancel string
	// Destructive colors the action button with the danger color.
	Destructive bool
	// Compact draws the dialog at the smaller size.
	Compact bool
	// OnConfirm runs for the action button, which also closes the
	// dialog.
	OnConfirm func()
	// OnCancel runs for the cancel button; the button closes the dialog
	// either way.
	OnCancel func()
}

// AlertDialog draws the dialog for props, or nothing while it is closed,
// and returns the panel. The backdrop dims the window; clicking it or
// pressing Escape closes the dialog.
func AlertDialog(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	cancel := p.Cancel
	if cancel == "" {
		cancel = "Cancel"
	}
	width := t.Space(48)
	pad := t.Space(5)
	if p.Compact {
		width = t.Space(40)
		pad = t.Space(4)
	}
	return ui.DialogBase(c, p.Open, func(back, panel ui.Element) {
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		panel.Padding(pad).Gap(t.Space(3)).Radius(t.Radius+4).Background(t.Background).
			MinWidth(width).MaxWidth(width).Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		panel.Children(func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				if p.Media != nil {
					// The icon in a tinted circle, centered above the
					// text, as an alert with a media head does.
					ui.Row(c).Justify(ui.Center).Children(func() {
						ui.Box(c).Size(t.Space(9), t.Space(9)).Radius(t.Radius).
							Background(t.Accent.Alpha(0.15)).Center().Children(func() {
							ui.Icon(c, p.Media).Size(t.Space(4.5), t.Space(4.5)).TextColor(t.Accent)
						})
					})
				}
				ui.Column(c).Gap(t.Space(1.5)).Children(func() {
					ui.Text(c, p.Title).Bold()
					if p.Description != "" {
						ui.Text(c, p.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
					}
				})
				ui.Row(c).Justify(ui.End).Gap(t.Space(1.5)).Children(func() {
					cancelButton(c, t, cancel, p)
					if p.Confirm != "" {
						confirmButton(c, t, p)
					}
				})
			})
		})
	})
}

// cancelButton draws the cancel button, which closes the dialog. Why not
// a hover face: the pointer is already inside the modal, so a solid face
// keeps the dialog calm, as the dialog of uix does.
func cancelButton(c *ui.Context, t *ui.Theme, label string, p Props) {
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	b.Children(func() { ui.Text(c, label).SingleLine() })
	if b.Clicked() {
		*p.Open = false
		if p.OnCancel != nil {
			p.OnCancel()
		}
	}
}

// confirmButton draws the action button, in the accent or the danger
// color, which closes the dialog and runs OnConfirm.
func confirmButton(c *ui.Context, t *ui.Theme, p Props) {
	face, fg := t.Accent, t.AccentText
	if p.Destructive {
		face, fg = t.Danger, ui.Hex("#ffffff")
	}
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).Background(face).TextColor(fg)
	b.Children(func() { ui.Text(c, p.Confirm).SingleLine() })
	if b.Clicked() {
		*p.Open = false
		if p.OnConfirm != nil {
			p.OnConfirm()
		}
	}
}
