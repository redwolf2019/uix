// Package settings composes an application settings surface from a
// catalog of categories, groups and items, ported from the settings view
// of the onevra component library (MIT). The host owns state and
// definitions; the component draws navigation and search around them.
// Place it in a parent with a finite height, and let it fill the parent.
// The host supplies the background or material; navigation adds a subtle
// tint.
//
//	settings.Settings(c, &state, categories, settings.Options{})
//
// Items build only while visible, and search matches their metadata, so
// applications own values, labels, persistence and effects.
package settings

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

// SettingsState is application-owned presentation state, independent of
// setting values. An empty or unknown Category selects the first
// category. Searching preserves Category.
type SettingsState struct {
	Category string
	Query    string
}

// SettingsCategory describes a navigation destination. IDs must be
// nonempty and unique among siblings; keep them stable across
// localization and rebuilds. Icon is the leading glyph of the sidebar
// item, drawn in the accent color.
type SettingsCategory struct {
	ID, Title, Section, Description string
	Icon                            *ui.SVG
	Groups                          []SettingsGroup
}

// SettingsGroup groups related items. Title may be empty for an unheaded
// group. Actions optionally builds trailing controls on the group's
// header row, such as a refresh button; like Build, it runs on the UI
// thread only when the header is visible and must not start background
// work during rendering. Actions are not searched.
type SettingsGroup struct {
	ID, Title string
	Items     []SettingsItem
	Actions   func(c *ui.Context)
}

// SettingsItem describes a searchable row. Build runs synchronously on
// the UI thread only when visible. Applications own values, accessible
// control labels, persistence and effects; Build must not start
// background work during rendering. IDs must be nonempty and unique
// within the group.
//
// DescriptionContent optionally replaces Description's text with inline
// text elements, such as Text and a Link, built into the description's
// paragraph and styled like it. Only text elements belong in it; like
// Build, it runs only when visible. Description is still what search
// matches, so give it the same words.
type SettingsItem struct {
	ID, Title, Description string
	Keywords               []string
	Build                  func(c *ui.Context)
	DescriptionContent     func(c *ui.Context)
}

// SettingsOptions controls presentation without imposing a window or
// dialog. Empty text fields use English defaults. Narrow is supplied by
// the host from its allocated width (not the window width); it replaces
// the sidebar with a category selector and stacks each row's control
// below its text.
type SettingsOptions struct {
	Narrow                                                                        bool
	Title, SearchLabel, SearchPlaceholder, CategoryLabel, ResultsTitle, EmptyText string
	// HeaderActions builds optional trailing actions, such as a dialog
	// close button.
	HeaderActions func(c *ui.Context)
}

// SettingsResult identifies an item by slice indices in the supplied
// catalog. Results are valid only for that catalog snapshot and preserve
// declaration order.
type SettingsResult struct{ Category, Group, Item int }

// SearchSettings searches every category using case-insensitive,
// whitespace-separated terms. Every term must occur in the combined
// category, group, title, description or keywords. An empty query returns
// all items. No builders run.
func SearchSettings(categories []SettingsCategory, query string) []SettingsResult {
	terms := strings.Fields(strings.ToLower(query))
	var results []SettingsResult
	for ci, category := range categories {
		for gi, group := range category.Groups {
			for ii, item := range group.Items {
				haystack := strings.ToLower(strings.Join([]string{category.Title, category.Section, category.Description, group.Title, item.Title, item.Description, strings.Join(item.Keywords, " ")}, " "))
				match := true
				for _, term := range terms {
					if !strings.Contains(haystack, term) {
						match = false
						break
					}
				}
				if match {
					results = append(results, SettingsResult{ci, gi, ii})
				}
			}
		}
	}
	return results
}

// Settings builds a bounded, scrollable settings surface using
// application-owned state and definitions. It fills its parent, and can
// be embedded in a dialog, shell content, or a standalone window. Use a
// keyed parent to distinguish multiple instances. state must not be nil.
// Navigation clears Query; clearing Query alone restores the selected
// category.
func Settings(c *ui.Context, state *SettingsState, categories []SettingsCategory, opts SettingsOptions) ui.Element {
	if state == nil {
		panic("settings: nil state")
	}
	if opts.Title == "" {
		opts.Title = "Settings"
	}
	if opts.SearchLabel == "" {
		opts.SearchLabel = "Search settings"
	}
	if opts.SearchPlaceholder == "" {
		opts.SearchPlaceholder = "Search settings"
	}
	if opts.CategoryLabel == "" {
		opts.CategoryLabel = "Category"
	}
	if opts.ResultsTitle == "" {
		opts.ResultsTitle = "Search results"
	}
	if opts.EmptyText == "" {
		opts.EmptyText = "No settings found"
	}
	selected := -1
	for i, category := range categories {
		if category.ID == state.Category {
			selected = i
			break
		}
	}
	if selected < 0 && len(categories) > 0 {
		selected = 0
		state.Category = categories[0].ID
	}
	t := c.Theme()
	search := func() {
		ui.SearchField(c, &state.Query).Label(opts.SearchLabel).Placeholder(opts.SearchPlaceholder).FillWidth().MinWidth(0)
	}
	heading := func(title string, actions bool) {
		ui.Row(c).FillWidth().MinHeight(t.Space(14)).Shrink(0).Padding(t.Space(3), t.Space(4)).Gap(t.Space(2)).Children(func() {
			ui.Text(c, title).Bold()
			ui.Spacer(c)
			if actions && opts.HeaderActions != nil {
				opts.HeaderActions(c)
			}
		})
	}
	return ui.Row(c).Fill().MinWidth(0).Gap(0).Children(func() {
		if !opts.Narrow {
			ui.Column(c).Width(t.Space(54)).FillHeight().Shrink(0).Background(t.Text.Alpha(.035)).Children(func() {
				heading(opts.Title, false)
				ui.Column(c).FillWidth().Padding(0, t.Space(3), t.Space(4), t.Space(3)).Children(search)
				ui.Scroll(c).Fill().Children(func() {
					ui.Column(c).FillWidth().MinWidth(0).Padding(0, t.Space(3), t.Space(3), t.Space(3)).Gap(t.Space(3)).Children(func() {
						for start := 0; start < len(categories); {
							section := categories[start].Section
							end := start + 1
							for end < len(categories) && categories[end].Section == section {
								end++
							}
							items := func() {
								for i := start; i < end; i++ {
									category := categories[i]
									// Why ButtonBase, not ui.Sidebar: the
									// settings rail stays always expanded
									// and each entry is one Tab stop, with
									// selection owned by the app; the
									// MyGo sidebar collapses sections and
									// moves among items with arrows.
									b := ui.ButtonBase(c).Key(category.ID).Label(category.Title).FillWidth().Justify(ui.Start).
										Padding(t.Space(1.75), t.Space(2)).Gap(t.Space(2)).Radius(t.Radius).
										Selected(selected == i).Background(ui.Color{})
									if selected == i {
										b.Background(t.Text.Alpha(.1))
									}
									b.Children(func() {
										if category.Icon != nil {
											ui.Icon(c, category.Icon).Size(t.Space(4.5), t.Space(4.5)).Shrink(0).TextColor(t.Accent)
										}
										ui.Text(c, category.Title).SingleLine()
									})
									if b.Clicked() {
										state.Category = category.ID
										selected = i
										state.Query = ""
									}
								}
							}
							if section != "" {
								ui.Column(c).FillWidth().MinWidth(0).Gap(t.Space(1)).Children(func() {
									ui.Text(c, section).FontSize(t.Rem(.8)).Bold().TextColor(t.TextMuted).
										Padding(t.Space(1), t.Space(2), t.Space(1), t.Space(2))
									items()
								})
							} else {
								ui.Column(c).FillWidth().MinWidth(0).Gap(t.Space(1)).Children(items)
							}
							start = end
						}
					})
				})
			})
			ui.Box(c).Width(1).FillHeight().Shrink(0).Background(t.Border)
		}
		ui.Column(c).Fill().MinWidth(0).Children(func() {
			title := opts.Title
			if !opts.Narrow && selected >= 0 {
				title = categories[selected].Title
			}
			if strings.TrimSpace(state.Query) != "" {
				title = opts.ResultsTitle
			}
			heading(title, true)
			if opts.Narrow {
				ui.Column(c).FillWidth().Padding(0, t.Space(3), t.Space(3), t.Space(3)).Gap(t.Space(2)).Children(func() {
					if selected >= 0 {
						sel := ui.SelectBase(c, &state.Category)
						sel.Trigger.Label(opts.CategoryLabel).FillWidth().Padding(t.Space(1.5), t.Space(3.5)).Gap(t.Space(1.5)).
							Radius(t.Radius).Justify(ui.SpaceBetween).MinWidth(t.Space(35)).
							Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
						sel.Trigger.Children(func() {
							ui.Text(c, categories[selected].Title).SingleLine()
						})
						sel.Popup(func(panel ui.Element) {
							panel.Margin(t.Space(1), 0, 0, 0).Padding(t.Space(1)).Radius(t.Radius+2).
								Background(t.Background).Border(1, t.Border).MinWidth(sel.Trigger.Bounds().W)
							panel.Shadow(0, 6, 20, 0, ui.RGBA(0, 0, 0, 0.18))
							for _, category := range categories {
								item := sel.Item(category.ID).Padding(t.Space(1.5), t.Space(2.5)).Radius(t.Radius)
								switch {
								case item.Highlighted():
									item.Background(t.Accent).TextColor(t.AccentText)
								case category.ID == state.Category:
									item.Background(t.Surface)
								}
								item.Children(func() { ui.Text(c, category.Title).SingleLine() })
							}
						})
						if sel.Trigger.Changed() {
							state.Query = ""
							for i := range categories {
								if categories[i].ID == state.Category {
									selected = i
								}
							}
						}
					}
					search()
				})
			}
			searching := strings.TrimSpace(state.Query) != ""
			results := SearchSettings(categories, state.Query)
			ui.Scroll(c).Fill().Children(func() {
				ui.Column(c).Key(state.Category).FillWidth().MinWidth(0).Padding(t.Space(5)).Gap(t.Space(5)).Children(func() {
					count, lastCategory, lastGroup := 0, -1, -1
					for _, result := range results {
						if !searching && result.Category != selected {
							continue
						}
						category := categories[result.Category]
						group := category.Groups[result.Group]
						item := group.Items[result.Item]
						if lastCategory != result.Category {
							if searching {
								ui.Text(c, category.Title)
							}
							if category.Description != "" {
								ui.Text(c, category.Description).TextColor(t.TextMuted).Wrap()
							}
						}
						if (lastCategory != result.Category || lastGroup != result.Group) && (group.Title != "" || group.Actions != nil) {
							ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
								ui.Row(c).FillWidth().MinWidth(0).AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
									ui.Text(c, group.Title).Grow(1).MinWidth(0).Wrap()
									if group.Actions != nil {
										group.Actions(c)
									}
								})
								ui.Box(c).FillWidth().Height(1).Background(t.Border)
							})
						}
						lastCategory, lastGroup = result.Category, result.Group
						count++
						var row ui.Element
						if opts.Narrow {
							row = ui.Column(c)
						} else {
							row = ui.Row(c)
						}
						row.Key(struct{ Category, Group, Item string }{category.ID, group.ID, item.ID}).FillWidth().MinWidth(0).Gap(t.Space(3)).Children(func() {
							ui.Column(c).FillWidth().Shrink(1).MinWidth(0).Gap(t.Space(1)).Children(func() {
								ui.Text(c, item.Title).Wrap()
								switch {
								case item.DescriptionContent != nil:
									ui.RichText(c).FillWidth().TextColor(t.TextMuted).FontSize(t.Rem(.875)).Wrap().Children(func() { item.DescriptionContent(c) })
								case item.Description != "":
									ui.Text(c, item.Description).TextColor(t.TextMuted).FontSize(t.Rem(.875)).Wrap()
								}
							})
							if item.Build != nil {
								item.Build(c)
							}
						})
					}
					if count == 0 {
						ui.Text(c, opts.EmptyText).TextColor(t.TextMuted).Wrap()
					}
				})
			})
		})
	})
}
