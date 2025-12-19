package inventoryHandler

import (

	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct{
		cfg 			*config.Config
		inventoryUsecase 	inventoryUsecase.InventoryUsecaseService
		
	}
)

func NewInventoryQueueHandler(	cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService)InventoryQueueHandlerService{
	return &inventoryQueueHandler{cfg,inventoryUsecase}
}