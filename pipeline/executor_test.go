package pipeline

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutorRunsStagesInOrder(t *testing.T) {
	p := New[testItem](DefaultConfig())

	p.Map("double", func(ctx context.Context, e testItem) (testItem, error) {
		e.Value *= 2
		return e, nil
	})

	p.Map("increment", func(ctx context.Context, e testItem) (testItem, error) {
		e.Value++
		return e, nil
	})

	executor := p.Executor()

	out, metadata, err := executor.Execute(
		context.Background(),
		[]testItem{{Value: 5}},
	)

	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Equal(t, 11, out[0].Value)
	require.Len(t, metadata.Stages, 2)
}

func TestExecutorRunsPipeline(
	t *testing.T,
) {
	p := New[testItem](DefaultConfig())

	p.Map(
		"double",
		func(ctx context.Context, e testItem) (testItem, error) {
			e.Value *= 2

			return e, nil
		},
	)

	p.Map(
		"increment",
		func(ctx context.Context, e testItem) (testItem, error) {
			e.Value++

			return e, nil
		},
	)

	out, metadata, err := p.Executor().Execute(

		context.Background(),

		[]testItem{
			{Value: 5},
		},
	)

	require.NoError(t, err)

	require.Equal(
		t,
		11,
		out[0].Value,
	)

	require.Len(
		t,
		metadata.Stages,
		2,
	)
}

func TestMapStage(
	t *testing.T,
) {
	p := New[testItem](DefaultConfig())

	p.Map(
		"map",
		func(ctx context.Context, e testItem) (testItem, error) {
			e.Value += 100

			return e, nil
		},
	)

	out, _, err := p.Executor().Execute(

		context.Background(),

		[]testItem{
			{Value: 1},
		},
	)

	require.NoError(t, err)

	require.Equal(
		t,
		101,
		out[0].Value,
	)
}
