package player

// Player contains the player data.
type Player struct {
	name string
	id   string
}

// New returns a new Player
func New(id, name string) Player {
	return Player{
		name,
		id,
	}
}

// IsEmpty returns true if .id is an empty string
func (p Player) IsEmpty() bool {
	return p.id == ""
}

// Is returns true if who is p
func (p Player) Is(who Player) bool {
	return p.id == who.id
}

// ID returns the current sd
func (p Player) ID() string {
	return p.id
}
