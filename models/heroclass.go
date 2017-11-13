package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type HeroClass uint8

const (
	Attack HeroClass = iota
	Defense
	Tank
	Support
)

func (h *HeroClass) Scan(value interface{}) error {
	// aux := GameType()
	var aux HeroClass
	v, ok := value.([]byte)
	if !ok {
		return errors.New("value is not of type []byte")
	}
	s := string(v)
	aux, ok = _HeroClassNameToValue[strings.Title(s)]
	if !ok {
		return fmt.Errorf("invalid HeroClass %q", s)
	}
	*h = aux
	return nil
}

func (h HeroClass) Value() (driver.Value, error) {
	return strings.ToLower(h.String()), nil
}
