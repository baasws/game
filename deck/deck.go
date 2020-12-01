package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/baasws/game/card"
	"github.com/baasws/game/seed"
)

// Deck struct
type Deck struct {
	cards          []card.Card
	availableCards []card.Card
	passedCards    []card.Card
}

const cardsPerSeed = 10

// New deck. Shuffling will reset the game
func New() (d Deck) {
	d.init()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.availableCards), func(i, j int) {
		d.availableCards[i], d.availableCards[j] =
			d.availableCards[j], d.availableCards[i]
	})
	return
}

// Pick a card from deck
func (d *Deck) Pick() (c card.Card, err error) {
	if len(d.availableCards) == 0 {
		return card.NewEmpty(), errors.New("No more cards in the deck")
	}
	// https://github.com/golang/go/wiki/SliceTricks
	// pick the last card, then remove from availables
	c = d.availableCards[len(d.availableCards)-1]
	d.passedCards = append(d.passedCards, c)
	d.burn(len(d.availableCards) - 1)
	return
}

// https://github.com/golang/go/wiki/SliceTricks#delete
func (d *Deck) burn(index int) {
	// a = append(a[:i], a[i+1:]...)
	d.availableCards = append(
		d.availableCards[:index], d.availableCards[index+1:]...)
}

// Drop a card only before the first Pick. Think of three players game
func (d *Deck) Drop() error {
	if len(d.availableCards) != len(d.cards) {
		fmt.Println("a card has already been dropped or played")
		return errors.New("Someone already picked a card")
	}
	for index, card := range d.availableCards {
		if card.IsExpendable() {
			fmt.Printf("dropping %dth card\n", index)
			d.burn(index)
		}
	}
	return nil
}

func (d *Deck) init() {
	// resetting d.cards & d.availableCards & d.passedCards
	d.cards = make([]card.Card, 0)
	d.availableCards = make([]card.Card, 0)
	d.passedCards = make([]card.Card, 0)

	// reinit deck
	for _, s := range seed.Iterable() {
		for i := 1; i < cardsPerSeed+1; i++ {
			card := card.New(s, i)
			d.cards = append(d.cards, card)
			d.availableCards = append(d.availableCards, card)
		}
	}
}
