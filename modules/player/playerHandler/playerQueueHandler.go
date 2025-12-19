package playerHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct{
		cfg 			*config.Config
		playerUsecase 	playerUsecase.PlayerUsecaseService
		
	}
)

func NewPlayerQueueHandler(	cfg *config.Config,playerUsecase playerUsecase.PlayerUsecaseService)PlayerQueueHandlerService{
	return &playerQueueHandler{cfg,playerUsecase}
}