package server

import(
		"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryRepository"
		"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryUsecase"
		"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryHandler"
)

func (s *server) inventoryService(){
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Check
	inventory.GET("", s.healthCheckService)
}