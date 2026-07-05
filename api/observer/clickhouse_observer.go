package observer

import (
	"context"
	"pipeline"

	"github.com/rs/zerolog"
)

type ClickHouseObserver struct {
	logger zerolog.Logger
	repo   PipelineMetricsRepository
}

func NewClickHouseObserver(logger zerolog.Logger, repo PipelineMetricsRepository) pipeline.Observer {
	return &ClickHouseObserver{
		logger: logger,
		repo:   repo,
	}
}

func (o *ClickHouseObserver) OnExecutionStart(
	ctx context.Context,
) {
}

func (o *ClickHouseObserver) OnStageComplete(
	ctx context.Context,
	m pipeline.StageMetrics,
) {

	err := o.repo.InsertStageMetric(ctx, m)
	if err != nil {
		o.logger.Error().Str("method", "OnStageComplete").Msg(err.Error())
	}
}

func (o *ClickHouseObserver) OnExecutionFinish(
	ctx context.Context,
	meta pipeline.ExecutionMetadata,
) {

	err := o.repo.InsertExecution(ctx, meta)
	if err != nil {
		o.logger.Error().Str("method", "OnExecutionFinish").Msg(err.Error())
	}
}
