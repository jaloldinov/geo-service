package storage

import (
	"context"
	"geo/models"
	"time"
)

type StorageI interface {
	User() UsersI
	Message() MessageI
}

type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
	GetUserByEmail(context.Context, *models.LoginUserReq) (*models.User, error)
}

type MessageI interface {
	CreateMessage(context.Context, *models.Message) (string, error)
}
