package models

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID         int    `json:"id" db:"user_id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	BlizzardID string `json:"blizzard_id" db:"blizzard_id"`
}

type UserController struct {
	db        *sqlx.DB
	tableName string
}

func NewUserDao(db *sqlx.DB) *UserController {
	return &UserController{db, "users"}
}

func (u *UserController) GetByID(id int) (User, error) {
	user := User{}
	err := u.db.Get(&user, "SELECT * FROM ? WHERE user_id=? ORDER BY last_updated DESC LIMIT 1", u.tableName, id)
	return user, err
}

func (u *UserController) List() ([]User, error) {
	users := []User{}
	err := u.db.Select(&users, "SELECT * FROM ?", u.tableName)
	return users, err
}
