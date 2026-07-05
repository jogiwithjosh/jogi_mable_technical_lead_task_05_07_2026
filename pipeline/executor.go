package pipeline

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Executor[T any] struct {
	stages   []Stage[T]
	observer Observer
}

func NewExecutor[T any](
	stages []Stage[T],
	observer Observer,
) *Executor[T] {
	return &Executor[T]{
		stages:   stages,
		observer: observer,
	}
}

func (e *Executor[T]) Execute(
	ctx context.Context,
	items []T,
) ([]T, ExecutionMetadata, error) {
	metadata := ExecutionMetadata{
		ID:        uuid.NewString(),
		StartedAt: time.Now(),
		Stages:    make([]StageMetrics, 0, len(e.stages)),
	}

	if e.observer != nil {
		e.observer.OnExecutionStart(ctx)
	}

	current := items

	for _, stage := range e.stages {

		stageStart := time.Now()

		result, err := stage.Process(ctx, current)

		stageMetrics := StageMetrics{
			ExecutionID: metadata.ID,
			Name:        stage.Name(),
			StartedAt:   stageStart,
			Input:       len(current),
		}

		stageMetrics.FinishedAt = time.Now()
		stageMetrics.Latency = stageMetrics.FinishedAt.Sub(stageStart)

		if err != nil {

			stageMetrics.Errors = 1
			stageMetrics.Output = 0
			stageMetrics.Dropped = len(current)

			if stageMetrics.Latency > 0 {
				stageMetrics.Throughput = CalculateThroughput(
					len(current),
					stageMetrics.Latency,
				)
			}

			metadata.Stages = append(
				metadata.Stages,
				stageMetrics,
			)

			metadata.FinishedAt = time.Now()
			metadata.Duration = metadata.FinishedAt.Sub(metadata.StartedAt)

			if e.observer != nil {
				e.observer.OnStageComplete(
					ctx,
					stageMetrics,
				)

				e.observer.OnExecutionFinish(
					ctx,
					metadata,
				)
			}

			return nil, metadata, err
		}

		stageMetrics.Output = len(result.Items)
		stageMetrics.Dropped = result.Dropped

		stageMetrics.Throughput = CalculateThroughput(
			len(result.Items),
			stageMetrics.Latency,
		)

		metadata.Stages = append(
			metadata.Stages,
			stageMetrics,
		)

		if e.observer != nil {
			e.observer.OnStageComplete(
				ctx,
				stageMetrics,
			)
		}

		current = result.Items
	}

	metadata.FinishedAt = time.Now()
	metadata.Duration = metadata.FinishedAt.Sub(metadata.StartedAt)

	if e.observer != nil {
		e.observer.OnExecutionFinish(
			ctx,
			metadata,
		)
	}

	return current, metadata, nil
}
