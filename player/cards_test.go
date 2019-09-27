package player

import (
	"fmt"
	"testing"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/errs"
	"github.com/briscola-as-a-service/game/seed"
)

func TestAdd(t *testing.T) {
	pc := Cards{}
	card := card.New(seed.Random(), 10)

	pc.Add(card)
	if pc.cards[0] != card {
		t.Error("Card is not card")
		return
	}

	err := pc.Add(card)
	if fmt.Sprint(err) != errs.CardAlreadyPresent {
		t.Error("We expect an error here")
		return
	}
}
