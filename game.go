package game

import (
	"errors"
	"fmt"

	"github.com/baasws/game/card"
	"github.com/baasws/game/deck"
	"github.com/baasws/game/errs"
	"github.com/baasws/game/hand"
	"github.com/baasws/game/player"
	"github.com/baasws/game/round"
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

// New starts a new game
// (
// 	  playerCards map[player.Player]*player.Cards,
//  	briscola card.Card,
// )
func New(players []player.Player) (dk Decker) {
	//
	fmt.Println("starting..")
	dk.deck = deck.New()

	// see if we need to drop a card
	if len(players) == 3 {
		dk.deck.Drop()
	}

	// we are not expecting errors here.. Tests are ok, huh?
	dk.briscola, _ = dk.deck.Pick()
	dk.currentRound = round.New()
	dk.playerCards = make(map[player.Player]*player.Cards)
	dk.players = players
	dk.nextPlayer = players[0]
	dk.rounds = make([]round.Round, 0)

	// preparing cards for each player
	for _, p := range players {
		fmt.Printf("Picking for player %v\n", p)
		pc := player.Cards{}
		for i := 0; i < 3; i++ {
			card, _ := dk.deck.Pick()
			fmt.Printf("Picked %v\n", card)
			pc.Add(card)
		}
		dk.playerCards[p] = &pc
	}

	// finally
	return
}

// GetBriscola returns the briscola
func (dk Decker) GetBriscola() card.Card {
	return dk.briscola
}

// GetPlayerCards returns the player cards
func (dk Decker) GetPlayerCards(player player.Player) *player.Cards {
	return dk.playerCards[player]
}

// PlayCard a player plays a card
func (dk *Decker) PlayCard(player player.Player, card card.Card) (
	next player.Player,
	desk []hand.Hand,
	roundEnd bool,
	err error,
) {

	// checking if round is terminated
	if dk.isRoundEnded() {
		err = errors.New(errs.RoundEnded)
		fmt.Println("round terminated. Please call NewRound()")
		return
	}

	// checking player turn
	if player != dk.nextPlayer {
		err = errors.New(errs.NotYourTurn)
		return
	}

	// checking if card is playable by that player
	playableCard := dk.playerCards[player].Have(card)

	if !playableCard {
		err = errors.New(errs.CardNotPlayable)
		return
	}

	// appending card to desk
	hand := hand.New(player, card)
	dk.currentRound.AddHand(hand)
	desk = dk.currentRound.GetHands()

	// remove played card from playerCards
	dk.playerCards[player].Drop(card)

	// finally
	dk.setNextPlayer()
	next = dk.nextPlayer
	roundEnd = len(dk.currentRound.GetHands()) == len(dk.players)
	return
}

// NewRound only if roundEnd received on PlayCard
func (dk *Decker) NewRound() (
	next player.Player,
	playerCards map[player.Player][]card.Card,
	err error) {
	// checking if round is terminated
	if !dk.isRoundEnded() {
		err = errors.New(errs.RoundNotEnded)
		fmt.Println("round not terminated yet")
		return
	}

	// compute winner and set as nextPlayer
	dk.currentRound.ComputeWinner(dk.briscola)

	// store current round
	dk.rounds = append(dk.rounds, dk.currentRound)
	dk.nextPlayer = dk.currentRound.GetWinner()
	next = dk.nextPlayer
	dk.currentRound = round.New()

	// pick card for players
	dk.pickCardsForPlayers()
	// fmt.Printf("Players cards: %v\n", dk.playerCards)
	return
}

func (dk *Decker) pickCardsForPlayers() {
	fmt.Println("picking cards for players")
	// starting from dk.nextPlayer pick a card
	for _, p := range dk.GetSortedPlayers() {
		fmt.Printf("picking a card for player: %v\n", p)
		card, err := dk.deck.Pick()
		if err != nil {
			fmt.Println("no more cards in deck")
		} else {
			fmt.Printf("Picked %v\n", card)
			dk.playerCards[p].Add(card)
		}
	}
}

// GetSortedPlayers gives an iterable with sorted players starting by nextPlayer
func (dk Decker) GetSortedPlayers() (ret []player.Player) {
	found := false
	toAppend := []player.Player{}
	for _, p := range dk.players {
		if p.Is(dk.nextPlayer) {
			found = true
		}
		if found {
			ret = append(ret, p)
		} else {
			toAppend = append(toAppend, p)
		}
	}
	// appending other players
	if len(toAppend) > 0 {
		for _, p := range toAppend {
			ret = append(ret, p)
		}
	}
	return
}

func (dk Decker) isRoundEnded() bool {
	return len(dk.currentRound.GetHands()) == len(dk.players)
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
