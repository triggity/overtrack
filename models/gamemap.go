package models

import (
	"github.com/jmoiron/sqlx"
)

// GameMap represents a game map
type GameMap struct {
	ID       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	FullName string   `json:"full_name" db:"full_name"`
	City     string   `json:"city" db:"city"`
	Country  string   `json:"country" db:"country"`
	GameType GameType `json:"game_type" db:"game_type"`
}

type GameMapController struct {
	db *sqlx.DB
}

func NewGameMapDao(db *sqlx.DB) *GameMapController {
	return &GameMapController{db}
}

func (g *GameMapController) GetByID(gameId int) (GameMap, error) {
	gameMap := GameMap{}
	err := g.db.Get(&gameMap, "SELECT * FROM maps WHERE name=$1 LIMIT 1", gameId)
	return gameMap, err
}

func (g *GameMapController) List() ([]GameMap, error) {
	gameMaps := []GameMap{}
	err := g.db.Select(&gameMaps, "SELECT * FROM maps")
	return gameMaps, err
}
