package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Game struct {
	ID          int           `json:"user_id"`
	UserID      int           `json:"user_id" db:"user_id"`
	Map         GameMap       `json:"map_id"`
	StartTime   time.Time     `json:"start_time" db:"start_time"`
	GameType    GameType      `json:"game_type" db:"game_type"`
	GroupSize   int           `json:"group_size" db:"group_size"`
	IsPlacement bool          `json:"is_placement" db:"is_placement"`
	Season      int           `json:"season" db:"season"`
	EndSR       sql.NullInt64 `json:"end_sr,omit_empty" db:"end_sr"`
	BeginSR     sql.NullInt64 `json:"begin_sr,omit_empty" db:"begin_sr"`
	Result      Result        `json:"result,string" db:"result"`
	Characters  []HeroResult  `json:"characters"`
	Stats       Stats         `json:"stats" db:"stats"`
	Disconnect  bool          `json:"disconnected" db:"disconnected"`
	Leavers     int           `json:"leavers" db:"leavers"`
}

// Auxilary struct for `Game` for retrieving from database
type gameDB struct {
	ID          int           `db:"id"`
	UserID      int           `db:"user_id"`
	MapID       int           `db:"map_id"`
	Result      Result        `db:"result"`
	StartTime   time.Time     `db:"start_time"`
	GroupSize   int           `db:"group_size"`
	IsPlacement bool          `db:"is_placement"`
	Season      int           `db:"season"`
	EndSR       sql.NullInt64 `db:"end_sr"`
	BeginSR     sql.NullInt64 `db:"begin_sr"`
	Leavers     int           `db:"leavers"`
	Disconnect  bool          `db:"disconnect"`
}

type GameController struct {
	db     *sqlx.DB
	logger *log.Entry
}

func NewGameController(db *sqlx.DB) *GameController {
	return &GameController{db, log.WithFields(log.Fields{
		"name": "game_controller",
	})}
}

func (g *GameController) GetByUser(id int) ([]Game, error) {
	// gameAux := gameDB{}
	// err := g.db.Get()
	return []Game{}, nil
}

func (g *GameController) GetByGame(userId int, gameId int) (Game, error) {
	aux := gameDB{}
	game := Game{}
	err := g.db.Get(&aux, "SELECT * FROM games WHERE id=$1", gameId)
	if err != nil {
		g.logger.Infof("failed to retrieve game %d", gameId, err)
		return game, err
	}
	// get map to fill
	gameMap, err := g.getMapByID(aux.MapID)
	if err != nil {
		g.logger.Infof("failed to get map %d for game %d", aux.MapID, gameId, err)
		return game, err
	}

	gameStats, err := g.getGameStatsByID(aux.UserID, gameId)
	if err != nil {
		g.logger.Infof("failed to get game stats for game %d for user %d", gameId, userId, err)
		return game, err
	}
	heroStats, err := g.getHeroStatsByID(userId, gameId)
	if err != nil {
		g.logger.Infof("failed to get hero stats for game %d for user %d", gameId, userId, err)
		return game, err
	}

	game.ID = aux.ID
	game.UserID = aux.UserID
	game.Map = gameMap
	game.StartTime = aux.StartTime
	// game.GameType = aux.GameType
	game.GroupSize = aux.GroupSize
	game.IsPlacement = aux.IsPlacement
	game.Season = aux.Season
	if aux.EndSR.Valid {
		game.EndSR = aux.EndSR
	}
	if aux.BeginSR.Valid {
		game.BeginSR = aux.BeginSR
	}
	game.Result = aux.Result
	game.Disconnect = aux.Disconnect
	game.Leavers = aux.Leavers
	game.Stats = gameStats
	game.Characters = heroStats

	return game, nil
}
func (g *GameController) getMapByID(gameId int) (GameMap, error) {
	gameMap := GameMap{}
	err := g.db.Get(&gameMap, "SELECT * FROM maps WHERE id=$1 LIMIT 1", gameId)
	return gameMap, err
}

func (g *GameController) getGameStatsByID(userId int, gameId int) (Stats, error) {
	stats := Stats{}
	err := g.db.Get(&stats, "SELECT * FROM game_stats WHERE user_id=$1 AND game_id=$2 LIMIT 1", userId, gameId)
	return stats, err
}

func (g *GameController) getGameStatsForUser(userId int) ([]Stats, error) {
	stats := []Stats{}
	err := g.db.Select(&stats, "SELECT * FROM game_stats WHERE user_id=$1", userId)
	return stats, err
}

// func (g *GameController) getHeroStatsByID(userId int, gameId int) ([])
