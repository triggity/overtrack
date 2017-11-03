package models

import (
	"context"
	"errors"
	"reflect"

	"github.com/triggity/overtrack/errs"
	elastic "gopkg.in/olivere/elastic.v5"
)

// GameMap represents a game map
type GameMap struct {
	Name     string   `json:"name"`
	FullName string   `json:"full_name"`
	City     string   `json:"city"`
	Country  string   `json:"country"`
	GameType GameType `json:"game_type"`
}

type GameMapDao struct {
	client *elastic.Client
}

func NewGameMapDao(client *elastic.Client) *GameMapDao {
	return &GameMapDao{client}
}

func (g *GameMapDao) GetByName(ctx context.Context, name string) (GameMap, error) {
	var result GameMap
	query := elastic.NewTermQuery("name", name)
	resp, err := g.client.Search().Index("ow").Type("maps").Query(query).Do(ctx)
	if err != nil {
		return result, err
	}
	for _, item := range resp.Each(reflect.TypeOf(result)) {
		var ok bool
		if result, ok = item.(GameMap); ok {
			return result, nil
		}
		return result, errors.New("Trouble deserializing gamemap")
	}
	return result, errs.ErrorNotFound

}

func (g *GameMapDao) List(ctx context.Context) ([]GameMap, error) {
	resp, err := g.client.Search().Index("ow").Type("maps").Do(ctx)
	if err != nil {
		return nil, err
	}
	var maps []GameMap
	for _, item := range resp.Each(reflect.TypeOf(GameMap{})) {
		var t GameMap
		var ok bool
		if t, ok = item.(GameMap); !ok {
			return nil, errors.New("Trouble deserializing gamemap")
		}
		maps = append(maps, t)
	}
	if len(maps) == 0 {
		return nil, errs.ErrorNotFound
	}

	return maps, nil
}
