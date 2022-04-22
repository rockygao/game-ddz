package data

import (
	"context"

	"game-ddz/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type gameRepo struct {
	data *Data
	log  *log.Helper
}

// NewgameRepo .
func NewgameRepo(data *Data, logger log.Logger) biz.gameRepo {
	return &gameRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gameRepo) Save(ctx context.Context, g *biz.game) (*biz.game, error) {
	return g, nil
}

func (r *gameRepo) Update(ctx context.Context, g *biz.game) (*biz.game, error) {
	return g, nil
}

func (r *gameRepo) FindByID(context.Context, int64) (*biz.game, error) {
	return nil, nil
}

func (r *gameRepo) ListByHello(context.Context, string) ([]*biz.game, error) {
	return nil, nil
}

func (r *gameRepo) ListAll(context.Context) ([]*biz.game, error) {
	return nil, nil
}
