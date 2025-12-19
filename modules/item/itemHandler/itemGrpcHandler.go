package itemHandler

import (
	"github.com/TewApirat/items-shop-ms/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		ItemUsecase itemUsecase.ItemUsecaseService
		
	}
)

func NewItemGrpcHandler(ItemUsecase itemUsecase.ItemUsecaseService)*itemGrpcHandler{
	return &itemGrpcHandler{ItemUsecase}
}