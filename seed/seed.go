package seed

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
	// todo randomize
	return Seed{seedSpade}
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
