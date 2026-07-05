package pipeline

import (
	"errors"

	"pipeline/metrics"
)

var ErrQueueClosed = errors.New("pipeline queue closed")

type Queue[T any] struct {
	ch chan T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		ch: make(chan T, size),
	}
}

func (q *Queue[T]) Publish(item T) {
	q.ch <- item
}

func (q *Queue[T]) Close() {
	metrics.QueueDepth.Set(0)
	close(q.ch)
}

func (q *Queue[T]) Subscribe() <-chan T {
	return q.ch
}

func (q *Queue[T]) Len() int {
	return len(q.ch)
}
