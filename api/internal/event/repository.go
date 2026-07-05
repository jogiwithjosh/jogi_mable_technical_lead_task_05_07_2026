package event

import (
	"context"
	"database/sql"
	"encoding/json"
	"pipeline/metrics"
	"time"

	"api/internal/database"
)

type Repository interface {
	Insert(
		ctx context.Context,
		event Event,
	) error

	InsertBatch(
		ctx context.Context,
		events []Event,
	) error

	Count(
		ctx context.Context,
	) (uint64, error)
}

type ClickHouseRepository struct {
	db *sql.DB
}

func NewClickHouseRepository(
	ch *database.ClickHouse,
) Repository {
	return &ClickHouseRepository{
		db: ch.DB,
	}
}

func (r *ClickHouseRepository) Insert(
	ctx context.Context,
	event Event,
) error {
	properties, err := json.Marshal(
		event.Properties,
	)
	if err != nil {
		return err
	}

	query := `
INSERT INTO events
(
    event_id,
    event_type,
    event_name,
    user_id,
    session_id,
    captured_at,
    page_url,
    page_path,
    page_title,
    properties,
    ip_address,
    country,
    city,
    user_agent,
    referrer,
    language,
    sdk_version,
    request_id
)
VALUES
(
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

	_, err = r.db.ExecContext(
		ctx,
		query,

		event.ID,
		event.Type,
		event.Name,
		event.UserID,
		event.SessionID,
		event.Timestamp,
		event.PageURL,
		event.Path,
		event.Title,
		string(properties),

		event.Metadata.IPAddress,
		event.Metadata.Country,
		event.Metadata.City,
		event.Metadata.UserAgent,
		event.Metadata.Referrer,
		event.Metadata.Language,
		event.Metadata.SDKVersion,
		event.Metadata.RequestID,
	)

	return err
}

func (r *ClickHouseRepository) InsertBatch(
	ctx context.Context,
	events []Event,
) error {
	if len(events) == 0 {
		return nil
	}

	batch, err := r.db.BeginTx(
		ctx,
		nil,
	)
	if err != nil {
		return err
	}
	start := time.Now()

	stmt, err := batch.PrepareContext(
		ctx,
		`
INSERT INTO events
(
event_id,
event_type,
event_name,
user_id,
session_id,
captured_at,
page_url,
page_path,
page_title,
properties,
ip_address,
country,
city,
user_agent,
referrer,
language,
sdk_version,
request_id
)
VALUES
(
?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`,
	)
	if err != nil {
		batch.Rollback()
		metrics.BatchInsertFailures.Inc()
		return err
	}

	defer stmt.Close()

	for _, event := range events {
		props, err := json.Marshal(
			event.Properties,
		)
		if err != nil {
			batch.Rollback()
			return err
		}

		_, err = stmt.ExecContext(
			ctx,

			event.ID,
			event.Type,
			event.Name,
			event.UserID,
			event.SessionID,
			event.Timestamp,
			event.PageURL,
			event.Path,
			event.Title,
			string(props),

			event.Metadata.IPAddress,
			event.Metadata.Country,
			event.Metadata.City,
			event.Metadata.UserAgent,
			event.Metadata.Referrer,
			event.Metadata.Language,
			event.Metadata.SDKVersion,
			event.Metadata.RequestID,
		)
		if err != nil {
			batch.Rollback()
			metrics.BatchInsertFailures.Inc()
			return err
		}
	}
	metrics.InsertLatency.Observe(
		time.Since(start).Seconds(),
	)

	return batch.Commit()
}

func (r *ClickHouseRepository) Count(
	ctx context.Context,
) (uint64, error) {
	var count uint64

	err := r.db.QueryRowContext(
		ctx,
		`SELECT count() FROM events`,
	).Scan(&count)

	return count, err
}
