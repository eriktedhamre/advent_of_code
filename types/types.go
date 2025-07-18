package types

// I'm starting to think that I do not like Go-generics :)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Coordinates struct {
	Row, Col int64
}

type Crucible struct {
	Row, Col         int64
	ConsecutiveMoves int
	Direction        Direction
	Parents          []Coordinates
}
