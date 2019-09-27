package round

import (
	"fmt"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/hand"
	"github.com/briscola-as-a-service/game/player"
)

// Round is an array of hands with a winner and won points
type Round struct {
	winner    player.Player
	hands     []hand.Hand
	wonPoints int
}

// New returns an empty Round
func New() Round {
	return Round{}
}

// ComputeWinner and returns the winning Hand
func (r *Round) ComputeWinner(briscola card.Card) (err error) {
	// 1: check se ci sono briscole nel round. Se sì, vince la più alta
	var winningBriscola hand.Hand
	var winningTaiet hand.Hand

	r.wonPoints = 0

	for _, hand := range r.hands {
		handValue := hand.GetCard().Points()
		fmt.Printf("handValue is %d\n", handValue)
		r.wonPoints += handValue
		fmt.Printf("WonPoints is now: %v\n", r.wonPoints)
		// briscola check
		if hand.GetCard().IsBriscola(briscola) {
			fmt.Println("Hanling with a 'briscola'")
			// se prima briscola che trovo o se è maggiore il valore della precedente
			if (winningBriscola.GetPlayer().IsEmpty()) ||
				(handValue > winningBriscola.GetCard().Points()) {
				winningBriscola = hand
			}
		} else {
			if (winningTaiet.GetPlayer().IsEmpty()) ||
				(handValue > (winningTaiet.GetCard().Points())) {
				winningTaiet = hand
			}
		}
	}
	// abbiamo il vincitore del round
	if !winningBriscola.GetPlayer().IsEmpty() {
		// winningBriscola is the winner
		r.winner = winningBriscola.GetPlayer()
		return
	}
	// 2: se non ci sono briscole, si conta solo il valore del GetCardValue()
	// winningTaiet is the winner
	r.winner = winningTaiet.GetPlayer()
	return
}

// AddHand adds and Hand to Hands
func (r *Round) AddHand(hand hand.Hand) {
	r.hands = append(r.hands, hand)
}

// GetHands returns []Hands
func (r Round) GetHands() []hand.Hand {
	return r.hands
}

// GetWinner returns the winner
func (r Round) GetWinner() player.Player {
	return r.winner
}