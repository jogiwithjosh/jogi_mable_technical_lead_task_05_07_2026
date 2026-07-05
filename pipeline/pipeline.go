package pipeline

import (
	"context"
	"errors"
	"time"
)

type Pipeline[T any] struct {
	config   Config
	stages   []Stage[T]
	runtime  *Runtime[T]
	observer Observer
}

func New[T any](
	cfg Config,
) *Pipeline[T] {
	return &Pipeline[T]{
		config: cfg,

		stages: make(
			[]Stage[T],
			0,
		),
	}
}

func (p *Pipeline[T]) Add(
	stage Stage[T],
) *Pipeline[T] {
	p.stages = append(
		p.stages,
		stage,
	)

	return p
}

func (p *Pipeline[T]) Execute(
	ctx context.Context,

	events []T,
) ([]T, ExecutionMetadata, error) {
	return execute(
		ctx,
		p,
		events,
	)
}

func execute[T any](
	ctx context.Context,
	p *Pipeline[T],
	items []T,
) ([]T, ExecutionMetadata, error) {
	meta := ExecutionMetadata{
		StartedAt: time.Now(),
	}

	current := items

	for _, stage := range p.stages {

		start := time.Now()

		result, err := stage.Process(
			ctx,
			current,
		)

		current = result.Items

		if err != nil {
			meta.FinishedAt = time.Now()
			return nil, meta, err
		}

		metrics := StageMetrics{
			Name:      stage.Name(),
			StartedAt: start,
			Input:     len(current),
		}

		metrics.FinishedAt = time.Now()
		metrics.Latency = metrics.FinishedAt.Sub(start)
		metrics.Output = len(result.Items)
		metrics.Dropped = result.Dropped
		metrics.Errors = 0

		metrics.Throughput = CalculateThroughput(
			len(result.Items),
			metrics.Latency,
		)
		meta.Stages = append(
			meta.Stages,
			metrics,
		)
		if p.observer != nil {
			p.observer.
				OnStageComplete(
					ctx,
					metrics,
				)
		}
		current = result.Items
	}

	meta.FinishedAt = time.Now()

	meta.Duration = meta.FinishedAt.Sub(
		meta.StartedAt,
	)
	if p.observer != nil {
		p.observer.
			OnExecutionFinish(
				ctx,
				meta,
			)
	}

	return current, meta, nil
}

func (p *Pipeline[T]) Start(
	ctx context.Context,
	sink Sink[T],
) error {
	if p.runtime != nil {
		return nil
	}

	p.runtime = NewRuntime(
		p.config,
		p.Executor(),
		sink,
	)
	p.runtime.Start(ctx)

	return nil
}

func (p *Pipeline[T]) Executor() *Executor[T] {
	return NewExecutor(
		p.stages,
		p.observer,
	)
}

func (p *Pipeline[T]) Stop() error {
	if p.runtime == nil {
		return nil
	}

	p.runtime.Stop()
	return nil
}

var ErrRuntimeNotStarted = errors.New("runtime not started")

func (p *Pipeline[T]) Publish(
	item T,
) error {
	if p.runtime == nil {
		return ErrRuntimeNotStarted
	}

	return p.runtime.Publish(
		item,
	)
}
