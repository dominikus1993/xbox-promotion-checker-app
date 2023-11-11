package games

import "context"

type XboxGameQuery struct {
	Page     int
	PageSize int
}

type XboxGamesReader interface {
	Read(ctx context.Context, q XboxGameQuery) ([]XboxStoreGame, error)
}
