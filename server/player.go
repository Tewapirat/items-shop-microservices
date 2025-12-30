package server

import (
	"log"

	"github.com/TewApirat/items-shop-ms/modules/player/playerHandler"
	playerPb "github.com/TewApirat/items-shop-ms/modules/player/playerPb"
	"github.com/TewApirat/items-shop-ms/modules/player/playerRepository"
	"github.com/TewApirat/items-shop-ms/modules/player/playerUsecase"
	"github.com/TewApirat/items-shop-ms/pkg/grpccon"
)

func (s *server) playerService(){
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	// gRPC
	go func ()  {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.App.Url)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)
}