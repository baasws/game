package card

import (
	"fmt"

	"github.com/baasws/game/seed"
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

// IsGreatherThan return true if Points() of the first card are greather than
// the given card.
func (c Card) IsGreatherThan(card Card) bool {
	return c.Points() > card.Points()
}

// IsValid return true if the card is valid
func (c Card) IsValid() bool {
	return c.value > 0 && c.value < 11 && c.seed.IsValid()
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

func (c Card) String() string {
	return fmt.Sprintf("%s%d", c.seed, c.value)
}
