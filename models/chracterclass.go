package models

import "strconv"

type CharacterClass uint8

const (
	_                     = iota
	Attack CharacterClass = 1
	Defense
	Tank
	Suppport
)

func (r CharacterClass) String() string {
	name := []string{"attack", "defense", "tank", "support"}
	i := uint8(r)
	switch {
	case i <= uint8(Suppport):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
