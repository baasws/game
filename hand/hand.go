package hand

import (
	"github.com/baasws/game/card"
	"github.com/baasws/game/player"
)

// Hand is a Player that played a Card
type Hand struct {
	player player.Player
	card   card.Card
}

// New returns a new Hand
func New(player player.Player, card card.Card) Hand {
	return Hand{
		player: player,
		card:   card,
	}
}

// GetCard returns the Card
func (h Hand) GetCard() card.Card {
	return h.card
}

// GetPlayer returns the Player
func (h Hand) GetPlayer() player.Player {
	return h.player
}
