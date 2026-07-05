package pipeline

import "context"

type FilterFunc[T any] func(context.Context, T) bool

type FilterStage[T any] struct {
	StageBase
	fn FilterFunc[T]
}

func NewFilterStage[T any](
	name string,
	fn FilterFunc[T],
) *FilterStage[T] {
	return &FilterStage[T]{
		StageBase: NewStageBase(name),
		fn:        fn,
	}
}

func (f *FilterStage[T]) Process(
	ctx context.Context,
	batch []T,
) (StageResult[T], error) {
	out := make([]T, 0, len(batch))

	dropped := 0

	for _, item := range batch {
		if f.fn(ctx, item) {
			out = append(out, item)
		} else {
			dropped++
		}
	}

	return StageResult[T]{
		Items:   out,
		Dropped: dropped,
	}, nil
}

func (p *Pipeline[T]) Filter(
	name string,
	fn FilterFunc[T],
) *Pipeline[T] {
	p.Add(NewFilterStage(
		name,
		fn,
	))

	return p
}
