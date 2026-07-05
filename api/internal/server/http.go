package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"api/internal/api"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTP(app *Application) *HTTPServer {
	gin.SetMode(gin.ReleaseMode)

	router := api.NewRouter(app.HealthHandler, app.Logger, app.Config, app.JWTManager, app.AuthService, app.EventService)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.Config.Port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &HTTPServer{
		server: srv,
	}
}

func (h *HTTPServer) Start() error {
	return h.server.ListenAndServe()
}

func (h *HTTPServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
