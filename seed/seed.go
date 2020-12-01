package seed

import (
	"math/rand"
	"time"
)

// Seed type
type Seed struct {
	seed string
}

// Card seeds
const (
	seedDenari  = "DENARI"
	seedSpade   = "SPADE"
	seedBastoni = "BASTONI"
	seedCoppe   = "COPPE"
)

// Random returns a random Seed
func Random() Seed {
	rand.Seed(time.Now().UnixNano())
	it := Iterable()
	return it[rand.Intn(len(it))]
}

// Denari returns a Denari
func Denari() Seed {
	return Seed{seedDenari}
}

// Spade returns a SeedSpade
func Spade() Seed {
	return Seed{seedSpade}
}

// Bastoni returns a Basotni
func Bastoni() Seed {
	return Seed{seedBastoni}
}

// Coppe returns a Coppe
func Coppe() Seed {
	return Seed{seedCoppe}
}

// Iterable returns an iterable object of strings
func Iterable() []Seed {
	return []Seed{
		Denari(),
		Spade(),
		Bastoni(),
		Coppe(),
	}
}

// IsValid checks if seed is valid
func (s Seed) IsValid() bool {
	return s.IsSpade() || s.IsDenari() || s.IsCoppe() || s.IsBastoni()
}

// IsSpade returns true if seed is SeedSpade
func (s Seed) IsSpade() bool {
	return s.seed == seedSpade
}

// IsDenari returns true if seed is SeedDenari
func (s Seed) IsDenari() bool {
	return s.seed == seedDenari
}

// IsBastoni returns true if seed is SeedBastoni
func (s Seed) IsBastoni() bool {
	return s.seed == seedBastoni
}

// IsCoppe returns true if seed is SeedCoppe
func (s Seed) IsCoppe() bool {
	return s.seed == seedCoppe
}

func (s Seed) String() string {
	switch true {
	case s.IsCoppe():
		return "C"
	case s.IsSpade():
		return "S"
	case s.IsDenari():
		return "D"
	}
	// case s.IsBastoni(): || default:
	return "B"
}
