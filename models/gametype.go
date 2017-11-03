package models

import "strconv"

// GameType represents one of the overwatch game types
type GameType uint8

const (
	_               = iota
	Escort GameType = 1
	Assault
	Hyrbid
	Control
)

func (g GameType) String() string {
	name := []string{"escort", "assault", "hybrid", "control"}
	i := uint8(g)
	switch {
	case i <= uint8(Control):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
