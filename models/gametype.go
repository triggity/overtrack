package models

// GameType represents one of the overwatch game types
type GameType uint8

const (
	Escort GameType = iota
	Assault
	Hyrbid
	Control
)
