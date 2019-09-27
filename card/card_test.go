package card

import (
	"testing"

	"github.com/briscola-as-a-service/game/seed"
)

func TestNew(t *testing.T) {
	c := New(seed.Bastoni(), 1)

	if !c.seed.IsBastoni() || c.value != 1 {
		t.Error("card is wrong!")
		return
	}
}

func TestNewEmpty(t *testing.T) {
	c := NewEmpty()

	if c.seed.IsValid() || c.value != 0 {
		t.Error("card should be invalid")
		return
	}
}

func TestEquals(t *testing.T) {
	c1 := New(seed.Bastoni(), 1)
	c2 := New(seed.Bastoni(), 1)
	c3 := New(seed.Bastoni(), 2)
	c4 := New(seed.Coppe(), 3)

	if !c1.Equals(c2) {
		t.Error("c1 should equals c2")
		return
	}

	if c1.Equals(c3) {
		t.Error("c1 != c3")
		return
	}

	if c3.Equals(c4) {
		t.Error("c3 != c4")
		return
	}
}

func TestIsBriscola(t *testing.T) {
	c1 := New(seed.Bastoni(), 2)
	c2 := New(seed.Denari(), 4)
	briscola := New(seed.Denari(), 1)

	if c1.IsBriscola(briscola) {
		t.Error("c1 is not a briscola")
		return
	}

	if !c2.IsBriscola(briscola) {
		t.Error("c2 is a briscola")
		return
	}
}

func TestPoints(t *testing.T) {
	zeroes := []int{2, 4, 5, 6, 7}
	nonZeroes := []int{1, 3, 8, 9, 10}
	// with this, we check consistency with points const in `const.go` file
	nonZeroValues := map[int]int{
		1:  11,
		3:  10,
		8:  2,
		9:  3,
		10: 4,
	}

	// zero values
	for _, zero := range zeroes {
		tmpCard := New(seed.Random(), zero)
		points := tmpCard.Points()
		if points != 0 {
			t.Error("Points should be == 0")
			return
		}
	}

	// others values
	for _, nonZero := range nonZeroes {
		tmpCard := New(seed.Random(), nonZero)
		points := tmpCard.Points()
		if points != nonZeroValues[nonZero] {
			t.Errorf("Points should be == %v and not %v",
				nonZeroValues[nonZero], points)
			return
		}
	}
}

func TestValue(t *testing.T) {
	for i := 1; i < 11; i++ {
		tmpCard := New(seed.Random(), i)
		if tmpCard.Value() != i {
			t.Error("wrong value")
			return
		}
	}
}

func TestIsExpendable(t *testing.T) {
	c1 := New(seed.Denari(), 2)
	c2 := New(seed.Spade(), 3)
	c3 := New(seed.Spade(), 2)

	if c1.IsExpendable() || c2.IsExpendable() {
		t.Error("c1 and c2 are not expendables")
		return
	}

	if !c3.IsExpendable() {
		t.Error("c3 is expendable")
		return
	}
}
