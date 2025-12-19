package itemUsecase

import (
	"github.com/TewApirat/items-shop-ms/modules/item/itemRepository"
)

type (
	ItemUsecaseService interface{}

	itemUsecase struct{
		itemRepository itemRepository.ItemRepositoryService

	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService)ItemUsecaseService{
	return &itemUsecase{itemRepository}
}
