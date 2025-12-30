package authHandler

import (
	"context"

	authPb "github.com/TewApirat/items-shop-ms/modules/auth/authPb"
	"github.com/TewApirat/items-shop-ms/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUseCase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUseCase authUsecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{
		authUseCase: authUseCase,
	}
}

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RoleCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}