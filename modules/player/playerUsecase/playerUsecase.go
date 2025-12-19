package playerUsecase

import (
	"github.com/TewApirat/items-shop-ms/modules/player/playerRepository"
)

type (
	PlayerUsecaseService interface{}

	playerUsecase struct{
		playerRepository playerRepository.PlayerRepositoryService
		
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService)PlayerUsecaseService{
	return &playerUsecase{playerRepository}
}