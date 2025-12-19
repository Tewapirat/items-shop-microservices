package itemHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	ItemHttpHandler struct{
		cfg 			*config.Config
		itemUsecase 	itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService)ItemHttpHandlerService{
	return &ItemHttpHandler{cfg,itemUsecase}
}

