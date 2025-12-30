package playerHandler

import (
	"context"

	playerPb "github.com/TewApirat/items-shop-ms/modules/player/playerPb"
	"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {

		playerUseCase playerUsecase.PlayerUsecaseService
		playerPb.UnimplementedPlayerGrpcServiceServer

	}
)

func NewPlayerGrpcHandler(playerUseCase playerUsecase.PlayerUsecaseService) *playerGrpcHandler{
	return &playerGrpcHandler{
		playerUseCase: playerUseCase,
	}
}

func(g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func(g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func(g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}