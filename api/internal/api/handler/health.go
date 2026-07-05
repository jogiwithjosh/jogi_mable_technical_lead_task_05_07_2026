package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/health"
)

type HealthHandler struct {
	service *health.Service
}

func NewHealthHandler(service *health.Service) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) Health(c *gin.Context) {
	if err := h.service.Ready(); err != nil {

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "down",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
