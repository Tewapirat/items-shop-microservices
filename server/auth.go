package server

import(
		"github.com/TewApirat/items-shop-ms/modules/auth/authRepository"
		"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
		"github.com/TewApirat/items-shop-ms/modules/auth/authHandler"
)

func (s *server) authService(){
	repo := authRepository.NewAuthRepositoryService(s.db)
	usecase := authUsecase.NewAuthUasecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	_ = auth
}
