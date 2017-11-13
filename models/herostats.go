package models

import (
	"github.com/jmoiron/sqlx/types"
	log "github.com/sirupsen/logrus"
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
	UserID         int            `db:"user_id"`
	GameID         int            `db:"game_id"`
	HeroID         int            `db:"hero_id"`
	Eliminations   int            `db:"eliminations"`
	ObjectiveKills int            `db:"objective_kills"`
	ObjectiveTime  int            `db:"objective_time"`
	HeroDamage     int            `db:"hero_damage"`
	Healing        int            `db:"healing"`
	Deaths         int            `db:"deaths"`
	Stats          types.JSONText `db:"custom_stats"`
	HeroTID        int            `db:"id"`
	Name           string         `db:"name"`
	Class          HeroClass      `db:"class"`
}

func (g *GameController) getHeroStatsByID(userID int, gameID int) ([]HeroResult, error) {
	aux := []heroResultDB{}
	results := []HeroResult{}
	_ = []struct {
		Name string
	}{
		{"foo"},
	}
	// err := g.db.Select(&aux, "SELECT * FROM hero_stats WHERE user_id=1 AND game_id=1")
	err := g.db.Select(&aux, "SELECT * FROM hero_stats INNER JOIN heros ON hero_stats.hero_id=heros.id WHERE hero_stats.user_id=$1 AND hero_stats.game_id=$2", userID, gameID)
	// err := g.db.Select(&aux, "SELECT hero_stats.id,user_id,game_id,hero_id,eliminations,objective_kills,hero_damage,healing,deaths,custom_stats,heros.id as hero_t_id,name,class FROM hero_stats INNER JOIN heros ON hero_stats.hero_id=heros.id WHERE hero_stats.user_id=$1 AND hero_stats.game_id=$2", userID, gameID)
	for i, p := range aux {
		log.Printf("%d => %v , %v", i, p.ID, string(p.Stats))
	}
	g.logger.Info("results user: %d, game: %d", userID, gameID, aux)
	return results, err
}

/*
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

*/
