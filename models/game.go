package models

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
	elastic "gopkg.in/olivere/elastic.v5"
)

type Game struct {
	UserID     int               `json:"user_id",db:"user_id"`
	Map        string            `json:"map",db:"map"`
	StartTime  time.Time         `json:"start_time,string",db:"start_time"`
	GameType   GameType          `json:"game_type, string",db:"game_type"`
	GroupSize  int               `json:"group_size",db:"group_size"`
	IsSeason   bool              `json:"is_placement",db:"is_placement"`
	Season     int               `json:"season",db:"season"`
	EndSR      int               `json:"end_sr, omit_empty",db:"end_sr"`
	BeginSR    int               `json:"begin_sr,omit_empty",db:"begin_sr"`
	Result     Result            `json:"result,string",db:"result"`
	Characters []CharacterResult `json:"characters",db:"characters"`
	Stats      Stats             `json:"stats",db:"stats"`
}

type GameDao struct {
	client    *elastic.Client
	sqlClient *sqlx.DB
	index     string
}

func NewGameDao(client *elastic.Client) *GameDao {
	return &GameDao{client, nil, "ow"}
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
