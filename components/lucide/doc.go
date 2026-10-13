// Package lucide provides the Lucide icons for MyGo, as a uix component
// for the icon library of the onevra project (MIT; the icon assets stay
// Lucide's, ISC, see LICENSE). Draw one with the icon component, which
// sizes it to the font and inherits the text color:
//
//	icon.Icon(c, icon.Props{SVG: lucide.Search, Label: "Search"})
//
// Each accessor lazily parses its SVG and returns the same shared
// pointer, so the whole set costs nothing until the icons an app uses
// are drawn, and each icon costs one parse. Names follow upstream
// kebab-case converted to PascalCase (arrow-up-right becomes
// ArrowUpRight); names starting with a digit use Icon as a prefix.
// Upstream aliases and lab icons are not included. The generated file
// is pinned to a Lucide release; to refresh it, regenerate from the
// upstream SVGs.
package lucide

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

// lazyIcon is a cache holding one parsed SVG. The zero value is ready to
// use; package init roots nothing, so unused icons never parse or stay
// in memory.
type lazyIcon struct {
	once sync.Once
	svg  *ui.SVG
}

func (i *lazyIcon) get(data string) *ui.SVG {
	i.once.Do(func() { i.svg = ui.MustParseSVG([]byte(data)) })
	return i.svg
}
