package pipeline

import (
	"context"
	"sync"
)

type testItem struct {
	ID    int
	Value int
}

type mockSink[T any] struct {
	mu    sync.Mutex
	items []T
	count int
}

func (m *mockSink[T]) Drain(
	ctx context.Context,
	items []T,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	//m.items = append(m.items, items...)
	m.count += len(items)

	return nil
}

func (m *mockSink[T]) Count() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	//return len(m.items)
	return m.count
}

func (m *mockSink[T]) Items() []T {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]T, len(m.items))
	copy(out, m.items)

	return out
}

func createTestPipeline() *Pipeline[testItem] {
	p := New[testItem](DefaultConfig())

	p.Map(
		"identity",
		func(ctx context.Context, t testItem) (testItem, error) {
			return t, nil
		},
	)

	return p
}

func createTestRuntime(cfg Config) (*Runtime[TestStruct], *mockSink[TestStruct]) {
	sink := &mockSink[TestStruct]{}

	pipe := New[TestStruct](DefaultConfig()).
		Map(
			"normalize",
			func(
				ctx context.Context,
				t TestStruct,
			) (TestStruct, error) {
				return t, nil
			},
		)

	rt := NewRuntime(
		cfg,
		pipe.Executor(),
		sink,
	)

	return rt, sink
}
