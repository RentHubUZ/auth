package redisdb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (r *RedisDB) StoreToken(ctx context.Context, userID, token string) error {
	if userID == "" || token == "" {
		return errors.New("invalid user id or token")
	}

	err := r.db.Set(ctx, r.prefix+userID, token, time.Hour*24).Err()
	if err != nil {
		return errors.Wrap(err, "failed to store token")
	}
	return nil
}

func (r *RedisDB) GetToken(ctx context.Context, userID string) (string, error) {
	if userID == "" {
		return "", errors.New("invalid user id")
	}

	token, err := r.db.Get(ctx, r.prefix+userID).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("token not found for " + userID)
		}
		return "", errors.Wrap(err, "failed to get token")
	}
	return token, nil
}

func (r *RedisDB) DeleteToken(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("invalid user id")
	}

	err := r.db.Del(ctx, r.prefix+userID).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete token")
	}
	return nil
}
