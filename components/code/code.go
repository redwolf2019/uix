// Package code provides retained, read-only syntax highlighting for MyGo
// text, ported from the code view of the onevra component library (MIT).
// Parse outside the frame builder and retain the document until the
// source changes. Render joins its parent's selection; call Selectable on
// the returned element for an independent code view. Applications own
// scrolling and background work.
//
//	doc := code.Parse(source, "go")
//	doc.Render(c).Padding(12)
package code

import (
	"strings"
	"unicode/utf8"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/egoist/mygo/ui"
)

// Highlighting limits bound lexer input and retained styling. Larger or
// more fragmented inputs remain readable as plain text; no source is
// truncated.
const (
	MaxHighlightBytes = 64 * 1024
	MaxHighlightRuns  = 4096
)

// Document is an immutable highlighted source, which several views can
// share. Its zero value is empty. It holds no UI handles and no global
// source cache. Both palettes are prepared once, so Render performs
// neither tokenization nor span allocation on theme changes.
type Document struct {
	source      string
	light, dark []ui.Span
}

// Parse highlights source in the named language, case-insensitively,
// accepting the aliases Chroma does (js, ts, py, sh and so on). Empty or
// unknown languages, invalid UTF-8, lexer failures and inputs beyond the
// limits use plain text. It never guesses the language, reads files or
// starts goroutines. Source bytes, tabs and line endings are preserved.
// Parsing is safe outside the UI thread; it is synchronous and makes no
// promise about its time for regex-based lexers.
func Parse(source, language string) (d *Document) {
	d = &Document{source: source}
	if source == "" || len(source) > MaxHighlightBytes || !utf8.ValidString(source) {
		return d
	}
	language = strings.ToLower(strings.TrimSpace(language))
	if language == "" || language == "text" || language == "plaintext" {
		return d
	}
	lexer := lexers.Get(language)
	if lexer == nil {
		return d
	}
	// Why a recover around the lexer: some third-party lexers panic on
	// malformed input, and a failed lexer must never make an otherwise
	// readable code block disappear.
	defer func() {
		if recover() != nil {
			d.light, d.dark = nil, nil
		}
	}()
	iterator, err := lexer.Tokenise(&chroma.TokeniseOptions{State: "root", EnsureLF: false}, source)
	if err != nil {
		return d
	}
	type run struct {
		start, end int
		kind       category
	}
	var runs []run
	offset := 0
	for token := iterator(); token != chroma.EOF; token = iterator() {
		value := token.Value
		// Lexers may append a synthetic final newline. Only original
		// bytes render.
		if offset == len(source) {
			break
		}
		if len(value) > len(source)-offset {
			value = value[:len(source)-offset]
		}
		if !strings.HasPrefix(source[offset:], value) {
			return d
		}
		if value == "" {
			continue
		}
		end, kind := offset+len(value), classify(token.Type)
		if len(runs) > 0 && runs[len(runs)-1].kind == kind {
			runs[len(runs)-1].end = end
		} else {
			if len(runs) == MaxHighlightRuns {
				return d
			}
			runs = append(runs, run{offset, end, kind})
		}
		offset = end
	}
	if offset != len(source) {
		return d
	}
	light, dark := make([]ui.Span, len(runs)), make([]ui.Span, len(runs))
	for i, run := range runs {
		light[i] = ui.Span{Text: source[run.start:run.end], Color: lightPalette[run.kind]}
		dark[i] = ui.Span{Text: source[run.start:run.end], Color: darkPalette[run.kind]}
	}
	d.light, d.dark = light, dark
	return d
}

// Text returns the exact input, including trailing newlines.
func (d *Document) Text() string {
	if d == nil {
		return ""
	}
	return d.source
}

// Highlighted reports whether tokenization completed within the limits.
func (d *Document) Highlighted() bool { return d != nil && d.light != nil }

// Render builds one wrapping monospace text element. It inherits the
// selection of its parent and adds no background or scroll container;
// apply Element methods for size, padding, NoWrap or an independent
// selection.
func (d *Document) Render(c *ui.Context) ui.Element {
	if !d.Highlighted() {
		return ui.Text(c, d.Text()).Font("monospace").MinWidth(0)
	}
	spans := d.light
	if c.Theme().Dark {
		spans = d.dark
	}
	return ui.RichText(c, spans...).Font("monospace").MinWidth(0)
}

type category uint8

const (
	plain category = iota
	keyword
	literal
	number
	comment
	name
	operator
	invalid
)

func classify(t chroma.TokenType) category {
	switch {
	case t == chroma.Error:
		return invalid
	case t.InCategory(chroma.Comment):
		return comment
	case t.InCategory(chroma.Keyword):
		return keyword
	case t.InSubCategory(chroma.LiteralString):
		return literal
	case t.InSubCategory(chroma.LiteralNumber):
		return number
	case t.InSubCategory(chroma.NameFunction), t == chroma.NameClass, t.InSubCategory(chroma.NameBuiltin), t == chroma.NameTag, t == chroma.NameAttribute:
		return name
	case t.InCategory(chroma.Operator):
		return operator
	default:
		return plain
	}
}

// Why the syntax colors stay fixed while the text around them follows the
// theme: they are content colors, not theme tokens — they tell the parts
// of the language apart in either appearance, so no token remaps them.
// Plain tokens inherit the host's text color.
var lightPalette = [...]ui.Color{{}, ui.Hex("#8b299e"), ui.Hex("#286534"), ui.Hex("#96520b"), ui.Hex("#656565"), ui.Hex("#225c9b"), ui.Hex("#845126"), ui.Hex("#b42318")}

var darkPalette = [...]ui.Color{{}, ui.Hex("#dca0eb"), ui.Hex("#a0ce95"), ui.Hex("#e5b779"), ui.Hex("#a0a0a0"), ui.Hex("#91bdf0"), ui.Hex("#e4b38b"), ui.Hex("#ff9b94")}
