package card

import (
	"github.com/briscola-as-a-service/game/seed"
	"github.com/labstack/gommon/log"
)

// Card type
type Card struct {
	seed  seed.Seed
	value int
}

// New returns a Card
func New(seed seed.Seed, value int) Card {
	if value < 1 || value > 10 {
		log.Error("invalid value for card")
		return Card{}
	}
	if !seed.IsValid() {
		log.Error("invalid seed")
		return Card{}
	}
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

// Points returns the points of a card
func (c Card) Points() int {
	return cardPoints[c.value]
}

// Value returns the card value
func (c Card) Value() int {
	return c.value
}

// Seed returns the card seed
func (c Card) Seed() seed.Seed {
	return c.seed
}

// IsExpendable returns true if the card is expendable
func (c Card) IsExpendable() bool {
	return c.seed.IsSpade() && c.value == 2
}
