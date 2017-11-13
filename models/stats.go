package models

type Stats struct {
	ID             int `json:"-" db:"id"`
	UserID         int `json:"-" db:"user_id"`
	GameID         int `json:"-" db:"game_id"`
	Eliminations   int `json:"eliminations" db:"eliminations"`
	ObjectiveKills int `json:"objective_kills,omitempty" db:"objective_kills"`
	ObjectiveTime  int `json:"objective_time" db:"objective_time"`
	HeroDamage     int `json:"hero_damage" db:"hero_damage"`
	Healing        int `json:"healing" db:"healing"`
	Deaths         int `json:"deaths" db:"deaths"`
}

func (s *Stats) CoreStats() *Stats {
	return s
}
