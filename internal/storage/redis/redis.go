package redis

import (
	"auth/internal/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	db       *redis.Client
	tokenTTL time.Duration
	codeTTL  time.Duration
	prefix   string
}

func NewRedisDB(cfg *config.Config) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_ADDRESS,
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DB,
	})

	return &RedisDB{
		db:       rdb,
		tokenTTL: cfg.REDIS_TOKEN_TTL,
		codeTTL:  cfg.REDIS_CODE_TTL,
		prefix:   cfg.REDIS_PREFIX,
	}
}

func (r *RedisDB) Close() {
	if err := r.db.Close(); err != nil {
		log.Printf("error while closing the redis connection: %v", err)
	}
}
