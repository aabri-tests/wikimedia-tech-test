package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Ctx    context.Context
	Client *redis.Client
	Ttl    time.Duration
}

func (r *Cache) Get(key string) ([]byte, error) {
	data, err := r.Client.Get(r.Ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Cache) Set(key string, data []byte) error {
	if r.Ttl < 0 {
		return errors.New("TTL must be a non-negative duration")
	}

	err := r.Client.Set(r.Ctx, key, data, r.Ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Cache) Delete(key string) error {
	err := r.Client.Del(r.Ctx, key).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	return nil
}
