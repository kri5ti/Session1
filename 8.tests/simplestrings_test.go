package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// test that Tuesday is a weekday

func Test1(t *testing.T) {
	p := Index(weekdays, "Tuesday")
	want := 7
	if p != want {
		t.Fatalf("got %d, want %d", p, want)
	}
}

// test that Sunday is not a weekday
func Test2(t *testing.T) {
	p := Index(weekdays, "Sunday")
	want := -1
	if p != want {
		t.Fatalf("got %d, want %d", p, want)
	}
}

// test that an empty search string returns 0
func Test3(t *testing.T) {
	p := Index(weekdays, "")
	want := 0
	if p != want {
		t.Fatalf("got %d, want %d", p, want)

	}
}

// test that the string Monday is not found in the empty string
func Test4(t *testing.T) {
	p := Index("", "Monday")
	want := -1
	if p != want {
		t.Fatalf("got %d, want %d", p, want)

	}
}
