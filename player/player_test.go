package player

import "testing"

func TestNew(t *testing.T) {
	p := New("id", "name")
	if p.id != "id" || p.name != "name" {
		t.Error("invalid id or name")
		return
	}
}

func TestIsEmpty(t *testing.T) {
	p1 := Player{}
	p2 := New("id", "name")
	if !p1.IsEmpty() {
		t.Error("p1 should be empty")
		return
	}
	if p2.IsEmpty() {
		t.Error("p2 should not be empty")
		return
	}
}

func TestIs(t *testing.T) {
	p1 := New("id", "name")
	p2 := New("id", "name2")
	p3 := New("id3", "name")

	if !p1.Is(p2) {
		t.Error("p1 is p2")
		return
	}

	if p1.Is(p3) {
		t.Error("p1 is not p3")
		return
	}
}
