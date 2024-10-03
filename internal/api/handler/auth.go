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

// Register godoc
// @Summary Registers user
// @Description Registers a new user
// @Tags auth
// @Param user body models.RegisterData true "User data"
// @Success 200 {object} models.RegisterResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /register [post]
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

// Login godoc
// @Summary Login user
// @Description Logs user in
// @Tags auth
// @Param user body models.LoginReq true "User data"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /login [post]
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

// ForgotPassword godoc
// @Summary Forgot password
// @Description Sends password reset email
// @Tags auth
// @Param email query string true "Email"
// @Success 200 {object} string "Check your email to reset password"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /forgot-password [post]
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

// ResetPassword godoc
// @Summary Resets password
// @Description Resets password
// @Tags auth
// @Param user body models.ResetPassReq true "User data"
// @Success 200 {object} string "Password reset successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reset-password [post]
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

// RefreshToken godoc
// @Summary Refreshes token
// @Description Refreshes refresh token
// @Tags auth
// @Param data body models.RefreshToken true "Refresh token"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /refresh [post]
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

// ValidateToken godoc
// @Summary Validates token
// @Description Validates access token
// @Tags auth
// @Param data body models.AccessToken true "Access token"
// @Success 200 {string} string "Access token is valid"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /validate [post]
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

// Logout godoc
// @Summary Logouts user
// @Description Logouts user
// @Tags auth
// @Param email query string true "User email"
// @Success 200 {string} string "User logged out successfully"
// @Failure 400 {object} string "Invalid email"
// @Failure 500 {object} string "Server error while processing request"
// @Router /logout [post]
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
