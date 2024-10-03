package api

import (
	_ "auth/internal/api/docs"
	"auth/internal/api/handler"
	"auth/internal/config"
	"auth/internal/storage"
	"auth/internal/storage/redis"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authorazation
// @version 1.0
// @description Authorazation API of RentHub
// @host localhost:8081
// @BasePath /auth
func NewRouter(s storage.IStorage, rdb *redis.RedisDB, cfg *config.Config) *gin.Engine {
	h := handler.NewHandler(s, rdb, cfg)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/forgot-password", h.ForgotPassword)
		auth.POST("/reset-password", h.ResetPassword)
		auth.POST("/refresh-token", h.RefreshToken)
		auth.POST("/validate-token", h.ValidateToken)
		auth.POST("/logout", h.Logout)
	}
	return r
}
