package observer

import (
	"api/internal/database"
	"context"
	"database/sql"
	"pipeline"
)

type PipelineMetricsRepository interface {
	InsertStageMetric(
		context.Context,
		pipeline.StageMetrics,
	) error

	InsertExecution(
		context.Context,
		pipeline.ExecutionMetadata,
	) error
}

type ClickHousePipelineMetricsRepository struct {
	db *sql.DB
}

func NewClickHousePipelineMetricsRepository(
	ch *database.ClickHouse,
) *ClickHousePipelineMetricsRepository {

	return &ClickHousePipelineMetricsRepository{
		db: ch.DB,
	}
}

func (r *ClickHousePipelineMetricsRepository) InsertExecution(
	ctx context.Context,
	meta pipeline.ExecutionMetadata,
) error {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(
		ctx,
		`
		INSERT INTO pipeline_stage_metrics
		(
			execution_id,
			stage,
			started_at,
			finished_at,
			latency_ms,
			input,
			output,
			dropped,
			errors,
			throughput
		)
		VALUES
		(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
		`,
	)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	defer stmt.Close()

	for _, stage := range meta.Stages {

		_, err = stmt.ExecContext(
			ctx,
			meta.ID,
			stage.Name,
			stage.StartedAt,
			stage.FinishedAt,
			float64(stage.Latency.Microseconds())/1000.0,
			stage.Input,
			stage.Output,
			stage.Dropped,
			stage.Errors,
			stage.Throughput,
		)

		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *ClickHousePipelineMetricsRepository) InsertStageMetric(
	context.Context,
	pipeline.StageMetrics,
) error {
	return nil
}
