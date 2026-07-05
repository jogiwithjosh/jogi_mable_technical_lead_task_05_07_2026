package pipeline

import (
	"context"
	"time"

	"pipeline/metrics"
)

type Batcher[T any] struct {
	size  int
	flush time.Duration
}

func NewBatcher[T any](
	size int,
	flush time.Duration,
) *Batcher[T] {
	return &Batcher[T]{
		size:  size,
		flush: flush,
	}
}

func (b *Batcher[T]) Run(
	ctx context.Context,
	input <-chan T,
	output chan<- []T,
) {
	timer := time.NewTicker(b.flush)
	defer timer.Stop()

	batch := make([]T, 0, b.size)

	flush := func() {
		if len(batch) == 0 {
			return
		}

		metrics.BatchSize.Observe(
			float64(len(batch)),
		)

		tmp := make([]T, len(batch))
		copy(tmp, batch)
		output <- tmp
		batch = batch[:0]
	}

	for {
		select {

		case <-ctx.Done():

			flush()
			return

		case item, ok := <-input:
			if !ok {
				flush()
				return
			}

			batch = append(batch, item)
			if len(batch) >= b.size {
				flush()
			}

		case <-timer.C:
			flush()
		}
	}
}
