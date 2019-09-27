package decker

import (
	"fmt"
	"testing"
)

func gimmePlayers(count int) (players []Player) {
	players = make([]Player, count)
	for i := 0; i < count; i++ {
		players[i] = Player{
			Name: fmt.Sprintf("pippo-%d", i),
			ID:   fmt.Sprintf("ID-%d", i),
		}
	}
	return
}

func TestNewGame1(t *testing.T) {
	var dk Decker
	players := gimmePlayers(2)
	dk.NewGame(players)

	for _, pc := range dk.playerCards {
		if pc.Len() != 3 {
			t.Errorf("Each player should have exactly 3 cards but we have %d",
				pc.Len())
			return
		}
	}
}

func TestNewGame2ThreePlayers(t *testing.T) {
	var dk Decker
	players := gimmePlayers(3)
	dk.NewGame(players)

	for _, pc := range dk.playerCards {
		if pc.Len() != 3 {
			t.Errorf("Each player should have exactly 3 cards but we have %d",
				pc.Len())
			return
		}
	}
}

func TestNewGame2(t *testing.T) {
	var dk Decker
	players := gimmePlayers(2)
	dk.NewGame(players)

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
	var dk Decker
	players := gimmePlayers(2)
	playerCards, _ := dk.NewGame(players)

	// should emit an error if we are playing as the wrong player
	_, _, _, err := dk.PlayCard(players[1], deck.Card{})
	if fmt.Sprint(err) != ErrNotYourTurn {
		t.Error("We are expecting an error here")
		return
	}

	// should emit an error if the card is not one the player can play
	_, _, _, err = dk.PlayCard(players[0], deck.Card{})
	if fmt.Sprint(err) != ErrCardNotPlayable {
		t.Error("We are expecting an error here")
		return
	}

	card := playerCards[players[0]].cards[0]
	next, desk, roundEnd, err := dk.PlayCard(players[0], card)
	if err != nil {
		t.Error("no error expected")
		return
	}

	if len(desk) != 1 {
		t.Error("desk len should be 1")
		return
	}

	if desk[0].Card != card {
		t.Error("desk card is different")
		return
	}

	if desk[0].Player != players[0] {
		t.Error("desk player is different")
		return
	}

	if next != players[1] {
		t.Error("next player should be players[1]")
		return
	}

	if roundEnd {
		t.Error("round should not be at the end")
		return
	}

	// trying to move to the next round, should give an error
	_, _, err = dk.NewRound()
	if fmt.Sprint(err) != ErrRoundNotEnded {
		t.Error("We expect an error here")
		return
	}

	// play another card as same user, we expect an error if setNextPlayer
	// correctly did his job
	_, _, _, err = dk.PlayCard(players[0], card)
	if fmt.Sprint(err) != ErrNotYourTurn {
		t.Error("We expect an error")
		return
	}

	// closing round
	card2 := playerCards[players[1]].cards[0]
	next, desk, roundEnd, err = dk.PlayCard(players[1], card2)
	if err != nil {
		t.Error("No error expected")
		return
	}
	if len(desk) != 2 {
		t.Error("Invalid desk len")
		return
	}
	if desk[0].Card != card || desk[1].Card != card2 {
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
	if fmt.Sprint(err) != ErrRoundEnded {
		t.Error("We expected an error here")
		return
	}

	// compute round and see what happens
	next, _, err = dk.NewRound()
	if err != nil {
		t.Error("No error expected here")
		return
	}
	if next == (Player{}) {
		t.Error("next player should be the winner")
		return
	}
}
