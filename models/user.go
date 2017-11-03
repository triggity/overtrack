package models

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"

	"github.com/triggity/overtrack/errs"
	elastic "gopkg.in/olivere/elastic.v5"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDao struct {
	client *elastic.Client
	index  string
}

func NewUserDao(client *elastic.Client) *UserDao {
	return &UserDao{client, "ow"}
}

func (u *UserDao) GetByID(ctx context.Context, id int) (User, error) {
	var result User
	resp, err := u.client.Get().Index(u.index).Type("users").Id(strconv.Itoa(id)).Do(ctx)
	if err != nil {
		return result, err
	}
	if !resp.Found {
		return result, errs.ErrorNotFound
	}
	err = json.Unmarshal(*resp.Source, &result)
	if err != nil {
		return result, err
	}
	result.ID = id
	return result, err
}

func (u *UserDao) List(ctx context.Context) ([]User, error) {
	resp, err := u.client.Search().Index(u.index).Type("users").Do(ctx)
	if err != nil {
		return nil, err
	}
	var users []User
	for _, item := range resp.Each(reflect.TypeOf(User{})) {
		var t User
		var ok bool
		if t, ok = item.(User); !ok {
			return nil, errors.New("Trouble deserializing User")
		}
		users = append(users, t)
	}

	return users, nil
}
