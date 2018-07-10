package deck

import "testing"

func TestRankString(t *testing.T) {
	a := ace.String()
	expected := "ace"
	if a != expected {
		t.Fatalf("Got: %v, Expected: %v", a, expected)
	}

	var f rank = -2
	expected = "rank(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}

func TestSuitString(t *testing.T) {
	s := spades.String()
	expected := "spades"
	if s != expected {
		t.Fatalf("Got: %v, Expected: %v", s, expected)
	}

	var f suit = -2
	expected = "suit(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}
