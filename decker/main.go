package decker

import (
	"errors"
	"fmt"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/player"
	"github.com/briscola-as-a-service/game/round"
)

// Decker is the Deck handler
type Decker struct {
	playerCards  map[player.Player]*player.Cards
	briscola     card.Card
	rounds       []round.Round
	currentRound round.Round
	players      []player.Player
	nextPlayer   player.Player
	deck         deck.Deck
}

// NewGame starts a new game
func (dk *Decker) NewGame(players []Player) (
	playerCards map[Player]*PlayerCards,
	briscola deck.Card,
) {
	//
	fmt.Println("starting..")
	dk.deck = deck.Deck{}

	dk.deck.Shuffle()

	// see if we need to drop a card
	if len(players) == 3 {
		dk.deck.Drop()
	}

	// we are not expecting errors here.. Tests are ok, huh?
	dk.briscola, _ = dk.deck.Pick()
	dk.currentRound = Round{}
	dk.playerCards = make(map[Player]*PlayerCards)
	dk.players = players
	dk.nextPlayer = players[0]
	dk.rounds = make([]Round, 0)

	// preparing cards for each player
	for _, p := range players {
		fmt.Printf("Picking for player %v\n", p)
		pc := PlayerCards{}
		for i := 0; i < 3; i++ {
			card, _ := dk.deck.Pick()
			fmt.Printf("Picked %v\n", card)
			pc.Add(card)
		}
		dk.playerCards[p] = &pc
	}

	// finally
	briscola = dk.briscola
	playerCards = dk.playerCards
	return
}

// PlayCard a player plays a card
func (dk *Decker) PlayCard(player Player, card deck.Card) (
	next Player,
	desk []Hand,
	roundEnd bool,
	err error,
) {

	// checking if round is terminated
	if dk.isRoundEnded() {
		err = errors.New(ErrRoundEnded)
		fmt.Println("round terminated. Please call NextRound()")
		return
	}

	// checking player turn
	if player != dk.nextPlayer {
		err = errors.New(ErrNotYourTurn)
		return
	}

	// checking if card is playable by that player
	playableCard := dk.playerCards[player].Have(card)

	if !playableCard {
		err = errors.New(ErrCardNotPlayable)
		return
	}

	// appending card to desk
	hand := Hand{
		Card:   card,
		Player: player,
	}
	dk.currentRound.AddHand(hand)
	desk = dk.currentRound.Hands

	// remove played card from playerCards
	dk.playerCards[player].Drop(card)

	// finally
	dk.setNextPlayer()
	next = dk.nextPlayer
	roundEnd = len(dk.currentRound.Hands) == len(dk.players)
	return
}

// NewRound only if roundEnd received on PlayCard
func (dk *Decker) NewRound() (
	next Player,
	playerCards map[Player][]deck.Card,
	err error) {
	// checking if round is terminated
	if !dk.isRoundEnded() {
		err = errors.New(ErrRoundNotEnded)
		fmt.Println("round not terminated yet")
		return
	}

	// compute winner and set as nextPlayer
	dk.currentRound.ComputeWinner(dk.briscola, dk.deck)

	// store current round
	dk.rounds = append(dk.rounds, dk.currentRound)
	dk.nextPlayer = dk.currentRound.Winner
	next = dk.nextPlayer
	dk.currentRound = Round{}

	// pick card for players

	return
}

func (dk Decker) isRoundEnded() bool {
	return len(dk.currentRound.Hands) == len(dk.players)
}

func (dk *Decker) setNextPlayer() {
	nextIndex := -1
	for i, p := range dk.players {
		if p == dk.nextPlayer {
			nextIndex = i + 1
		}
	}
	if nextIndex >= len(dk.players) {
		nextIndex = 0
	}
	dk.nextPlayer = dk.players[nextIndex]
}
