package models

import (
	"github.com/jmoiron/sqlx"
)

type Hero struct {
	ID    int       `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	Class HeroClass `json:"class" db:"class"`
}

type HeroController struct {
	db        *sqlx.DB
	tableName string
}

func NewHeroDao(sqlClient *sqlx.DB) *HeroController {
	return &HeroController{sqlClient, "heros"}
}

func (h *HeroController) GetHero(id int) (Hero, error) {
	hero := Hero{}
	err := h.db.Get(&hero, "SELECT * FROM heros WHERE id=$1 LIMIT 1", id)
	return hero, err
}

func (h *HeroController) List() ([]Hero, error) {
	heros := []Hero{}
	err := h.db.Select(&heros, "SELECT * FROM heros")
	return heros, err
}
