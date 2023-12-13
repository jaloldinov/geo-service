package redis

import (
	"context"
	"fmt"
	"geo/models"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/pkg/errors"
)

type cacheRepo struct {
	cache *cache.Cache
}

func NewCacheRepo(cache *cache.Cache) *cacheRepo {
	return &cacheRepo{cache: cache}
}

func (u cacheRepo) Create(ctx context.Context, id string, obj models.Location, ttl time.Duration) error {
	//first check if that id exists or not, if yes append otherwise create set new data to redis
	// Retrieve the existing array from Redis
	var existingArray []models.Location
	err := u.cache.Get(ctx, id, &existingArray)
	if err != nil && err != cache.ErrCacheMiss {
		//println("redis.Create.Error:", err.Error(), "\nkey:", id)
		fmt.Printf("id not found in redis. new data creating...")
		err = u.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   id,
			Value: obj,
			TTL:   ttl,
		})
		if err != nil {
			println("redis.Create.Error:", err.Error(), "\nkey:", id)
			return errors.Wrap(err, "error while creating cache in redis")
		}
	}

	// Append the new data to the existing array
	existingArray = append(existingArray, obj)

	// Store the updated array in Redis
	err = u.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id,
		Value: existingArray,
		TTL:   ttl,
	})
	if err != nil {
		println("redis.Create.Error:", err.Error(), "\nkey:", id)
		return errors.Wrap(err, "error while appending array cache in redis")
	}

	fmt.Println("created/append in redis", id)
	return nil
}

func (u cacheRepo) Get(ctx context.Context, id string, response interface{}) (interface{}, error) {
	// var response interface{}

	err := u.cache.Get(ctx, id, response)
	if err != nil {
		println("redis.Get.Error:", err.Error(), "\nkey:", id)
		return false, err
	}
	fmt.Println("get from redis", id)
	return response, nil
}

func (u cacheRepo) Delete(ctx context.Context, id string) error {

	err := u.cache.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "error while deleting cache in redis")
	}
	fmt.Printf("delete from redis %s", id)
	return nil
}
