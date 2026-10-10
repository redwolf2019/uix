package magicback

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			MagicBack(c, p)
		})
	}
}

// TestRenders draws the round button, alone or with its label.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{}), 320, 100)
	if len(tt.Texts()) != 0 {
		t.Fatalf("the arrow button shows text %q", tt.Texts())
	}
	tt = ui.NewTester(frame(Props{Label: "Back"}), 320, 100)
	if !tt.HasText("Back") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestBack runs OnBack while the app can go back.
func TestBack(t *testing.T) {
	back, home := 0, 0
	tt := ui.NewTester(frame(Props{
		Back:   true,
		OnBack: func() { back++ },
		OnHome: func() { home++ },
	}), 320, 100)
	tt.ClickAt(20, 20)
	if back != 1 || home != 0 {
		t.Fatalf("back %d home %d, want back 1 home 0", back, home)
	}
}

// TestHome runs OnHome while there is nowhere to go back to.
func TestHome(t *testing.T) {
	back, home := 0, 0
	tt := ui.NewTester(frame(Props{
		OnBack: func() { back++ },
		OnHome: func() { home++ },
	}), 320, 100)
	tt.ClickAt(20, 20)
	if back != 0 || home != 1 {
		t.Fatalf("back %d home %d, want back 0 home 1", back, home)
	}
}

// TestDarkMode draws the button under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Back"}), 320, 100)
	tt.SetDark(true)
	if !tt.HasText("Back") {
		t.Fatal("button missing under dark mode")
	}
}
