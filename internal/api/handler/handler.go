package handler

import (
	"auth/internal/config"
	"auth/internal/logger"
	"auth/internal/storage"
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	Storage        storage.IStorage
	Logger         *slog.Logger
	Config         *config.Config
	ContextTimeout time.Duration
}

func NewHandler(s storage.IStorage, cfg *config.Config) *Handler {
	return &Handler{
		Storage:        s,
		Logger:         logger.NewLogger(),
		Config:         cfg,
		ContextTimeout: time.Second * 5,
	}
}

func makeContext(h *Handler) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), h.ContextTimeout)
}

func handleError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Logger.Error(er)
}
