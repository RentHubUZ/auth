package handler

import (
	"auth/internal/api/email"
	"auth/internal/api/tokens"
	"auth/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(c *gin.Context) {
	h.Logger.Info("Register handler is invoked")

	var req models.RegisterData
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	if req.Password != req.ConfirmPassword {
		handleError(c, h, errors.New("passwords do not match"), "invalid data", http.StatusBadRequest)
		return
	}

	passByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handleError(c, h, err, "error hashing password", http.StatusInternalServerError)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	resp, err := h.Storage.User().Add(ctx, &models.RegisterReq{
		FullName:       req.FullName,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: string(passByte),
	})
	if err != nil {
		handleError(c, h, err, "error registering user", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Register handler is completed successfully")
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
	h.Logger.Info("Login handler is invoked")

	var req models.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	user, err := h.Storage.User().GetByEmail(ctx, req.Email)
	if err != nil {
		handleError(c, h, err, "error getting user details", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		handleError(c, h, err, "invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(h.Config, user.ID, user.Role)
	if err != nil {
		handleError(c, h, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := tokens.GenerateRefreshToken(ctx, h.Config, h.RedisDB, user.ID)
	if err != nil {
		handleError(c, h, err, "error generating refresh token", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Login handler is completed successfully")
	c.JSON(http.StatusOK, models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	h.Logger.Info("ForgotPassword handler is invoked")

	req := c.Query("email")
	if req == "" {
		handleError(c, h, errors.New("email is required"), "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	_, err := h.Storage.User().GetByEmail(ctx, req)
	if err != nil {
		handleError(c, h, err, "invalid credentials", http.StatusUnauthorized)
		return
	}

	err = email.SendEmail(ctx, h.Config, h.RedisDB, req)
	if err != nil {
		handleError(c, h, err, "error sending email", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("ForgotPassword handler is completed successfully")
	c.JSON(http.StatusOK, gin.H{"message": "check your email to reset password"})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	h.Logger.Info("ResetPassword handler is invoked")

	var req models.ResetPassReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		handleError(c, h, errors.New("all fields are required"), "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	user, err := h.Storage.User().GetByEmail(ctx, req.Email)
	if err != nil {
		handleError(c, h, err, "invalid credentials", http.StatusUnauthorized)
		return
	}

	err = h.Storage.User().UpdatePassword(ctx, user.ID, req.NewPassword)
	if err != nil {
		handleError(c, h, err, "error updating password", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("ResetPassword handler is completed successfully")
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	h.Logger.Info("RefreshToken handler is invoked")

	var req models.RefreshToken
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	valid, err := tokens.ValidateRefreshToken(ctx, h.Config, h.RedisDB, req.RefreshToken)
	if !valid || err != nil {
		handleError(c, h, err, "error validating refresh token", http.StatusInternalServerError)
		return
	}

	id, err := tokens.ExtractRefreshID(h.Config, req.RefreshToken)
	if err != nil {
		handleError(c, h, err, "error extracting user id from refresh token", http.StatusInternalServerError)
		return
	}

	role, err := h.Storage.User().GetRole(ctx, id)
	if err != nil {
		handleError(c, h, err, "error getting user role", http.StatusInternalServerError)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(h.Config, id, role)
	if err != nil {
		handleError(c, h, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("RefreshToken handler is completed successfully")
	c.JSON(http.StatusOK, models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	})
}

func (h *Handler) ValidateToken(c *gin.Context) {
	h.Logger.Info("ValidateToken handler is invoked")

	var req models.AccessToken
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	valid, err := tokens.ValidateAccessToken(h.Config, req.AccessToken)
	if !valid || err != nil {
		handleError(c, h, err, "error validating access token", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("ValidateToken handler is completed successfully")
	c.JSON(http.StatusOK, "Access token is valid")
}

func (h *Handler) Logout(c *gin.Context) {
	h.Logger.Info("Logout handler is invoked")

	email := c.Query("email")
	if email == "" {
		handleError(c, h, errors.New("email is required"), "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := makeContext(h)
	defer cancel()

	user, err := h.Storage.User().GetByEmail(ctx, email)
	if err != nil {
		handleError(c, h, err, "error getting user details", http.StatusInternalServerError)
		return
	}

	err = h.RedisDB.DeleteToken(ctx, user.ID)
	if err != nil {
		handleError(c, h, err, "error logging out", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Logout handler is completed successfully")
	c.JSON(http.StatusOK, "User logged out successfully")
}
