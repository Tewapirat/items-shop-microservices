package authUsecase

import (
	"github.com/TewApirat/items-shop-ms/modules/auth/authRepository"
)

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUasecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}