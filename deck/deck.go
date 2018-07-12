package deck

import (
	"math/rand"
	"time"
)

type suit int

const (
	clubs suit = iota
	diamonds
	hearts
	spades
)

type rank int

const (
	ace rank = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
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

// New returns deck TODO
func New(opts ...func([]Card)) []Card {
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
		opt(cards)
	}

	return cards
}

// Shuffle the deck
func Shuffle() func([]Card) {
	return func(cards []Card) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}
}
