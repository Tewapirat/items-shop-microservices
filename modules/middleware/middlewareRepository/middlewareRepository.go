package middlewareRepository

type (
	MiddlewareRepositoryService interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepositoryHandler()MiddlewareRepositoryService{
	return &middlewareRepository{}
}