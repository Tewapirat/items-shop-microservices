package server

import (
	"log"

	"github.com/TewApirat/items-shop-ms/modules/item/itemHandler"
	itemPb "github.com/TewApirat/items-shop-ms/modules/item/itemPb"
	"github.com/TewApirat/items-shop-ms/modules/item/itemRepository"
	"github.com/TewApirat/items-shop-ms/modules/item/itemUsecase"
	"github.com/TewApirat/items-shop-ms/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func ()  {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.App.Url)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
