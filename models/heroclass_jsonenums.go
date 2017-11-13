// generated by jsonenums -type=HeroClass; DO NOT EDIT

package models

import (
	"encoding/json"
	"fmt"
)

var (
	_HeroClassNameToValue = map[string]HeroClass{
		"Attack":  Attack,
		"Defense": Defense,
		"Tank":    Tank,
		"Support": Support,
	}

	_HeroClassValueToName = map[HeroClass]string{
		Attack:  "Attack",
		Defense: "Defense",
		Tank:    "Tank",
		Support: "Support",
	}
)

func init() {
	var v HeroClass
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_HeroClassNameToValue = map[string]HeroClass{
			interface{}(Attack).(fmt.Stringer).String():  Attack,
			interface{}(Defense).(fmt.Stringer).String(): Defense,
			interface{}(Tank).(fmt.Stringer).String():    Tank,
			interface{}(Support).(fmt.Stringer).String(): Support,
		}
	}
}

// MarshalJSON is generated so HeroClass satisfies json.Marshaler.
func (r HeroClass) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _HeroClassValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid HeroClass: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so HeroClass satisfies json.Unmarshaler.
func (r *HeroClass) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("HeroClass should be a string, got %s", data)
	}
	v, ok := _HeroClassNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid HeroClass %q", s)
	}
	*r = v
	return nil
}