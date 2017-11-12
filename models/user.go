package models

import (
	"github.com/jmoiron/sqlx"

	elastic "gopkg.in/olivere/elastic.v5"
)

type User struct {
	ID         int    `json:"id" db:"user_id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	BlizzardID string `json:"blizzard_id" db:"blizzard_id"`
}

type UserDao struct {
	client    *elastic.Client
	db        *sqlx.DB
	tableName string
}

func NewUserDao(client *elastic.Client, db *sqlx.DB) *UserDao {
	return &UserDao{client, db, "users"}
}

func (u *UserDao) GetByID(id int) (User, error) {
	user := User{}
	err := u.db.Get(&user, "SELECT * FROM ? WHERE user_id=? ORDER BY last_updated DESC LIMIT 1", u.tableName, id)
	return user, err
}

func (u *UserDao) List() ([]User, error) {
	users := []User{}
	err := u.db.Select(&users, "SELECT * FROM ?", u.tableName)
	return users, err
}
