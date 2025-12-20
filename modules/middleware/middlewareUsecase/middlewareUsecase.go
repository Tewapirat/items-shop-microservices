package middlewareUsecase

import(
	"github.com/TewApirat/items-shop-ms/modules/middleware/middlewareRepository"
)

type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct{
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUasecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService)MiddlewareUsecaseService{
	return &middlewareUsecase{middlewareRepository}
}