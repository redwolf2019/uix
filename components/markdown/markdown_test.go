// Package markdown tests the document renderer: the blocks it draws, the
// interactions it owns and the pieces the host owns.
package markdown

import (
	"image"
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame parses md once and draws it headless, as an app retains the
// document across the frames.
func frame(md string, opts ...Options) func(c *ui.Context) {
	d := Parse(md)
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			d.Render(c, opts...)
		})
	}
}

// TestRenders draws the common blocks: heading, emphasis, list, fenced
// code, block quote and rule.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame("# Title\n\nSome **bold** and *italic* text.\n\n- one\n- two\n\n```go\nfunc main() {}\n```\n\n> a quote\n\n---\n"), 420, 480)
	for _, s := range []string{
		"Title", "Some", "bold", "italic", "one", "two",
		"func main() {}", "a quote",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTable draws a GFM table with its header and cells.
func TestTable(t *testing.T) {
	tt := ui.NewTester(frame("| A | B |\n|---|---|\n| 1 | 2 |\n"), 420, 200)
	for _, s := range []string{"A", "B", "1", "2"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTask draws the read-only task checkboxes of a list.
func TestTask(t *testing.T) {
	tt := ui.NewTester(frame("- [x] done\n- [ ] todo\n"), 420, 140)
	for _, s := range []string{"done", "todo"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestLink hands a click on a link to the host's handler.
func TestLink(t *testing.T) {
	var got string
	tt := ui.NewTester(frame("[docs](https://mygo.egoist.dev/docs)", Options{
		OnLink: func(destination string) { got = destination },
	}), 420, 100)
	if !tt.HasText("docs") {
		t.Fatalf("no link in %q", tt.Texts())
	}
	// The link is the only inline of its paragraph, so its words start
	// at the line box's left edge; the line box, not the link's own
	// words, is what Find reports for inline text.
	r, ok := tt.Find("docs")
	if !ok {
		t.Fatal("no link line")
	}
	tt.ClickAt(r.X+5, r.Y+r.H/2)
	if got != "https://mygo.egoist.dev/docs" {
		t.Fatalf("link %q, want the destination", got)
	}
}

// TestCopyCode copies the source of a fenced block, keeping its trailing
// newline, from the header button.
func TestCopyCode(t *testing.T) {
	tt := ui.NewTester(frame("```go\nfmt.Println(\"hi\")\n```", Options{CodeHeaders: true}), 420, 160)
	r, ok := tt.Find("Copy code")
	if !ok {
		t.Fatalf("no copy button in %q", tt.Texts())
	}
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	if got := tt.Clipboard(); got != "fmt.Println(\"hi\")\n" {
		t.Fatalf("clipboard %q, want the code with its newline", got)
	}
}

// TestImage draws a paragraph of only an image at its size, named by its
// alt text.
func TestImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 40, 20))
	for x := 0; x < 40; x++ {
		for y := 0; y < 20; y++ {
			img.SetRGBA(x, y, color.RGBA{200, 100, 50, 255})
		}
	}
	opts := Options{Images: map[string]*ui.Bitmap{"pic.png": ui.NewBitmap(img)}}
	tt := ui.NewTester(frame("![a cat](pic.png)", opts), 420, 120)
	if !tt.HasText("a cat") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestDark draws the same document under the dark appearance too.
func TestDark(t *testing.T) {
	tt := ui.NewTester(frame("# Title\n\nbody\n"), 420, 200)
	tt.SetDark(true)
	for _, s := range []string{"Title", "body"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}
