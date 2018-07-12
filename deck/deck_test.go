package deck

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cards := New()
	expected := []Card{
		Card{
			Rank: ace,
			Suit: clubs,
		},
		Card{
			Rank: two,
			Suit: clubs,
		},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{},
		Card{
			Rank: ace,
			Suit: diamonds,
		},
	}
	if cards[0] != expected[0] || cards[1] != expected[1] || cards[13] != expected[13] {
		t.Fatalf("Got:\n%v\nWant:\n%v\n", cards, expected)
	}

	deck1 := New(Shuffle())
	deck2 := New(Shuffle())
	if reflect.DeepEqual(deck1, deck2) {
		t.Fatalf("Got two the same decks")
	}
}

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
