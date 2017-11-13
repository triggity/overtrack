package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Game struct {
	ID          int               `json:"user_id"`
	UserID      int               `json:"user_id" db:"user_id"`
	Map         GameMap           `json:"map_id"`
	StartTime   time.Time         `json:"start_time" db:"start_time"`
	GameType    GameType          `json:"game_type" db:"game_type"`
	GroupSize   int               `json:"group_size" db:"group_size"`
	isPlacement bool              `json:"is_placement" db:"is_placement"`
	Season      int               `json:"season" db:"season"`
	EndSR       sql.NullInt64     `json:"end_sr,omit_empty" db:"end_sr"`
	BeginSR     sql.NullInt64     `json:"begin_sr,omit_empty" db:"begin_sr"`
	Result      Result            `json:"result,string" db:"result"`
	Characters  []CharacterResult `json:"characters" db:"characters"`
	Stats       Stats             `json:"stats" db:"stats"`
	Disconnect  bool              `json:"disconnected" db:"disconnected"`
	leavers     int               `json:"leavers" db:"leavers"`
}

// Auxilary struct for `Game` for retrieving from database
type gameDB struct {
	ID          int           `db:"id"`
	UserID      int           `db:"user_id"`
	Map         string        `db:"map_id"`
	StartTime   time.Time     `db:"start_time"`
	GameType    GameType      `db:"game_type"`
	GroupSize   int           `db:"group_size"`
	isPlacement bool          `db:"is_placement"`
	Season      int           `db:"season"`
	EndSR       sql.NullInt64 `db:"end_sr"`
	BeginSR     sql.NullInt64 `db:"begin_sr"`
	Result      Result        `db:"result"`
	Disconnect  bool          `db:"disconnected"`
	leavers     int           `db:"leavers"`
}

type GameController struct {
	db    *sqlx.DB
	index string
}

func NewGameController(db *sqlx.DB) *GameController {
	return &GameController{db, "ow"}
}

func (g *GameController) GetByUser(id int) ([]Game, error) {
	// gameAux := gameDB{}
	// err := g.db.Get()
	return []Game{}, nil
}

func (g *GameController) GetByGame(gameId int) (Game, error) {
	aux := gameDB{}
	game := Game{}
	err := g.db.Get(&aux, "SELECT * FROM games WHERE id=$1", gameId)
	if err != nil {
		return game, err
	}
	game.ID = aux.ID
	game.UserID = aux.UserID
	// map
	game.StartTime = aux.StartTime
	game.GameType = aux.GameType
	game.GroupSize = aux.GroupSize
	game.isPlacement = aux.isPlacement
	game.Season = aux.Season
	if aux.EndSR.Valid {
		game.EndSR = aux.EndSR
	}
	if aux.BeginSR.Valid {
		game.BeginSR = aux.BeginSR
	}
	game.Result = aux.Result
	game.Disconnect = aux.Disconnect
	game.leavers = aux.leavers
	return game, nil
}
