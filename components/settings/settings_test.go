// Package settings tests the settings surface: the search it performs
// over the catalog and the rendering and navigation of the surface.
package settings

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// catalog returns a two-category fixture. record records every Build
// invocation, so tests see which items are visible. The first category
// holds two items in a headed group with an action; the second holds one.
func catalog() ([]SettingsCategory, *SettingsState, *[]string) {
	var built []string
	record := func(id string) func(c *ui.Context) {
		return func(c *ui.Context) { built = append(built, id) }
	}
	state := &SettingsState{}
	categories := []SettingsCategory{
		{
			ID: "general", Title: "General", Section: "Personal", Description: "General preferences",
			Groups: []SettingsGroup{
				{ID: "appearance", Title: "Appearance", Actions: func(c *ui.Context) { ui.Text(c, "Reset") }, Items: []SettingsItem{
					{ID: "theme", Title: "Theme", Description: "Light or dark", Keywords: []string{"dark", "light"}, Build: record("general.theme")},
					{ID: "accent", Title: "Accent", Description: "Highlight color", Build: record("general.accent")},
				}},
			},
		},
		{
			ID: "network", Title: "Network", Section: "Personal", Description: "Connection preferences",
			Groups: []SettingsGroup{
				{ID: "proxy", Title: "Proxy", Items: []SettingsItem{
					{ID: "proxy-mode", Title: "Proxy mode", Description: "How connections route", Keywords: []string{"http", "socks"}, Build: record("network.proxy")},
				}},
			},
		},
	}
	return categories, state, &built
}

// frame draws the surface over the catalog, as a host draws it over its
// own state.
func frame(categories []SettingsCategory, state *SettingsState, opts ...SettingsOptions) func(c *ui.Context) {
	return func(c *ui.Context) {
		var o SettingsOptions
		if len(opts) == 1 {
			o = opts[0]
		}
		Settings(c, state, categories, o)
	}
}

// TestSearch matches items by every term, case-insensitively, across the
// category, group, item and keywords, without running any builder.
func TestSearch(t *testing.T) {
	categories, _, _ := catalog()
	find := func(query string, want int) {
		t.Helper()
		if got := len(SearchSettings(categories, query)); got != want {
			t.Fatalf("query %q matched %d items, want %d", query, got, want)
		}
	}
	find("", 3)
	find("dark", 1)
	find("DARK", 1)
	find("network", 1)
	find("proxy route", 1)
	find("dark proxy", 0)
	find("nope", 0)
	r := SearchSettings(categories, "network")
	if r[0] != (SettingsResult{1, 0, 0}) {
		t.Fatalf("result %+v, want the network item", r[0])
	}
}

// TestRenders draws the first category and its items, the sidebar titles
// and the group's action; only the selected category's builders run.
func TestRenders(t *testing.T) {
	categories, state, built := catalog()
	tt := ui.NewTester(frame(categories, state), 640, 520)
	for _, s := range []string{"Settings", "General", "Network", "Theme", "Accent", "Light or dark", "Reset"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if !contains(*built, "general.theme") || !contains(*built, "general.accent") {
		t.Fatalf("selected category did not build: %q", *built)
	}
	if contains(*built, "network.proxy") {
		t.Fatalf("hidden category built: %q", *built)
	}
}

// TestNavigate switches the category with a sidebar click: the clicked
// category's items build, the previous category's items stop building,
// and the query clears.
func TestNavigate(t *testing.T) {
	categories, state, built := catalog()
	tt := ui.NewTester(frame(categories, state), 640, 520)
	if err := tt.Click("Network"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	// The click mutates state mid-frame, so the frame that processed it
	// still drew the previous category; measure the frame after it.
	*built = nil
	tt.Frame()
	if !contains(*built, "network.proxy") {
		t.Fatalf("clicked category did not build: %q", *built)
	}
	if contains(*built, "general.theme") {
		t.Fatalf("previous category kept building: %q", *built)
	}
	if tt.HasText("Theme") {
		t.Fatalf("previous category items still visible: %q", tt.Texts())
	}
	if state.Category != "network" || state.Query != "" {
		t.Fatalf("state after click %+v, want network with no query", *state)
	}
}

// TestSearchView searches from the start: the heading names the results,
// matching items from every category build, and others do not.
func TestSearchView(t *testing.T) {
	categories, state, built := catalog()
	state.Query = "network"
	tt := ui.NewTester(frame(categories, state), 640, 520)
	for _, s := range []string{"Search results", "Network", "Proxy mode", "How connections route"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if contains(*built, "general.theme") {
		t.Fatalf("non-matching item built: %q", *built)
	}
	if !contains(*built, "network.proxy") {
		t.Fatalf("matching item did not build: %q", *built)
	}
}

// TestEmpty shows the empty text when no item matches.
func TestEmpty(t *testing.T) {
	categories, state, _ := catalog()
	state.Query = "zzz"
	tt := ui.NewTester(frame(categories, state), 640, 520)
	if !tt.HasText("No settings found") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestNarrow replaces the sidebar with a category selector and stacks
// the items below.
func TestNarrow(t *testing.T) {
	categories, state, _ := catalog()
	tt := ui.NewTester(frame(categories, state, SettingsOptions{Narrow: true}), 420, 520)
	for _, s := range []string{"General", "Theme", "Accent"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if tt.HasText("Network") {
		t.Fatalf("other category visible in narrow view: %q", tt.Texts())
	}
}

// TestDark draws the surface under the dark appearance too.
func TestDark(t *testing.T) {
	categories, state, _ := catalog()
	tt := ui.NewTester(frame(categories, state), 640, 520)
	tt.SetDark(true)
	if !tt.HasText("Theme") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
