package interactivestepper

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			InteractiveStepper(c, p)
		})
	}
}

// steps builds three steps whose bodies name themselves.
func steps() []Step {
	return []Step{
		{Title: "Account", Body: func(c *ui.Context) { ui.Text(c, "Step 1 of 3") }},
		{Title: "Profile", Body: func(c *ui.Context) { ui.Text(c, "Step 2 of 3") }},
		{Title: "Done", Body: func(c *ui.Context) { ui.Text(c, "Step 3 of 3") }},
	}
}

// TestRenders draws the indicators and the first step's content.
func TestRenders(t *testing.T) {
	cur := 0
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps()}), 420, 180)
	for _, s := range []string{"Account", "Profile", "Done", "Step 1 of 3"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestNext moves to the following step.
func TestNext(t *testing.T) {
	cur := 0
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps()}), 420, 180)
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if cur != 1 || !tt.HasText("Step 2 of 3") {
		t.Fatalf("Next left the stepper at %d, texts %q", cur, tt.Texts())
	}
}

// TestBack moves to the previous step.
func TestBack(t *testing.T) {
	cur := 1
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps()}), 420, 180)
	if err := tt.Click("Back"); err != nil {
		t.Fatal(err)
	}
	if cur != 0 || !tt.HasText("Step 1 of 3") {
		t.Fatalf("Back left the stepper at %d, texts %q", cur, tt.Texts())
	}
}

// TestFinish runs OnComplete on the last step.
func TestFinish(t *testing.T) {
	cur := 2
	done := false
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps(), OnComplete: func() { done = true }}), 420, 180)
	if err := tt.Click("Finish"); err != nil {
		t.Fatal(err)
	}
	if !done {
		t.Fatal("Finish did not complete the stepper")
	}
}

// TestJumps moves to the step whose indicator is clicked.
func TestJumps(t *testing.T) {
	cur := 0
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps()}), 420, 180)
	if err := tt.Click("3"); err != nil {
		t.Fatal(err)
	}
	if cur != 2 {
		t.Fatalf("the indicator left the stepper at %d, want 2", cur)
	}
}

// TestDarkMode draws the stepper under the dark appearance too.
func TestDarkMode(t *testing.T) {
	cur := 0
	tt := ui.NewTester(frame(Props{Step: &cur, Steps: steps()}), 420, 180)
	tt.SetDark(true)
	if !tt.HasText("Step 1 of 3") {
		t.Fatal("stepper missing under dark mode")
	}
}
