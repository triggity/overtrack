package models

import (
	"context"
	"errors"
	"reflect"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Game struct {
	UserID     int               `json:"user_id"`
	Map        string            `json:"map"`
	StartTime  time.Time         `json:"start_time,string"`
	GameType   GameType          `json:"game_type, string"`
	GroupSize  int               `json:"group_size"`
	IsSeason   bool              `json:"is_placement"`
	Season     int               `json:"season"`
	EndSR      int               `json:"end_sr, omit_empty"`
	BeginSR    int               `json:"begin_sr, omit_empty"`
	Result     Result            `json:"result, string"`
	Characters []CharacterResult `json:"characters"`
	Stats      Stats             `json:"stats"`
}

type GameDao struct {
	client *elastic.Client
	index  string
}

func NewGameDao(client *elastic.Client) *GameDao {
	return &GameDao{client, "ow"}
}

func (g *GameDao) GetByUser(ctx context.Context, id int) ([]Game, error) {
	query := elastic.NewTermQuery("user_id", id)
	resp, err := g.client.Search().Index(g.index).Type("games").Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}
	var games []Game
	for _, item := range resp.Each(reflect.TypeOf(Game{})) {
		var g Game
		var ok bool
		if g, ok = item.(Game); !ok {
			return nil, errors.New("Trouble deserializing Game")
		}
		games = append(games, g)
	}
	return games, nil
}
