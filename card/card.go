package card

import "github.com/briscola-as-a-service/game/seed"

// Card type
type Card struct {
	seed  seed.Seed
	value int
}

// New returns a Card
func New(seed seed.Seed, value int) Card {
	// check seed validity
	return Card{
		seed,
		value,
	}
}

// NewEmpty returns an empty Card
func NewEmpty() Card {
	return Card{}
}

// Equals checks cards equality
func (c Card) Equals(card Card) bool {
	return c.seed == card.seed && c.value == card.value
}

// IsBriscola returns true if card has the same seed of the one passed
func (c Card) IsBriscola(briscola Card) bool {
	return c.seed == briscola.seed
}

// Value returns the value points of a card
func (c Card) Value() int {
	return cardPoints[c.value]
}

// IsExpendable returns true if the card is expendable
func (c Card) IsExpendable() bool {
	return c.seed.IsSpade() && c.value == 2
}
