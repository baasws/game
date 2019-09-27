package round

import (
	"fmt"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/hand"
	"github.com/briscola-as-a-service/game/player"
)

// Round is an array of hands with a winner and won points
type Round struct {
	Winner    player.Player
	Hands     []hand.Hand
	WonPoints int
}

// ComputeWinner and returns the winning Hand
func (r *Round) ComputeWinner(briscola card.Card) (err error) {
	// 1: check se ci sono briscole nel round. Se sì, vince la più alta
	var winningBriscola hand.Hand
	var winningTaiet hand.Hand

	r.WonPoints = 0

	for _, hand := range r.Hands {
		handValue := hand.GetCard().Value()
		fmt.Printf("handValue is %d\n", handValue)
		r.WonPoints += handValue
		fmt.Printf("WonPoints is now: %v\n", r.WonPoints)
		// briscola check
		if hand.GetCard().IsBriscola(briscola) {
			fmt.Println("Hanling with a 'briscola'")
			// se prima briscola che trovo o se è maggiore il valore della precedente
			if (winningBriscola.GetPlayer().IsEmpty()) ||
				(handValue > winningBriscola.GetCard().Value()) {
				winningBriscola = hand
			}
		} else {
			if (winningTaiet.GetPlayer().IsEmpty()) ||
				(handValue > (winningTaiet.GetCard().Value())) {
				winningTaiet = hand
			}
		}
	}
	// abbiamo il vincitore del round
	if !winningBriscola.GetPlayer().IsEmpty() {
		// winningBriscola is the winner
		r.Winner = winningBriscola.GetPlayer()
		return
	}
	// 2: se non ci sono briscole, si conta solo il valore del GetCardValue()
	// winningTaiet is the winner
	r.Winner = winningTaiet.GetPlayer()
	return
}

// AddHand adds and Hand to Hands
func (r *Round) AddHand(hand hand.Hand) {
	r.Hands = append(r.Hands, hand)
}
