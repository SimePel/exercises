package deck

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
	cardsForOneDeck = 54
	cardsForOneSuit = 13
	allSuits        = 4
)

// New returns deck TODO
func New() []Card {
	cards := make([]Card, 0, cardsForOneDeck)
	for i := 0; i < allSuits; i++ {
		for j := 1; j <= cardsForOneSuit; j++ {
			cards = append(cards, Card{
				Rank: rank(j),
				Suit: suit(i),
			})
		}
	}
	return cards
}
