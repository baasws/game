package game

import (
	"fmt"
	"testing"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/errs"
	"github.com/briscola-as-a-service/game/player"
)

func gimmePlayers(count int) (players []player.Player) {
	players = make([]player.Player, count)
	for i := 0; i < count; i++ {
		players[i] = player.New(fmt.Sprintf("ID-%d", i), fmt.Sprintf("pippo-%d", i))
	}
	return
}

func TestNewGame1(t *testing.T) {
	players := gimmePlayers(2)
	dk := New(players)

	for _, pc := range dk.playerCards {
		if pc.Len() != 3 {
			t.Errorf("Each player should have exactly 3 cards but we have %d",
				pc.Len())
			return
		}
	}
}

func TestNewGame2ThreePlayers(t *testing.T) {
	players := gimmePlayers(3)
	dk := New(players)

	for _, pc := range dk.playerCards {
		if pc.Len() != 3 {
			t.Errorf("Each player should have exactly 3 cards but we have %d",
				pc.Len())
			return
		}
	}
}

func TestNewGame2(t *testing.T) {
	players := gimmePlayers(2)
	dk := New(players)

	for i, p := range players {
		if dk.players[i] != p {
			t.Errorf("Player #%d is not the same", i)
			return
		}
	}

	if dk.nextPlayer != players[0] {
		// check nextPlayer
		t.Errorf("nextPlayer should be %v", players[0])
		return
	}
}

func TestPlayCard(t *testing.T) {
	players := gimmePlayers(2)
	dk := New(players)
	playerCards := dk.GetPlayerCards(players[0])

	// should emit an error if we are playing as the wrong player
	_, _, _, err := dk.PlayCard(players[1], card.NewEmpty())
	if fmt.Sprint(err) != errs.NotYourTurn {
		t.Error("We are expecting an error here")
		return
	}

	// should emit an error if the card is not one the player can play
	_, _, _, err = dk.PlayCard(players[0], card.NewEmpty())
	if fmt.Sprint(err) != errs.CardNotPlayable {
		t.Error("We are expecting an error here")
		return
	}

	card := playerCards.Get(0)
	next, desk, roundEnd, err := dk.PlayCard(players[0], card)
	if err != nil {
		t.Error("no error expected")
		return
	}

	if len(desk) != 1 {
		t.Error("desk len should be 1")
		return
	}

	if !desk[0].GetCard().Equals(card) {
		t.Error("desk card is different")
		return
	}

	if !desk[0].GetPlayer().Is(players[0]) {
		t.Error("desk player is different")
		return
	}

	if !next.Is(players[1]) {
		t.Error("next player should be players[1]")
		return
	}

	if roundEnd {
		t.Error("round should not be at the end")
		return
	}

	// trying to move to the next round, should give an error
	_, _, err = dk.NewRound()
	if fmt.Sprint(err) != errs.RoundNotEnded {
		t.Error("We expect an error here")
		return
	}

	// play another card as same user, we expect an error if setNextPlayer
	// correctly did his job
	_, _, _, err = dk.PlayCard(players[0], card)
	if fmt.Sprint(err) != errs.NotYourTurn {
		t.Error("We expect an error")
		return
	}

	// closing round
	card2 := dk.GetPlayerCards(players[1]).Get(0)
	next, desk, roundEnd, err = dk.PlayCard(players[1], card2)
	if err != nil {
		t.Error("No error expected")
		return
	}
	if len(desk) != 2 {
		t.Error("Invalid desk len")
		return
	}
	if !desk[0].GetCard().Equals(card) ||
		!desk[1].GetCard().Equals(card2) {
		t.Error("Invalid cards in deck")
		return
	}

	if next != players[0] {
		t.Error("invalid next player")
		return
	}

	if !roundEnd {
		t.Error("Round should be terminated!")
		return
	}

	// playing on an ended round should fail
	_, _, _, err = dk.PlayCard(players[0], card)
	if fmt.Sprint(err) != errs.RoundEnded {
		t.Error("We expected an error here")
		return
	}

	// compute round and see what happens
	next, _, err = dk.NewRound()
	if err != nil {
		t.Error("No error expected here")
		return
	}
	if next.IsEmpty() {
		t.Error("next player should be the winner")
		return
	}
}
