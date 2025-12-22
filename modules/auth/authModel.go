package auth

import (
	"time"

	"github.com/TewApirat/items-shop-ms/modules/player"
)

type (
	PLayerLogingReq struct {
		Email 		string `json:"email" form:"email" validate:"required,email,max=255"`
		Password 	string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenReq struct {
		RefreshToken string `json:"refresh_token" from:"refresh_token" validate:"required,max=500"`
	}

	InsertPlayerRole struct {
		PlayerId string `json:"player_id" validate:"required"`
		RoleCode []int 	`json:"role_id" validare:"required"`
	}

	ProfileIntercepter struct {
		*player.PlayerProfile
		Credential *Credential `json:"credentail"`
	}

	CredentialRes struct{
		Id 			 string				`json:"_id""`
		PlayerId 	 string 			`json:"player_id"`
		RoleCode 	 int 				`json:"role_code"`
		AccessToken  string 			`json:"access_token"`
		RefreshToken string 			`json:"refresh_token"`
		CreateAt 	 time.Time 			`json:"created_at"`
		UpdateAt 	 time.Time 			`json:"updated_at"`

	}
)