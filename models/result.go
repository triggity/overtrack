package models

import "strconv"

type Result uint8

const (
	_          = iota
	Win Result = 1
	Loss
	Draw
)

func (r Result) String() string {
	name := []string{"win", "loss", "draw"}
	i := uint8(r)
	switch {
	case i <= uint8(Draw):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
