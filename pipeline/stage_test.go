package pipeline

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterStage(t *testing.T) {
	p := New[testItem](DefaultConfig())

	p.Filter("even", func(ctx context.Context, e testItem) bool {
		return e.Value%2 == 0
	})

	executor := p.Executor()

	out, _, err := executor.Execute(
		context.Background(),
		[]testItem{
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
		},
	)

	require.NoError(t, err)
	require.Len(t, out, 2)
	require.Equal(t, 2, out[0].Value)
	require.Equal(t, 4, out[1].Value)
}

func TestGenerateStage(t *testing.T) {
	p := New[testItem](DefaultConfig())

	p.Generate(
		"duplicate",
		func(ctx context.Context, e testItem) ([]testItem, error) {
			return []testItem{
				{Value: e.Value + 100},
			}, nil
		},
	)

	executor := p.Executor()

	out, _, err := executor.Execute(
		context.Background(),
		[]testItem{{1, 1}},
	)

	require.NoError(t, err)
	require.Len(t, out, 2)
}

func TestIfStage(t *testing.T) {
	left := New[testItem](DefaultConfig())

	left.Map("left", func(ctx context.Context, e testItem) (testItem, error) {
		e.Value *= 2
		return e, nil
	})

	right := New[testItem](DefaultConfig())

	right.Map("right", func(ctx context.Context, e testItem) (testItem, error) {
		e.Value *= 3
		return e, nil
	})

	main := New[testItem](DefaultConfig())

	main.If(
		"branch",
		func(ctx context.Context, e testItem) bool {
			return e.Value < 10
		},
		left,
		right,
	)

	out, _, err := main.Executor().Execute(
		context.Background(),
		[]testItem{
			{5, 20},
			{25, 20},
		},
	)

	require.NoError(t, err)
	require.Len(t, out, 2)
}
