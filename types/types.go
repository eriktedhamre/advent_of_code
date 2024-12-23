package types

// I'm starting to think that I do not like Go-generics :)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Crucible struct {
	Row, Col         int
	ConsecutiveMoves int
	Direction        Direction
}