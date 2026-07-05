package handler

import (
	"errors"
	"net/http"

	"api/internal/api/middleware"
	"api/internal/auth"
	"api/internal/config"
	"api/internal/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	config      *config.Config
	authService auth.Service
}

func NewAuthHandler(
	config *config.Config,
	authService auth.Service,
) *AuthHandler {
	return &AuthHandler{
		config:      config,
		authService: authService,
	}
}

func (h *AuthHandler) Signup(
	c *gin.Context,
) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	user, err := h.authService.Signup(
		c.Request.Context(),
		req,
	)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(
			err,
			auth.ErrUserAlreadyExists,
		) {
			status = http.StatusConflict
		}

		c.JSON(status, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusCreated,
		user.Public(),
	)
}

func (h *AuthHandler) Login(
	c *gin.Context,
) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	user, token, err := h.authService.Login(
		c.Request.Context(),
		req,
	)
	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	middleware.SetAuthCookie(
		c.Writer,
		token,
		h.config.JWT,
	)

	c.JSON(
		http.StatusOK,
		gin.H{
			"user": user.Public(),
		},
	)
}

func (h *AuthHandler) Logout(
	c *gin.Context,
) {
	middleware.ClearAuthCookie(
		c.Writer,
	)

	c.Status(
		http.StatusNoContent,
	)
}

func (h *AuthHandler) Me(
	c *gin.Context,
) {
	user := middleware.User(c)

	if user == nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "unauthorized",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		user.Public(),
	)
}
