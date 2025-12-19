package middlewarehandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct{
		cfg 				*config.Config
		middleWareUsecase 	middlewareusecase.MiddlewareUsecaseService
		
	}
)

func NewMiddlewareHandler(cfg *config.Config, middleWareUsecase middlewareusecase.MiddlewareUsecaseService)MiddlewareHandlerService{
	return &middlewareHandler{cfg, middleWareUsecase}
}