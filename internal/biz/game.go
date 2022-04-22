package biz

import (
	"context"

	v1 "game-ddz/api/game/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// game is a game model.
type game struct {
	Hello string
}

// gameRepo is a Greater repo.
type gameRepo interface {
	Save(context.Context, *game) (*game, error)
	Update(context.Context, *game) (*game, error)
	FindByID(context.Context, int64) (*game, error)
	ListByHello(context.Context, string) ([]*game, error)
	ListAll(context.Context) ([]*game, error)
}

// gameUsecase is a game usecase.
type gameUsecase struct {
	repo gameRepo
	log  *log.Helper
}

// NewgameUsecase new a game usecase.
func NewgameUsecase(repo gameRepo, logger log.Logger) *gameUsecase {
	return &gameUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Creategame creates a game, and returns the new game.
func (uc *gameUsecase) Creategame(ctx context.Context, g *game) (*game, error) {
	uc.log.WithContext(ctx).Infof("Creategame: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
