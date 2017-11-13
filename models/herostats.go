package models

import (
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx/types"
)

type HeroStats interface {
	CoreStats() *Stats
	CharacterStats() map[string]float32
}

type HeroResult struct {
	Name  string    `json:"name"`
	Class HeroClass `json:"class"`
	Stats HeroStats `json:"stats"`
}

// used for unmarshaling to specific character's results
// WARNING: order matters here
type heroResultDB struct {
	ID             int            `db:"id"`
	GameID         int            `db:"game_id"`
	HeroID         int            `db:"hero_id"`
	Eliminations   int            `db:"eliminations"`
	ObjectiveKills int            `db:"objective_kills"`
	ObjectiveTime  int            `db:"objective_time"`
	HeroDamage     int            `db:"hero_damage"`
	Healing        int            `db:"healing"`
	Deaths         int            `db:"deaths"`
	Stats          types.JSONText `db:"custom_stats"`
	Name           string         `db:"name"`
	Class          HeroClass      `db:"class"`
}

func (g *GameController) getHeroStatsByID(userID int, gameID int) ([]HeroResult, error) {
	stats := []HeroResult{}
	err := g.db.Select(&stats, "SELECT *,heros.name,heros.class FROM hero_stats INNER JOIN heros ON hero_stats.hero_id=heros.id WHERE hero_stats.user_id=$1 AND hero_stats.game_id=$2", userID, gameID)
	return stats, err
}

func (c *HeroResult) UnmarshalJSON(d []byte) error {
	var cr heroResultDB
	err := json.Unmarshal(d, &cr)
	if err != nil {
		return err
	}
	c.Name = cr.Name
	switch c.Name {
	case "orisa":
		var o HeroResultOrisa
		err := json.Unmarshal(cr.Stats, &o)
		if err != nil {
			return err
		}
		c.Stats = &o
		return nil
	}
	return errors.New("foobar error")
}

type HeroResultOrisa struct {
	*Stats
	HeroClass      HeroClass `json:"class" db:"class"`
	DamagedBlocked int       `json:"damage_blocked" db:"damage_blocked"`
}

func (o *HeroResultOrisa) CharacterStats() map[string]float32 {
	return map[string]float32{
		"damage_blocked": float32(o.DamagedBlocked),
	}
}

type HeroResultReindhardt struct {
	*Stats
	HeroClass      HeroClass `json:"class" db:"class"`
	DamagedBlocked int       `json:"damage_blocked" db:"damage_blocked"`
}

func (r *HeroResultReindhardt) CharacterStats() map[string]float32 {
	return map[string]float32{
		"damage_blocked": float32(r.DamagedBlocked),
	}
}
