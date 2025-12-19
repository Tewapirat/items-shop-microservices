package playerHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct{
		cfg 			*config.Config
		playerUsecase 	playerUsecase.PlayerUsecaseService
		
	}
)

func NewPlayerHttpHandler(	cfg *config.Config,playerUsecase playerUsecase.PlayerUsecaseService)PlayerHttpHandlerService{
	return &playerHttpHandler{cfg,playerUsecase}
}