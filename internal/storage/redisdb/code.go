package redisdb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (r *RedisDB) StoreCode(ctx context.Context, email, code string) error {
	if email == "" || code == "" {
		return errors.New("invalid email or code")
	}

	err := r.db.Set(ctx, r.prefix+email, code, time.Minute*3).Err()
	if err != nil {
		return errors.Wrap(err, "failed to store code")
	}
	return nil
}

func (r *RedisDB) GetCode(ctx context.Context, email string) (string, error) {
	if email == "" {
		return "", errors.New("invalid email")
	}

	code, err := r.db.Get(ctx, r.prefix+email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("code not found for " + email)
		}
		return "", errors.Wrap(err, "failed to get code")
	}
	return code, nil
}

func (r *RedisDB) DeleteCode(ctx context.Context, email string) error {
	if email == "" {
		return errors.New("invalid email")
	}

	err := r.db.Del(ctx, r.prefix+email).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete code")
	}
	return nil
}
