package authhandler

import (
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authUseCase authusecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUseCase authusecase.AuthUsecaseService)authusecase.AuthUsecaseService{
	return &authGrpcHandler{authUseCase}
}