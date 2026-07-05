package event

import (
	"time"

	"api/internal/dto"
)

func FromDTO(d dto.TrackEvent) (Event, error) {
	ts, err := time.Parse(
		time.RFC3339Nano,
		d.Timestamp,
	)
	if err != nil {
		return Event{}, err
	}

	return Event{
		ID: d.ID,

		Type: d.Type,

		Name: d.Name,

		UserID: d.UserID,

		SessionID: d.SessionID,

		Timestamp: ts,

		PageURL: d.URL,

		Path: d.Page,

		Properties: d.Properties,

		Metadata: Metadata{
			Referrer: d.Referrer,
		},
	}, nil
}
