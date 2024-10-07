package tokens

import (
	"auth/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateAccessToken(cfg *config.Config, id, role string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(cfg.ACCESS_TOKEN_KEY))
	if err != nil {
		return "", errors.Wrap(err, "access token generation failed")
	}

	return newToken, nil
}

func ValidateAccessToken(cfg *config.Config, token string) (bool, error) {
	_, err := ExtractAccessTokenClaims(cfg, token)
	if err != nil {
		return false, errors.Wrap(err, "validation failed")
	}
	return true, nil
}

func ExtractAccessTokenClaims(cfg *config.Config, tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.ACCESS_TOKEN_KEY), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "access token parsing failed")
	}
	if !token.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
