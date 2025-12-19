package inventoryUsecase

import (
		"github.com/TewApirat/items-shop-ms/modules/inventory/inventoryRepository"
)

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService)InventoryUsecaseService{
	return &inventoryUsecase{inventoryRepository}
}
