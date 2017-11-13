package models

import "github.com/jmoiron/sqlx"

type Stats struct {
	GameID         int `json:"game_id" db:"game_id"`
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

type StatsController struct {
	db *sqlx.DB
}

func NewStatsController(sqlClient *sqlx.DB) *StatsController {
	return &StatsController{sqlClient}
}

func (s *StatsController) GetByGame(userId int, gameId int) (Stats, error) {
	stats := Stats{}
	err := s.db.Get(&stats, "SELECT * FROM game_stats WHERE user_id=$1 AND game_id=$2 LIMIT 1", userId, gameId)
	return stats, err
}

func (s *StatsController) GetByUser(userId int) ([]Stats, error) {
	stats := []Stats{}
	err := s.db.Select(&stats, "SELECT * FROM game_stats WHERE user_id=$1", userId)
	return stats, err
}
