package pipeline

import "context"

type MapFunc[T any] func(context.Context, T) (T, error)

type MapStage[T any] struct {
	StageBase
	fn MapFunc[T]
}

func NewMapStage[T any](
	name string,
	fn MapFunc[T],
) *MapStage[T] {
	return &MapStage[T]{
		StageBase: NewStageBase(name),
		fn:        fn,
	}
}

func (m *MapStage[T]) Process(
	ctx context.Context,
	batch []T,
) (StageResult[T], error) {
	out := make([]T, 0, len(batch))

	for _, item := range batch {

		result, err := m.fn(ctx, item)
		if err != nil {
			return StageResult[T]{}, err
		}

		out = append(out, result)
	}

	return StageResult[T]{
		Items: out,
	}, nil
}

func (p *Pipeline[T]) Map(
	name string,
	fn MapFunc[T],
) *Pipeline[T] {
	p.Add(NewMapStage(
		name,
		fn,
	))

	return p
}
