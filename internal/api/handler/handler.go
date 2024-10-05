package handler

import (
	"auth/internal/config"
	logger "auth/internal/logs"
	"auth/internal/storage"
	"auth/internal/storage/redis"
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	Storage        storage.IStorage
	RedisDB        *redis.RedisDB
	Logger         *slog.Logger
	Config         *config.Config
	ContextTimeout time.Duration
}

func NewHandler(s storage.IStorage, rdb *redis.RedisDB, cfg *config.Config) *Handler {
	return &Handler{
		Storage:        s,
		RedisDB:        rdb,
		Logger:         logger.NewLogger(),
		Config:         cfg,
		ContextTimeout: time.Second * 5,
	}
}

func makeContext(h *Handler, c context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c, h.ContextTimeout)
}

func handleError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Logger.Error(er)
}
