package event

import (
	"context"
	"time"
)

type Event struct {
	ID string

	Type string

	Name string

	UserID string

	SessionID string

	Timestamp time.Time

	PageURL string

	Path string

	Title string

	Properties map[string]any

	Metadata Metadata
}

type RequestMetadata struct {
	IPAddress string

	UserAgent string

	Referrer string

	Language string

	RequestID string

	SDKVersion string
}

type Metadata struct {
	IPAddress string

	UserAgent string

	Referrer string

	Language string

	Country string

	City string

	SessionID string

	RequestID string

	SDKVersion string
}

type metadataKey struct{}

func WithRequestMetadata(
	ctx context.Context,

	meta RequestMetadata,
) context.Context {
	return context.WithValue(

		ctx,

		metadataKey{},

		meta,
	)
}

func RequestMetadataFromContext(
	ctx context.Context,
) (RequestMetadata, bool) {
	value := ctx.Value(
		metadataKey{},
	)

	meta, ok := value.(RequestMetadata)

	return meta, ok
}
