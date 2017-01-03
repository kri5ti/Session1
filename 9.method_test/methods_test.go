package methods

import (
	"testing"
)

func TestPointString(t *testing.T) {
	p := Point{X: 300, Y: 60}
	got := p.String()
	want := "point: x=300, y=60"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
