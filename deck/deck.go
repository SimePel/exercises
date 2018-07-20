package deck

import (
	"math/rand"
	"sort"
	"time"
)

type suit int

const (
	// Clubs ♣
	Clubs suit = iota
	// Diamonds ♦
	Diamonds
	// Hearts ♥
	Hearts
	// Spades ♠
	Spades
)

type rank int

const (
	// Ace 🂡
	Ace rank = iota + 1
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
	Rank rank
	Suit suit
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
				Rank: rank(j),
				Suit: suit(i),
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
func Remove(ranks ...rank) Deck {
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

// SortByRank sorts the deck by rank with provided function
func SortByRank(f func(i, j rank) bool) Deck {
	return func(cards []Card) []Card {
		sort.Slice(cards, transformRank(cards, f))
		return cards
	}
}

// SortBySuit sorts the deck by suit with provided function
func SortBySuit(f func(i, j suit) bool) Deck {
	return func(cards []Card) []Card {
		sort.Slice(cards, transformSuit(cards, f))
		return cards
	}
}

// transformSuit transforms the f to standard function for sorting
func transformSuit(cards []Card, f func(suit, suit) bool) func(int, int) bool {
	return func(i, j int) bool {
		return f(cards[i].Suit, cards[j].Suit)
	}
}

// transformRank transforms the f to standard function for sorting
func transformRank(cards []Card, f func(rank, rank) bool) func(int, int) bool {
	return func(i, j int) bool {
		return f(cards[i].Rank, cards[j].Rank)
	}
}
