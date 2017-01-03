package pointers

import "testing"

func TestPointString(t *testing.T) {
	p := Point{X: 300, Y: 60}
	p.Move(-200, 40)
	got := p.String()
	want := "point: x=100, y=100"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
