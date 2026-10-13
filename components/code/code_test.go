// Package code tests the highlighted code view: what it highlights, what
// it leaves plain, and the limits that keep it readable.
package code

import (
	"strings"
	"testing"

	"github.com/egoist/mygo/ui"
)

// frameWith draws one document headless, as an app retains it across the
// frames, parsing only once.
func frameWith(d *Document) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			d.Render(c)
		})
	}
}

// TestHighlights colors the tokens of Go source and keeps the exact text.
func TestHighlights(t *testing.T) {
	src := "package main\n\nfunc main() {\n\tfmt.Println(\"hi\")\n}\n"
	d := Parse(src, "go")
	if !d.Highlighted() {
		t.Fatal("the source did not highlight")
	}
	if d.Text() != src {
		t.Fatalf("text %q, want the exact source", d.Text())
	}
	var colored bool
	for _, s := range d.light {
		if s.Color.A > 0 {
			colored = true
		}
	}
	if !colored {
		t.Fatal("no span carries a syntax color")
	}
}

// TestPlain leaves unknown, empty and unnamed languages as plain text.
func TestPlain(t *testing.T) {
	for _, in := range []struct{ src, lang string }{
		{"hello", ""},
		{"hello", "no-such-language"},
		{"hello", "text"},
	} {
		if d := Parse(in.src, in.lang); d.Highlighted() {
			t.Fatalf("%q in %q highlighted, want plain", in.src, in.lang)
		}
	}
}

// TestRenders draws the highlighted code as text.
func TestRenders(t *testing.T) {
	d := Parse("package main\nfunc main() {}", "go")
	tt := ui.NewTester(frameWith(d), 320, 100)
	if !tt.HasText("package main") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestDark draws the same document under the dark appearance too.
func TestDark(t *testing.T) {
	d := Parse("package main\nfunc main() {}", "go")
	tt := ui.NewTester(frameWith(d), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("package main") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestLimits leaves oversized or over-fragmented input as plain text, with
// every byte kept.
func TestLimits(t *testing.T) {
	big := strings.Repeat("x = 1\n", 12000) // > MaxHighlightBytes
	d := Parse(big, "python")
	if d.Highlighted() {
		t.Fatal("oversized source highlighted, want plain")
	}
	if d.Text() != big {
		t.Fatal("oversized source lost bytes")
	}
	tt := ui.NewTester(frameWith(d), 320, 100)
	if !tt.HasText("x = 1") {
		t.Fatal("oversized source did not render")
	}

	fragmented := strings.Repeat("a=1\n", 4000) // within the bytes, past the run cap
	if d := Parse(fragmented, "python"); d.Highlighted() {
		t.Fatal("over-fragmented source highlighted, want plain")
	}
}

// TestPreserves keeps the trailing newline, which a render must not trim.
func TestPreserves(t *testing.T) {
	src := "func main() {}\n"
	if got := Parse(src, "go").Text(); got != src {
		t.Fatalf("text %q, want the trailing newline kept", got)
	}
}
