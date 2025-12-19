package authusecase

import (
	"github.com/TewApirat/items-shop-ms/modules/auth/authRepository"
)

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository authrepository.AuthRepositoryService
	}
)

func NewAuthUasecase(authRepository authrepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}