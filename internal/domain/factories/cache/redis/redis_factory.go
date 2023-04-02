package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wikimedia/internal/application/ports"
	kv "github.com/wikimedia/internal/infrastructure/adapters/cache/redis"
	"github.com/wikimedia/pkg/config"
)

// DefaultCacheTtl Hold the cache for 10 minutes
var (
	DefaultCacheTtl        = 5 * time.Minute
	DefaultRedisHostEnvVar = "REDIS_HOST"
	DefaultRedisPortEnvVar = "REDIS_PORT"
	DefaultRedisPassEnvVar = "REDIS_PASS"
)

type RedisFactory struct {
	Cfg *config.Config
}

func New(c *config.Config) ports.CacheFactory {
	return &RedisFactory{Cfg: c}
}

func (r RedisFactory) GetCache() (ports.Cache, error) {
	port, _ := strconv.Atoi(os.Getenv(DefaultRedisPortEnvVar))
	url := fmt.Sprintf(r.Cfg.Cache.Redis.URL, os.Getenv(DefaultRedisHostEnvVar), port)
	pass := fmt.Sprintf(r.Cfg.Cache.Redis.Pass, os.Getenv(DefaultRedisPassEnvVar))
	opt := &redis.Options{
		Addr:     url,
		Password: pass,
		DB:       r.Cfg.Cache.Redis.DB, // Default DB
	}
	ctx := context.Background()
	redisClient := redis.NewClient(opt)
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	cache := &kv.Cache{
		Ctx:    ctx,
		Client: redisClient,
		Ttl:    DefaultCacheTtl,
	}
	return cache, nil
}
