package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"

	"api/internal/api/handler"
	"api/internal/api/middleware"
	"api/internal/auth"
	"api/internal/config"
	"api/internal/event"
)

func NewRouter(
	health *handler.HealthHandler,
	logger zerolog.Logger,
	cfg *config.Config,
	jwtManager *auth.JWTManager,
	authService auth.Service,
	eventService event.Service,
) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Recovery())
	router.Use(middleware.Logger(logger))
	router.Use(middleware.CORS(cfg))

	router.GET("/health", health.Health)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	authHandler := handler.NewAuthHandler(cfg, authService)
	orderHandler := handler.NewOrderHandler()

	eventHandler := handler.NewEventHandler(eventService)

	api := router.Group("/uapi")
	{
		api.POST(
			"/signup",
			authHandler.Signup,
		)

		api.POST(
			"/login",
			authHandler.Login,
		)
	}

	protected := router.Group("/api")

	protected.Use(
		middleware.RequireAuth(jwtManager, authService),
	)

	protected.GET(
		"/me",
		authHandler.Me,
	)

	protected.POST(
		"/logout",
		authHandler.Logout,
	)

	protected.POST(
		"/order",
		orderHandler.Order,
	)

	protected.POST(
		"/events",
		eventHandler.Track,
	)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	return router
}
