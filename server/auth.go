package server

import (
	"log"

	"github.com/TewApirat/items-shop-ms/modules/auth/authHandler"
	authPb "github.com/TewApirat/items-shop-ms/modules/auth/authPb"
	"github.com/TewApirat/items-shop-ms/modules/auth/authRepository"
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
	"github.com/TewApirat/items-shop-ms/pkg/grpccon"
)

func (s *server) authService(){
	repo := authRepository.NewAuthRepositoryService(s.db)
	usecase := authUsecase.NewAuthUasecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func ()  {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.App.Url)
		grpcServer.Serve(lis)

	}()


	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
