package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Register(c *gin.Context) {}

func (h *Handler) Login(c *gin.Context) {}

func (h *Handler) ForgotPassword(c *gin.Context) {}

func (h *Handler) ResetPassword(c *gin.Context) {}

func (h *Handler) RefreshToken(c *gin.Context) {}

func (h *Handler) ValidateToken(c *gin.Context) {}

func (h *Handler) Logout(c *gin.Context) {}
