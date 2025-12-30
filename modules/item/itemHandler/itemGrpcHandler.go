package itemHandler

import (
	"context"

	itemPb "github.com/TewApirat/items-shop-ms/modules/item/itemPb"
	"github.com/TewApirat/items-shop-ms/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
		
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService)*itemGrpcHandler{
	return &itemGrpcHandler{
		itemUsecase: itemUsecase}
}

func (g *itemGrpcHandler) FindItemInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}