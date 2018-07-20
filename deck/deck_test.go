package deck

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cards := New()
	expected := []Card{
		Card{
			Rank: Ace,
			Suit: Clubs,
		},
		Card{
			Rank: Two,
			Suit: Clubs,
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
			Rank: Ace,
			Suit: Diamonds,
		},
	}
	if cards[0] != expected[0] || cards[1] != expected[1] || cards[13] != expected[13] {
		t.Fatalf("Got:\n%v\nWant:\n%v\n", cards, expected)
	}

	deck1 := New(Shuffle())
	deck2 := New(Shuffle())
	if reflect.DeepEqual(deck1, deck2) {
		t.Fatalf("Got the same decks")
	}

	cards = New(MultiDeck(1))
	if len(cards) != 52 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(MultiDeck(3))
	if len(cards) != 156 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(Remove(Ace, Two, Three, Four))
	if len(cards) != 36 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(Remove(King, Jack), MultiDeck(2))
	if len(cards) != 88 {
		t.Fatalf("Got:\n%v\nlen=%v, cap=%v", cards, len(cards), cap(cards))
	}

	cards = New(SortByRank(func(i, j Rank) bool {
		return i < j
	}))
	if cards[0].Rank != Ace {
		t.Fatalf("Got=\n%v\nExpected that first Rank will be %v", cards, Ace.String())
	}

	cards = New(SortByRank(func(i, j Rank) bool {
		return i > j
	}))
	if cards[0].Rank != King {
		t.Fatalf("Got=\n%v\nExpected that first Rank will be %v", cards, King.String())
	}

	cards = New(SortBySuit(func(i, j Suit) bool {
		return i > j
	}))
	if cards[0].Suit != Spades {
		t.Fatalf("Got=\n%v\nExpected that first Suit will be %v", cards, Spades.String())
	}

	cards = New(SortByRank(func(i, j Rank) bool {
		return i > j
	}), SortBySuit(func(i, j Suit) bool {
		return i > j
	}))
	if cards[0].Rank != King && cards[0].Suit != Spades {
		t.Fatalf("Got=\n%v\nExpected that first card will be %v %v", cards, King.String(), Spades.String())
	}
}

func TestRankString(t *testing.T) {
	a := Ace.String()
	expected := "Ace"
	if a != expected {
		t.Fatalf("Got: %v, Expected: %v", a, expected)
	}

	var f Rank = -2
	expected = "Rank(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}

func TestSuitString(t *testing.T) {
	s := Spades.String()
	expected := "Spades"
	if s != expected {
		t.Fatalf("Got: %v, Expected: %v", s, expected)
	}

	var f Suit = -2
	expected = "Suit(-2)"
	if f.String() != expected {
		t.Fatalf("Got: %v, Expected: %v", f.String(), expected)
	}
}
