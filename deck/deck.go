package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Suit of the card
type Suit int

const (
	// Clubs ♣
	Clubs Suit = iota
	// Diamonds ♦
	Diamonds
	// Hearts ♥
	Hearts
	// Spades ♠
	Spades
)

// Rank of the card
type Rank int

const (
	// Ace 🂡
	Ace Rank = iota + 1
	// Two 🂢
	Two
	// Three 🂣
	Three
	// Four 🂤
	Four
	// Five 🂥
	Five
	// Six 🂦
	Six
	// Seven 🂧
	Seven
	// Eight 🂨
	Eight
	// Nine 🂩
	Nine
	// Ten 🂪
	Ten
	// Jack 🂫
	Jack
	// Queen 🂭
	Queen
	// King 🂮
	King
)

// Card is representing abstract playing card
type Card struct {
	Rank Rank
	Suit Suit
}

const (
	cardsForOneDeck = 52
	cardsForOneSuit = 13
	allSuits        = 4
)

// Deck is set of cards
type Deck func([]Card) []Card

// New returns deck with applied opts
func New(opts ...Deck) []Card {
	cards := make([]Card, 0, cardsForOneDeck)
	for i := 0; i < allSuits; i++ {
		for j := 1; j <= cardsForOneSuit; j++ {
			cards = append(cards, Card{
				Rank: Rank(j),
				Suit: Suit(i),
			})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// Shuffle the deck
func Shuffle() Deck {
	return func(cards []Card) []Card {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

// MultiDeck appends q decks to the original one
func MultiDeck(q int) Deck {
	if q <= 1 {
		return func(cards []Card) []Card { return cards }
	}
	return func(cards []Card) []Card {
		tmp := make([]Card, len(cards), q*len(cards))
		copy(tmp, cards)
		for q != 1 {
			tmp = append(tmp, cards...)
			q--
		}
		return tmp
	}
}

// Remove specified cards from the deck
func Remove(ranks ...Rank) Deck {
	return func(cards []Card) []Card {
		tmp := make([]Card, 0, len(cards))
		var needToAdd bool
		var newLen int
		for _, c := range cards {
			needToAdd = true
			for _, r := range ranks {
				if c.Rank == r {
					needToAdd = false
				}
			}
			if needToAdd {
				tmp = append(tmp, c)
				newLen++
			}
		}
		res := make([]Card, newLen)
		copy(res, tmp)
		return res
	}
}

// SortByRank sorts the deck by Rank with provided function
func SortByRank(f func(i, j Rank) bool) Deck {
	return func(cards []Card) []Card {
		sort.Slice(cards, transformRank(cards, f))
		return cards
	}
}

// transformRank transforms the f to standard function for sorting
func transformRank(cards []Card, f func(Rank, Rank) bool) func(int, int) bool {
	return func(i, j int) bool {
		return f(cards[i].Rank, cards[j].Rank)
	}
}
