package pipeline

import "context"

type GenerateFunc[T any] func(context.Context, T) ([]T, error)

type GenerateStage[T any] struct {
	StageBase
	fn GenerateFunc[T]
}

func NewGenerateStage[T any](
	name string,
	fn GenerateFunc[T],
) *GenerateStage[T] {
	return &GenerateStage[T]{
		StageBase: NewStageBase(name),
		fn:        fn,
	}
}

func (g *GenerateStage[T]) Process(
	ctx context.Context,
	batch []T,
) (StageResult[T], error) {
	out := make([]T, 0, len(batch))

	for _, item := range batch {

		out = append(out, item)

		items, err := g.fn(ctx, item)
		if err != nil {
			return StageResult[T]{}, err
		}

		out = append(out, items...)
	}

	return StageResult[T]{
		Items: out,
	}, nil
}

func (p *Pipeline[T]) Generate(
	name string,
	fn GenerateFunc[T],
) *Pipeline[T] {
	p.Add(NewGenerateStage(
		name,
		fn,
	))

	return p
}
