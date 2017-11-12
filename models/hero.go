package models

import (
	"github.com/jmoiron/sqlx"
)

type Hero struct {
	ID    int       `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	Class HeroClass `json:"class" db:"class"`
}

type HeroDao struct {
	db        *sqlx.DB
	tableName string
}

func NewHeroDao(sqlClient *sqlx.DB) *HeroDao {
	return &HeroDao{sqlClient, "heros"}
}

func (h *HeroDao) GetHero(name string) (Hero, error) {
	hero := Hero{}
	err := h.db.Get("SELECT * FROM ? WHERE name=? LIMIT 1;", h.tableName, name)
	return hero, err
}

func (h *HeroDao) GetHeros() ([]Hero, error) {
	heros := []Hero{}
	err := h.db.Select("SELECT * FROM ?;", h.tableName)
	return heros, err
}
