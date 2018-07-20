package blackjack

import (
	"github.com/SimePel/exercises/deck"
)

const (
	blackjack = 21
)

// Scoring sets value of the each card in blackjack.
// Ace = 11 is processing separately
var Scoring = map[deck.Rank]int{
	deck.Ace:   1,
	deck.Two:   2,
	deck.Three: 3,
	deck.Four:  4,
	deck.Five:  5,
	deck.Six:   6,
	deck.Seven: 7,
	deck.Eight: 8,
	deck.Nine:  9,
	deck.Ten:   10,
	deck.Jack:  10,
	deck.Queen: 10,
	deck.King:  10,
}

// Player represents TODO
type Player struct {
	DealtCards []deck.Card
}

// GetScore sums each value of your dealt card
func (p *Player) GetScore() (score int) {
	hasAce := false
	for _, c := range p.DealtCards {
		if c.Rank == deck.Ace {
			hasAce = true
		}
		score += Scoring[c.Rank]
	}
	if hasAce {
		if score+10 <= blackjack {
			score += 10
		}
	}
	return score
}
