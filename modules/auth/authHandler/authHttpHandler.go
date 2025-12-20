package authHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface {}

	authHttpHandler struct {
		cfg 		*config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg * config.Config, authUsecase authUsecase.AuthUsecaseService)AuthHttpHandlerService{
	return &authHttpHandler{cfg, authUsecase}
}