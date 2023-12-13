package storage

import (
	"context"
	"geo/models"
)

type StorageI interface {
	User() UsersI
	Message() MessageI
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
	GetUserByEmail(context.Context, *models.LoginUserReq) (*models.User, error)
}

type MessageI interface {
	CreateMessage(context.Context, *models.Message) (string, error)
}

/*
type StorageI interface {
	User() UsersI
	Message() MessageI
	UserGeo() UserGeoI
}

type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj models.Location, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (interface{}, error)
	Delete(ctx context.Context, key string) error
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRespond, error)
	GetUserByID(context.Context, *models.IdRequest) (*models.User, error)
}

type UserGeoI interface {
	GetUserGeo(context.Context, *models.GetUserGeoRequest) (*models.UserGeoRespond, error)
}

/*
type MessageI interface {
	CreateMessage(context.Context, *models.Message) (string, error)
}
*/
