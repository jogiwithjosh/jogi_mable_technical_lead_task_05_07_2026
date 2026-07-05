package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api/internal/config"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/server"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg)

	db, err := database.New(cfg.ClickHouse)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to connect to ClickHouse")
	}

	app := server.New(
		cfg,
		log,
		db,
	)
	httpServer := server.NewHTTP(app)

	log.Info().
		Msg("ClickHouse connected")

	go func() {
		if err := httpServer.Start(); err != nil &&
			err != http.ErrServerClosed {

			app.Logger.Fatal().
				Err(err).
				Msg("HTTP server failed")
		}
	}()

	app.Logger.Info().
		Str("port", cfg.Port).
		Msg("API started")

	waitForShutdown(app, httpServer)
}

func waitForShutdown(
	app *server.Application,
	httpServer *server.HTTPServer,
) {
	// Wait for Ctrl+C or docker stop
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	app.Logger.Info().
		Msg("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		app.Logger.Error().
			Err(err).
			Msg("HTTP shutdown failed")
	}

	if err := app.Close(); err != nil {
		app.Logger.Error().
			Err(err).
			Msg("Database shutdown failed")
	}

	app.Logger.Info().
		Msg("Shutdown complete")
}
