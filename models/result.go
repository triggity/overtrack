package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type Result uint8

const (
	Win Result = iota
	Loss
	Draw
)

func (r *Result) Scan(value interface{}) error {
	// aux := GameType()
	var aux Result
	v, ok := value.([]byte)
	if !ok {
		return errors.New("value is not of type []byte")
	}
	s := string(v)
	aux, ok = _ResultNameToValue[strings.Title(s)]
	if !ok {
		return fmt.Errorf("invalid Result %q", s)
	}
	*r = aux
	return nil
}

func (r Result) Value() (driver.Value, error) {
	return strings.ToLower(r.String()), nil
}
