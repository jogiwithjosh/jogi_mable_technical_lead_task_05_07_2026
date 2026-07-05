package server

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"api/internal/event"

	"github.com/google/uuid"
)

func validateEvent(
	ctx context.Context,
	e event.Event,
) (event.Event, error) {
	if strings.TrimSpace(e.UserID) == "" {
		return e, errors.New("missing user id")
	}

	if strings.TrimSpace(e.SessionID) == "" {
		return e, errors.New("missing session id")
	}

	if strings.TrimSpace(e.Name) == "" {
		return e, errors.New("missing event name")
	}

	if strings.TrimSpace(e.Type) == "" {
		return e, errors.New("missing event type")
	}

	return e, nil
}

func normalizeEvent(
	ctx context.Context,
	e event.Event,
) (event.Event, error) {
	e.Name = strings.TrimSpace(e.Name)
	e.PageURL = strings.TrimSpace(e.PageURL)
	e.Path = strings.TrimSpace(e.Path)
	e.Title = strings.TrimSpace(e.Title)
	if e.PageURL != "" {
		if parsed, err := url.Parse(e.PageURL); err == nil {
			parsed.Fragment = ""
			e.PageURL = parsed.String()
		}
	}
	if e.Properties == nil {
		e.Properties = map[string]any{}
	}
	return e, nil
}

func enrichEvent(
	ctx context.Context,
	e event.Event,
) (event.Event, error) {
	meta, ok := event.RequestMetadataFromContext(
		ctx,
	)

	if !ok {
		return e, nil
	}

	e.Metadata.IPAddress = meta.IPAddress
	e.Metadata.UserAgent = meta.UserAgent
	e.Metadata.Referrer = meta.Referrer
	e.Metadata.Language = meta.Language
	e.Metadata.RequestID = meta.RequestID
	e.Metadata.SDKVersion = meta.SDKVersion

	return e, nil
}

func isValidEvent(
	ctx context.Context,
	e event.Event,
) bool {
	if e.UserID == "" {
		return false
	}

	if e.Type == "" {
		return false
	}

	if e.Timestamp.After(time.Now().Add(time.Minute)) {
		return false
	}

	return true
}

func createDerivedEvents(
	ctx context.Context,
	e event.Event,
) ([]event.Event, error) {
	derived := event.Event{
		ID:        uuid.NewString(),
		Type:      "pipeline.capture_time",
		Timestamp: time.Now().UTC(),
		UserID:    e.UserID,
		SessionID: e.SessionID,

		Properties: map[string]any{
			"source_event": e.ID,
			"event_type":   e.Type,
		},
	}

	return []event.Event{
		derived,
	}, nil
}
