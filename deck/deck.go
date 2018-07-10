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
