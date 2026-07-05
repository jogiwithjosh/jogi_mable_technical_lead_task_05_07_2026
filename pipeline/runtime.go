package pipeline

import (
	"context"
	"sync"

	"pipeline/metrics"
)

type Runtime[T any] struct {
	cfg      Config
	queue    *Queue[T]
	batcher  *Batcher[T]
	executor *Executor[T]
	sink     Sink[T]
	batches  chan []T
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func NewRuntime[T any](
	cfg Config,
	executor *Executor[T],
	sink Sink[T],
) *Runtime[T] {
	return &Runtime[T]{
		cfg:   cfg,
		queue: NewQueue[T](cfg.BufferSize),
		batcher: NewBatcher[T](
			cfg.BatchSize,
			cfg.FlushEvery,
		),
		executor: executor,
		sink:     sink,
		batches: make(
			chan []T,
			cfg.Workers,
		),
	}
}

func (r *Runtime[T]) Start(
	parent context.Context,
) {
	ctx, cancel := context.WithCancel(parent)
	r.cancel = cancel
	r.wg.Add(1)

	go func() {
		defer r.wg.Done()

		r.batcher.Run(
			ctx,
			r.queue.Subscribe(),
			r.batches,
		)

		close(r.batches)
	}()

	for i := 0; i < r.cfg.Workers; i++ {

		r.wg.Add(1)

		go func(id int) {
			defer r.wg.Done()

			NewWorker(
				id,
				r.executor,
				r.sink,
			).Run(
				ctx,
				r.batches,
			)
		}(i + 1)
	}
}

func (r *Runtime[T]) Publish(
	item T,
) error {
	metrics.QueueDepth.Set(
		float64(r.queue.Len()),
	)
	r.queue.Publish(item)
	return nil
}

func (r *Runtime[T]) Stop() {
	// Stop accepting new events
	r.queue.Close()

	// Wait for batcher to drain queue
	r.wg.Wait()

	// Now cancel background context
	if r.cancel != nil {
		r.cancel()
	}
}
