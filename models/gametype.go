package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// GameType represents one of the overwatch game types
type GameType uint8

const (
	Escort GameType = iota
	Assault
	Hyrbid
	Control
)

func ParseGameType(data string) (GameType, error) {
	return Escort, nil
}

func (g *GameType) Scan(value interface{}) error {
	// aux := GameType()
	var aux GameType
	v, ok := value.([]byte)
	if !ok {
		return errors.New("value is not of type []byte")
	}
	s := string(v)
	aux, ok = _GameTypeNameToValue[strings.Title(s)]
	if !ok {
		return fmt.Errorf("invalid GameType %q", s)
	}
	*g = aux
	return nil
}

func (g GameType) Value() (driver.Value, error) {
	return strings.ToLower(g.String()), nil
}
