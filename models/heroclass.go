package models

type HeroClass uint8

const (
	Attack HeroClass = iota
	Defense
	Tank
	Suppport
)
