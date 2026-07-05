package server

import (
	"context"
	"log/slog"
	"pipeline"
	"pipeline/metrics"

	"github.com/rs/zerolog"

	"api/internal/api/handler"
	"api/internal/auth"
	"api/internal/config"
	"api/internal/database"
	"api/internal/event"
	"api/internal/health"
	"api/observer"
)

type Application struct {
	Config        *config.Config
	Logger        zerolog.Logger
	ClickHouse    *database.ClickHouse
	HealthHandler *handler.HealthHandler
	JWTManager    *auth.JWTManager
	AuthService   auth.Service
	EventService  event.Service

	runtime *pipeline.Runtime[event.Event]
}

func New(
	cfg *config.Config,
	logger zerolog.Logger,
	db *database.ClickHouse,
) *Application {
	jwtManager := auth.NewJWTManager(
		cfg.JWT.Secret,
		cfg.JWT.Expiry,
	)
	userRepository := auth.NewClickHouseRepository(db)
	authService := auth.NewService(
		userRepository,
		jwtManager,
	)

	eventRepository := event.NewClickHouseRepository(
		db,
	)

	pipelines := pipeline.New[event.Event](
		pipeline.DefaultConfig(),
	)

	pipelines.
		Map("validate", validateEvent).
		Map("normalize", normalizeEvent).
		Map("enrich", enrichEvent).
		Filter("remove-invalid", isValidEvent).
		Generate("derived-events", createDerivedEvents).
		WithObserver(observer.NewClickHouseObserver(logger, observer.NewClickHousePipelineMetricsRepository(db)))

	executor := pipelines.Executor()
	runtime := pipeline.NewRuntime(pipeline.DefaultConfig(), executor, NewClickHouseSink(eventRepository))
	runtime.Start(context.Background())

	eventService := event.NewService(
		*slog.Default(),
		runtime,
	)

	metrics.Register()
	return &Application{
		Config:     cfg,
		Logger:     logger,
		ClickHouse: db,
		HealthHandler: handler.NewHealthHandler(
			health.New(db),
		),

		JWTManager:   jwtManager,
		AuthService:  authService,
		EventService: eventService,

		runtime: runtime,
	}
}

func (a *Application) Close() error {
	a.runtime.Stop()
	return a.ClickHouse.Close()
}

type ClickHouseSink struct {
	repository event.Repository
}

func NewClickHouseSink(repository event.Repository) *ClickHouseSink {
	return &ClickHouseSink{
		repository: repository,
	}
}

func (s *ClickHouseSink) Drain(
	ctx context.Context,
	events []event.Event,
) error {
	return s.repository.InsertBatch(
		ctx,
		events,
	)
}
