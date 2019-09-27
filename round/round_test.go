package round

import (
	"testing"

	"github.com/briscola-as-a-service/game/card"
	"github.com/briscola-as-a-service/game/hand"
	"github.com/briscola-as-a-service/game/player"
	"github.com/briscola-as-a-service/game/seed"
)

func TestNew(t *testing.T) {
	r := New()
	if !r.winner.IsEmpty() {
		t.Error("winner should be an empty Player")
		return
	}
	if len(r.hands) != 0 {
		t.Error("hands should be an empty array")
		return
	}
	if r.wonPoints != 0 {
		t.Error("wonPoints should be zeroed")
		return
	}
}

func TestAddHand(t *testing.T) {
	r := New()

	player1 := player.New("P1", "Player1")
	card1 := card.New(seed.Denari(), 3)
	hand1 := hand.New(player1, card1)
	r.AddHand(hand1)

	if len(r.hands) != 1 {
		t.Error("Hands len should be 1")
		return
	}

	if r.hands[0] != hand1 {
		t.Error("Hand is wrong")
		return
	}

	if r.GetHands()[0] != r.hands[0] {
		t.Error("hands should be equal to GetHands()")
		return
	}

	// 2
	player2 := player.New("P2", "Player2")

	card2 := card.New(seed.Denari(), 10)
	hand2 := hand.New(player2, card2)
	r.AddHand(hand2)

	if len(r.hands) != 2 {
		t.Error("Hands len should be 2")
		return
	}

	if r.hands[0] != hand1 && r.hands[1] != hand2 {
		t.Error("Hands are wrong")
		return
	}

	if r.GetHands()[0] != r.hands[0] || r.GetHands()[1] != r.hands[1] {
		t.Error("hands should be equal to GetHands()")
		return
	}

	return
}

func TestComputeWinnerWithLoad(t *testing.T) {
	r := Round{}
	// p1:
	player1 := player.New("P1", "Player1")
	card1 := card.New(seed.Denari(), 3)
	hand1 := hand.New(player1, card1)
	// p2:
	player2 := player.New("P2", "Player2")
	card2 := card.New(seed.Denari(), 10)
	hand2 := hand.New(player2, card2)
	// compute winner with a briscola
	briscola := card.New(seed.Denari(), 1)
	//
	r.AddHand(hand1)
	r.AddHand(hand2)
	r.ComputeWinner(briscola)
	// the winner should be player 1
	if r.winner != player1 {
		t.Error("Winner should be player1")
		return
	}
	if r.wonPoints != 14 {
		t.Error("WonPoints should be 14")
		return
	}

	// compute without a briscola (taiet)
	briscola = card.New(seed.Spade(), 1)
	r.ComputeWinner(briscola)
	// the winner should be player 1
	if r.winner != player1 {
		t.Error("Winner should be player1")
		return
	}
	if r.wonPoints != 14 {
		t.Error("WonPoints should be 14")
		return
	}

	// play another round, reversing player order
	// briscola time
	briscola = card.New(seed.Denari(), 1)
	r = Round{}
	r.AddHand(hand2)
	r.AddHand(hand1)
	r.ComputeWinner(briscola)

	if !hand1.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}

	if !r.winner.Is(r.GetWinner()) {
		t.Error("Wrong winner got by GetWinner")
		return
	}

	// no briscola
	briscola = card.New(seed.Spade(), 1)
	r.ComputeWinner(briscola)

	if !hand1.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}
	return
}

func TestComputeWinnerWithoutLoad(t *testing.T) {
	r := Round{}
	// p1:
	player1 := player.New("P1", "Player1")
	card1 := card.New(seed.Denari(), 4)
	hand1 := hand.New(player1, card1)
	// p2:
	player2 := player.New("P2", "Player2")
	card2 := card.New(seed.Spade(), 6)
	hand2 := hand.New(player2, card2)
	// compute winner with a briscola
	briscola := card.New(seed.Denari(), 1)
	//
	r.AddHand(hand1)
	r.AddHand(hand2)
	r.ComputeWinner(briscola)
	// the winner should be player 1
	if !player1.Is(r.winner) {
		t.Error("Winner should be player1")
		return
	}
	if r.wonPoints != 0 {
		t.Error("WonPoints should be 0")
		return
	}

	// compute without a briscola (taiet)
	briscola = card.New(seed.Bastoni(), 1)
	r.ComputeWinner(briscola)
	// the winner should be player 1
	if r.winner != player1 {
		t.Error("Winner should be player1")
		return
	}
	if r.wonPoints != 0 {
		t.Error("WonPoints should be 0")
		return
	}

	// play another round, reversing player order
	// briscola time
	briscola = card.New(seed.Denari(), 1)
	r = Round{}
	r.AddHand(hand2)
	r.AddHand(hand1)
	r.ComputeWinner(briscola)

	if !hand1.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}

	// no briscola
	briscola = card.New(seed.Spade(), 1)
	r.ComputeWinner(briscola)

	if !hand2.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}

	// winning briscola 1
	briscola = hand1.GetCard()
	r.ComputeWinner(briscola)

	if !hand1.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}

	// winning briscola 2
	briscola = hand2.GetCard()
	r.ComputeWinner(briscola)

	if !hand2.GetPlayer().Is(r.winner) {
		t.Error("Wrong winner")
		return
	}
	return
}
