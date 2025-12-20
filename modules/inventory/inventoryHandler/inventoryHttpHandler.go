package inventoryHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryUsecase"
)

type (
		InventoryHttpHandlerService interface{}
		
		inventoryHttpHandler struct{
			cfg 	*config.Config
			inventoryUsecase inventoryUsecase.InventoryUsecaseService
		}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService)InventoryHttpHandlerService{
	return &inventoryHttpHandler{cfg,inventoryUsecase}
}