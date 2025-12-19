package authhandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface {}

	authHttpHandler struct {
		cfg 		*config.Config
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthHandler(cfg * config.Config, authUsecase authusecase.AuthUsecaseService)AuthHttpHandlerService{
	return &authHttpHandler{cfg, authUsecase}
}