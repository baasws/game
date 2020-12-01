package player

import (
	"fmt"
	"testing"

	"github.com/baasws/game/card"
	"github.com/baasws/game/errs"
	"github.com/baasws/game/seed"
)

func TestAddLenDrop(t *testing.T) {
	pc := Cards{}
	c := card.New(seed.Random(), 10)
	c2 := card.New(seed.Random(), 9)

	pc.Add(c)
	if pc.cards[0] != c {
		t.Error("Card is not card")
		return
	}
	if pc.Len() != 1 {
		t.Error("Len() should returns 1")
	}

	err := pc.Add(c)
	if fmt.Sprint(err) != errs.CardAlreadyPresent {
		t.Error("We expect an error here")
		return
	}

	if pc.Len() != 1 {
		t.Error("Len() should returns 1, before drop")
		return
	}

	err = pc.Drop(c2)
	if fmt.Sprint(err) != errs.CardNotFound {
		t.Error("We expect an error since the card is not present")
		return
	}

	if pc.Len() != 1 {
		t.Error("Len() should returns 1, after an invalid drop")
		return
	}

	pc.Add(c2)

	if pc.Len() != 2 {
		t.Error("Len() should returns 2, before drop")
		return
	}
	if !pc.Get(1).Equals(c2) {
		t.Error("card 1 should be c2")
		return
	}

	err = pc.Drop(c2)
	if err != nil {
		t.Error("no error expected here")
		return
	}

	if pc.Len() != 1 {
		t.Error("Len() should returns 1, after a valid drop")
		return
	}
}
