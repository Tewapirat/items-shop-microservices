package authrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (

	AuthRepositoryService interface{}

	authrepository struct {
		db *mongo.Client
	}
)

func NewAuthRepositoryService(db *mongo.Client)AuthRepositoryService {
	return &authrepository{db}
}


func (r * authrepository)authDbConn(pctx context.Context) *mongo.Database{
	return r.db.Database("auth_db")
}