package inventory

import (
	"github.com/TewApirat/items-shop-ms/modules/item"
	"github.com/TewApirat/items-shop-ms/modules/models"
)

type(
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId string `json:"item_id" validate:"required,max=64"`

	}

	ItemInventory struct{
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlayerInventory struct{
		PlayerId string `json:"player_id"`
		*models.PaginateRes 

	}
)