package deck

import (
	"fmt"
	"testing"

	"github.com/briscola-as-a-service/game/seed"
)

func TestPick1(t *testing.T) {
	var d Deck

	// - this should fail. We must shuffle the deck, first
	_, err := d.Pick()
	if err == nil {
		t.Error("We should not be able to Pick on a non shuffled deck")
		return
	}

	// 	 	- cards array:
	expectedDeckSize := cardsPerSeed * len(seed.Iterable())
	if len(d.cards) != expectedDeckSize {
		t.Errorf("Wrong deck size, expected size: %d, got: %d",
			expectedDeckSize, len(d.cards))
		return
	}
	//		- passedCards array
	if len(d.passedCards) != 0 {
		t.Error("passedCards array should have zero elements")
		return
	}
	//		- availableCards array
	if len(d.availableCards) != expectedDeckSize {
		t.Error("availableCards array should have `expectedDeckSize` elements")
		return
	}
	// after the Shuffle, the deck should be ready
	if !d.ready {
		t.Error("Deck should be ready")
		return
	}
}

func TestPick2(t *testing.T) {
	d := New()
	availableCardsLenBeforePick := len(d.availableCards)

	card, err := d.Pick()
	if err != nil {
		t.Error("We are not expecing any error now")
		return
	}
	cardValue := card.Value()
	if cardValue < 1 || cardValue > cardsPerSeed {
		t.Errorf("Invalid card.Value: %v\n", cardValue)
		return
	}

	// availableCards should be 1 card less than before!
	if len(d.availableCards) != availableCardsLenBeforePick-1 {
		t.Errorf("availableCards should be %v and not %v",
			availableCardsLenBeforePick-1, len(d.availableCards))
		return
	}

	// the card picked should not be in the availableCards array
	for _, c := range d.availableCards {
		if c == card {
			t.Error("The card should not be in the d.availableCards array")
			return
		}
	}

	// the card picked should be in the passedCards array
	found := false
	for _, c := range d.passedCards {
		if c == card {
			found = true
		}
	}
	if !found {
		t.Error("The card should be in the d.passedCards array")
		return
	}

	// cards array should always be the same!
	if len(d.cards) != len(seed.Iterable())*cardsPerSeed {
		t.Error("cards array should always be the same after a Pick")
		return
	}
}

func TestPick3(t *testing.T) {
	d := New()

	deckSize := len(seed.Iterable()) * cardsPerSeed

	// let's pick all the cards
	for i := 0; i < deckSize; i++ {
		d.Pick()
	}

	// this time, we expect an error
	_, err := d.Pick()
	if err == nil {
		t.Error("We expect an error here")
		return
	}

	if len(d.availableCards) != 0 {
		t.Error("availableCards should be empty")
		return
	}

	if len(d.passedCards) != deckSize {
		t.Error("passedCards should be 'full'")
		return
	}
}

// we should not be able to Drop a deck after a Pick
func TestDrop1(t *testing.T) {
	d := New()
	d.Pick()

	err := d.Drop()
	if err == nil {
		t.Error("We expect an error here")
		return
	}
}

// we now should be able to drop a card
// That card should be the `expendable` one.
func TestDrop2(t *testing.T) {
	d := New()
	err := d.Drop()
	if err != nil {
		t.Error("We are not expecing an error here")
		return
	}

	// let's check that expendableCard is not in the availableCards but
	// not in passedCards
	if len(d.passedCards) > 0 {
		t.Error("PassedCards should be zeroed")
		return
	}

	if len(d.availableCards) != (len(seed.Iterable())*cardsPerSeed)-1 {
		t.Error("Invalid availableCards len")
		return
	}

	fmt.Println(d.availableCards)

	for _, card := range d.availableCards {
		if card.IsExpendable() {
			t.Error("This card should be dropped and not available")
			return
		}
	}
}
