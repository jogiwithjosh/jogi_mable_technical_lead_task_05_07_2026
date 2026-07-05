package pipeline

import (
	"context"
	"time"

	"pipeline/metrics"
)

type Worker[T any] struct {
	id       int
	executor *Executor[T]
	sink     Sink[T]
}

func NewWorker[T any](
	id int,
	executor *Executor[T],
	sink Sink[T],
) *Worker[T] {
	return &Worker[T]{
		id:       id,
		executor: executor,
		sink:     sink,
	}
}

func (w *Worker[T]) Run(
	ctx context.Context,
	batches <-chan []T,
) {
	for {
		select {

		case <-ctx.Done():

			return

		case batch, ok := <-batches:

			if !ok || len(batch) <= 0 {
				return
			}
			start := time.Now()
			items, _, err := w.executor.Execute(
				ctx,
				batch,
			)

			metrics.PipelineLatency.Observe(
				time.Since(start).Seconds(),
			)

			if err != nil {
				metrics.EventsFailed.Add(
					float64(len(batch)),
				)
				continue
			}
			metrics.EventsProcessed.Add(
				float64(len(items)),
			)

			if w.sink != nil {
				_ = w.sink.Drain(
					ctx,
					items,
				)
			}
		}
	}
}
