package seed

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	count := 10
	var randoms []Seed
	for i := 0; i < count; i++ {
		randoms = append(randoms, Random())
	}

	atLeastOneDiffers := false
	for i := 0; i < count-1; i++ {
		if randoms[i] != randoms[i+1] {
			atLeastOneDiffers = true
		}
	}
	if !atLeastOneDiffers {
		t.Error("We expect a random here")
		return
	}
}

func TestSeeds(t *testing.T) {
	if Denari().seed != seedDenari {
		t.Error("invalid seed for denari")
		return
	}
	if Coppe().seed != seedCoppe {
		t.Error("invalid seed for coppe")
		return
	}
	if Spade().seed != seedSpade {
		t.Error("invalid seed for spade")
		return
	}
	if Bastoni().seed != seedBastoni {
		t.Error("invalid seed for bastoni")
		return
	}
}

func TestIterable(t *testing.T) {
	it := Iterable()

	if len(it) != 4 {
		t.Error("len should be 4")
		return
	}

	for i := 0; i < 3; i++ {
		for j := 3; j > i; j-- {
			fmt.Printf("comparing %d with %d\n", i, j)
			if it[i] == it[j] {
				t.Error("Found dupes")
				return
			}
		}
	}

	for _, s := range it {
		if !s.IsValid() {
			t.Error("found invalid seed")
			return
		}
	}
}

func TestIsValid(t *testing.T) {
	valid1 := Denari()
	valid2 := Coppe()
	valid3 := Spade()
	valid4 := Bastoni()
	invalid := Seed{}

	if !valid1.IsValid() ||
		!valid2.IsValid() ||
		!valid3.IsValid() ||
		!valid4.IsValid() {
		t.Error("those are valid")
		return
	}

	if invalid.IsValid() {
		t.Error("this is not valid")
		return
	}
}

func TestString(t *testing.T) {
	if fmt.Sprintf("%s", Denari()) != "D" {
		t.Error("invalid for Denari()")
		return
	}
	if fmt.Sprintf("%s", Coppe()) != "C" {
		t.Error("invalid for Coppe()")
		return
	}
	if fmt.Sprintf("%s", Spade()) != "S" {
		t.Error("invalid for Spade()")
		return
	}
	if fmt.Sprintf("%s", Bastoni()) != "B" {
		t.Error("invalid for Bastoni()")
		return
	}
}
