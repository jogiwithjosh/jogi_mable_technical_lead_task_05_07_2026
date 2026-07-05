package pipeline

import "context"

type Stage[T any] interface {
	Name() string

	Process(ctx context.Context, batch []T) (StageResult[T], error)
}

type StageResult[T any] struct {
	Items   []T
	Dropped int
}

type StageBase struct {
	name string
}

func (s StageBase) Name() string {
	return s.name
}

func NewStageBase(name string) StageBase {
	return StageBase{name: name}
}
