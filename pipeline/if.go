package pipeline

import "context"

type Predicate[T any] func(context.Context, T) bool

type IfStage[T any] struct {
	StageBase

	predicate Predicate[T]

	trueExecutor  *Executor[T]
	falseExecutor *Executor[T]
}

func NewIfStage[T any](
	name string,
	predicate Predicate[T],
	trueExecutor *Executor[T],
	falseExecutor *Executor[T],
) *IfStage[T] {
	return &IfStage[T]{
		StageBase:     NewStageBase(name),
		predicate:     predicate,
		trueExecutor:  trueExecutor,
		falseExecutor: falseExecutor,
	}
}

func (i *IfStage[T]) Process(
	ctx context.Context,
	batch []T,
) (StageResult[T], error) {
	left := make([]T, 0)
	right := make([]T, 0)

	for _, item := range batch {
		if i.predicate(ctx, item) {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	out := make([]T, 0, len(batch))

	if len(left) > 0 {

		items, _, err := i.trueExecutor.Execute(ctx, left)
		if err != nil {
			return StageResult[T]{}, err
		}

		out = append(out, items...)
	}

	if len(right) > 0 {

		items, _, err := i.falseExecutor.Execute(ctx, right)
		if err != nil {
			return StageResult[T]{}, err
		}

		out = append(out, items...)
	}

	return StageResult[T]{
		Items: out,
	}, nil
}

func (p *Pipeline[T]) If(
	name string,
	predicate Predicate[T],
	truePipeline *Pipeline[T],
	falsePipeline *Pipeline[T],
) *Pipeline[T] {
	p.Add(NewIfStage(
		name,
		predicate,
		truePipeline.Executor(),
		falsePipeline.Executor(),
	))

	return p
}
