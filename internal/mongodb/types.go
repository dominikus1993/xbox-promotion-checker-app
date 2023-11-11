package mongodb

import (
	"strconv"

	"github.com/dominikus1993/xbox-promotion-checker-app/pkg/games"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type screapedXboxHame struct {
	ID        string             `bson:"_id,omitempty"`
	Title     string             `bson:"Title"`
	Link      string             `bson:"Link"`
	Price     string             `bson:"Price"`
	OldPrice  string             `bson:"OldPrice"`
	Promotion string             `bson:"Promotion"`
	CrawledAt primitive.DateTime `bson:"CrawledAt"`
}

func parsePrice(price string) float64 {
	if s, err := strconv.ParseFloat(price, 64); err == nil {
		return s
	}
	return 0
}

func toXboxGame(game screapedXboxHame) games.XboxStoreGame {
	return games.XboxStoreGame{
		ID:       game.ID,
		Title:    games.Title(game.Title),
		Link:     games.Link(game.Link),
		Price:    parsePrice(game.Price),
		OldPrice: parsePrice(game.OldPrice),
	}
}
