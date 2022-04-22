package service

import (
	"context"

	v1 "game-ddz/api/game/v1"
	"game-ddz/internal/biz"
)

// gameService is a game service.
type gameService struct {
	v1.UnimplementedgameServer

	uc *biz.gameUsecase
}

// NewgameService new a game service.
func NewgameService(uc *biz.gameUsecase) *gameService {
	return &gameService{uc: uc}
}

// SayHello implements game.gameServer.
func (s *gameService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.Creategame(ctx, &biz.game{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
