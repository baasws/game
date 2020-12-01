package hand

import (
	"testing"

	"github.com/baasws/game/card"
	"github.com/baasws/game/player"
	"github.com/baasws/game/seed"
)

func TestNew(t *testing.T) {
	p := player.New("id", "name")
	c := card.New(seed.Random(), 5)
	h := New(p, c)

	if !h.card.Equals(c) {
		t.Error("card is different")
		return
	}

	if !h.player.Is(p) {
		t.Error("player is different")
	}

	// getters
	if !h.GetCard().Equals(c) {
		t.Error("card got by GetCard() is different")
		return
	}

	if !h.GetPlayer().Is(p) {
		t.Error("player got by GetPlayer() is different")
	}
}
