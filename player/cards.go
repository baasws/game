package player

import (
	"errors"

	"github.com/baasws/game/card"
	"github.com/baasws/game/errs"
)

// Cards matains the card a Player have in hands
type Cards struct {
	cards []card.Card
}

// Add a card to cards
func (p *Cards) Add(card card.Card) error {
	if p.Have(card) {
		return errors.New(errs.CardAlreadyPresent)
	}
	p.cards = append(p.cards, card)
	return nil
}

// Get returns the index-th card
func (p Cards) Get(index int) card.Card {
	return p.cards[index]
}

// Drop a card from Player's hand
func (p *Cards) Drop(card card.Card) error {
	for i, c := range p.cards {
		if c == card {
			p.cards = append(p.cards[:i], p.cards[i+1:]...)
			return nil
		}
	}
	return errors.New(errs.CardNotFound)
}

// Len returns the len of cards
func (p Cards) Len() int {
	return len(p.cards)
}

// Have that card?
func (p Cards) Have(card card.Card) bool {
	for _, c := range p.cards {
		if c == card {
			return true
		}
	}
	return false
}
