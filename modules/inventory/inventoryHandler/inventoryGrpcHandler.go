package inventoryHandler

import (
		"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct{
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{inventoryUsecase}
}