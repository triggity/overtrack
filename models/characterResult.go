package models

import (
	"encoding/json"
	"errors"
)

type CharacterStats interface {
	CoreStats() *Stats
	CharacterStats() map[string]float32
}

type CharacterResult struct {
	Name  string         `json:"name" db:"name"`
	Stats CharacterStats `json:"stats" db:"stats"`
}

// used for unmarshaling to specific character's results
type characterResultRaw struct {
	Name  string          `json:"name" db:"name"`
	Stats json.RawMessage `json:"stats" db:"stats"`
}

func (c *CharacterResult) UnmarshalJSON(d []byte) error {
	var cr characterResultRaw
	err := json.Unmarshal(d, &cr)
	if err != nil {
		return err
	}
	c.Name = cr.Name
	switch c.Name {
	case "orisa":
		var o CharacterResultOrisa
		err := json.Unmarshal(cr.Stats, &o)
		if err != nil {
			return err
		}
		c.Stats = &o
		return nil
	}
	return errors.New("foobar error")
}

type CharacterResultOrisa struct {
	*Stats
	CharacterClass CharacterClass `json:"class" db:"class"`
	DamagedBlocked int            `json:"damage_blocked" db:"damage_blocked"`
}

func (o *CharacterResultOrisa) CharacterStats() map[string]float32 {
	return map[string]float32{
		"damage_blocked": float32(o.DamagedBlocked),
	}
}

type CharacterResultReinhardt struct {
	*Stats
	CharacterClass CharacterClass `json:"class" db:"class"`
	DamagedBlocked int            `json:"damage_blocked" db:"damage_blocked"`
}

func (r *CharacterResultReinhardt) CharacterStats() map[string]float32 {
	return map[string]float32{
		"damage_blocked": float32(r.DamagedBlocked),
	}
}
