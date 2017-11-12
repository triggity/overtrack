package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Game struct {
	UserID     int               `json:"user_id" db:"user_id"`
	Map        string            `json:"map" db:"map"`
	StartTime  time.Time         `json:"start_time,string" db:"start_time"`
	GameType   GameType          `json:"game_type, string" db:"game_type"`
	GroupSize  int               `json:"group_size" db:"group_size"`
	IsSeason   bool              `json:"is_placement" db:"is_placement"`
	Season     int               `json:"season" db:"season"`
	EndSR      int               `json:"end_sr,omit_empty" db:"end_sr"`
	BeginSR    int               `json:"begin_sr,omit_empty" db:"begin_sr"`
	Result     Result            `json:"result,string" db:"result"`
	Characters []CharacterResult `json:"characters" db:"characters"`
	Stats      Stats             `json:"stats" db:"stats"`
}

type GameDao struct {
	db    *sqlx.DB
	index string
}

func NewGameDao(db *sqlx.DB) *GameDao {
	return &GameDao{db, "ow"}
}

func (g *GameDao) GetByUser(id int) ([]Game, error) {
	return []Game{}, nil
}
