package tokens

import (
	"auth/internal/config"
	"auth/internal/storage/redis"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateRefreshToken(ctx context.Context, cfg *config.Config, rdb *redis.RedisDB, userID string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(cfg.REFRESH_TOKEN_KEY))
	if err != nil {
		return "", errors.Wrap(err, "refresh token generation failed")
	}

	err = rdb.StoreToken(ctx, userID, newToken)
	if err != nil {
		return "", errors.Wrap(err, "refresh token storage failed")
	}
	return newToken, nil
}

func ValidateRefreshToken(ctx context.Context, cfg *config.Config, rdb *redis.RedisDB, tokenStr string) (bool, error) {
	claims, err := ExtractRefreshClaims(cfg, tokenStr)
	if err != nil {
		return false, errors.Wrap(err, "validation failed")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return false, errors.New("invalid token claims: user id not found")
	}

	token, err := rdb.GetToken(ctx, userID)
	if err != nil {
		return false, errors.Wrap(err, "validation failed")
	}

	if token != tokenStr {
		return false, errors.New("invalid refresh token")
	}
	return true, nil
}

func ExtractRefreshID(cfg *config.Config, token string) (string, error) {
	claims, err := ExtractRefreshClaims(cfg, token)
	if err != nil {
		return "", errors.Wrap(err, "claims extraction failed")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token claims: user id not found")
	}
	return userID, nil
}

func ExtractRefreshClaims(cfg *config.Config, tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.REFRESH_TOKEN_KEY), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "refresh token parsing failed")
	}
	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
