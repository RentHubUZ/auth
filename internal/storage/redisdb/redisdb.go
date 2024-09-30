package redisdb

import (
	"auth/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	db     *redis.Client
	prefix string
}

func NewRedisDB(cfg *config.Config) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_ADDRESS,
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DB,
	})

	return &RedisDB{
		db:     rdb,
		prefix: "rent-hub:",
	}
}

func (r *RedisDB) Close() {
	if err := r.db.Close(); err != nil {
		log.Printf("error while closing the redis connection: %v", err)
	}
}
