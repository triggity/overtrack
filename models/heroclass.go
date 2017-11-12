package models

import "strconv"

type HeroClass uint8

const (
	_                = iota
	Attack HeroClass = 1
	Defense
	Tank
	Suppport
)

func (r HeroClass) String() string {
	name := []string{"attack", "defense", "tank", "support"}
	i := uint8(r)
	switch {
	case i <= uint8(Suppport):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
