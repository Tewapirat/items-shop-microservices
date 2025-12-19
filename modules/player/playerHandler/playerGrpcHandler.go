package playerHandler

import (
	"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
)

type (
	playerGrpcHandlerService struct {
		PlayerUseCase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(PlayerUseCase playerUsecase.PlayerUsecaseService) *playerGrpcHandlerService{
	return &playerGrpcHandlerService{PlayerUseCase}
}