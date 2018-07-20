package blackjack

import (
	"testing"

	"github.com/SimePel/exercises/deck"
)

func TestGetScore(t *testing.T) {
	p := Player{
		DealtCards: []deck.Card{
			deck.Card{
				Rank: deck.Two,
				Suit: deck.Spades,
			},
			deck.Card{
				Rank: deck.Jack,
				Suit: deck.Diamonds,
			},
		},
	}
	expectedScore := 12
	if p.GetScore() != expectedScore {
		t.Fatalf("Got: %v, Expected: %v\n", p.GetScore(), expectedScore)
	}

	p = Player{
		DealtCards: []deck.Card{
			deck.Card{
				Rank: deck.Ace,
				Suit: deck.Hearts,
			},
			deck.Card{
				Rank: deck.Six,
				Suit: deck.Diamonds,
			},
		},
	}
	expectedScore = 17
	if p.GetScore() != expectedScore {
		t.Fatalf("Got: %v, Expected: %v\n", p.GetScore(), expectedScore)
	}

	p = Player{
		DealtCards: []deck.Card{
			deck.Card{
				Rank: deck.Ten,
				Suit: deck.Clubs,
			},
			deck.Card{
				Rank: deck.Seven,
				Suit: deck.Diamonds,
			},
			deck.Card{
				Rank: deck.Ace,
				Suit: deck.Spades,
			},
		},
	}
	expectedScore = 18
	if p.GetScore() != expectedScore {
		t.Fatalf("Got: %v, Expected: %v\n", p.GetScore(), expectedScore)
	}
}
