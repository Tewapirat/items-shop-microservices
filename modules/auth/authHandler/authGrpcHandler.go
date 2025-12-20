package authHandler

import (
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authUseCase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUseCase authUsecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{authUseCase}
}