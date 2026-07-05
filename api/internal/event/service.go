package event

import (
	"context"
	"log/slog"
	"pipeline"
)

type Service interface {
	Track(
		ctx context.Context,
		event Event,
	) error

	TrackBatch(
		ctx context.Context,
		events []Event,
	) error
}

type EventService struct {
	logger  slog.Logger
	runtime *pipeline.Runtime[Event]
}

func NewService(
	logger slog.Logger,
	runtime *pipeline.Runtime[Event],
) Service {
	return &EventService{
		logger:  logger,
		runtime: runtime,
	}
}

func (s *EventService) Track(
	ctx context.Context,
	event Event,
) error {
	if err := s.runtime.Publish(event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) TrackBatch(
	ctx context.Context,
	events []Event,
) error {
	if len(events) == 0 {
		return nil
	}

	for _, event := range events {
		if err := s.runtime.Publish(event); err != nil {
			return err
		}
	}

	return nil
}
