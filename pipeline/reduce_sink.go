package pipeline

import "context"

type Reducer[T any, R any] func(
	context.Context,
	R,
	T,
) R

func Reduce[T any, R any](
	ctx context.Context,
	items []T,
	initial R,
	fn Reducer[T, R],
) R {
	result := initial

	for _, item := range items {
		result = fn(
			ctx,
			result,
			item,
		)
	}

	return result
}

type Sink[T any] interface {
	Drain(
		context.Context,
		[]T,
	) error
}

func Drain[T any](
	ctx context.Context,
	items []T,
	sink Sink[T],
) error {
	if sink == nil {
		return nil
	}

	return sink.Drain(
		ctx,
		items,
	)
}
