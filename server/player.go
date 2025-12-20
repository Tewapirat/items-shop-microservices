package server

import(
		"github.com/TewApirat/items-shop-ms/modules/player/playerRepository"
		"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
		"github.com/TewApirat/items-shop-ms/modules/player/playerHandler"
)

func (s *server) playerService(){
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	_ = player
}