package models

import (
	"github.com/jmoiron/sqlx"
)

// GameMap represents a game map
type GameMap struct {
	ID       int      `db:"id"`
	Name     string   `json:"name" db:"name"`
	FullName string   `json:"full_name" db:"full_name"`
	City     string   `json:"city" db:"city"`
	Country  string   `json:"country" db:"country"`
	GameType GameType `json:"game_type" db:"game_type"`
}

type GameMapDao struct {
	db        *sqlx.DB
	tableName string
}

func NewGameMapDao(db *sqlx.DB) *GameMapDao {
	return &GameMapDao{db, "maps"}
}

func (g *GameMapDao) GetByName(name string) (GameMap, error) {
	gameMap := GameMap{}
	err := g.db.Get("SELECT * FROM ? WHERE name=? LIMIT 1", g.tableName, name)
	return gameMap, err
}

func (g *GameMapDao) List() ([]GameMap, error) {
	gameMaps := []GameMap{}
	err := g.db.Select("SELECT * FROM ?", g.tableName)
	return gameMaps, err
}
