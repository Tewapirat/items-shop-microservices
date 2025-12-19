package middlewareusecase

import(
	"github.com/TewApirat/items-shop-ms/modules/middleware/middlewareRepository"
)

type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct{
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUasecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService)MiddlewareUsecaseService{
	return &middlewareUsecase{middlewareRepository}
}