package server

import (
	"log"

	"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryHandler"
	inventoryPb "github.com/TewApirat/items-shop-ms/modules/inventory/inventoryPb"
	"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryRepository"
	"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryUsecase"
	"github.com/TewApirat/items-shop-ms/pkg/grpccon"
	
)

func (s *server) inventoryService(){
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)


	// gRPC
	go func ()  {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)
		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.App.Url)
		grpcServer.Serve(lis)

	}()


	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Check
	inventory.GET("", s.healthCheckService)
}