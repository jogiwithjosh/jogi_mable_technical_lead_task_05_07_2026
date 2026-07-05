package pipeline

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataCollected(t *testing.T) {
	p := New[testItem](DefaultConfig())

	p.Map(
		"stage1",
		func(ctx context.Context, e testItem) (testItem, error) {
			return e, nil
		},
	)

	_, metadata, err := p.Executor().Execute(
		context.Background(),
		[]testItem{{1, 1}},
	)

	require.NoError(t, err)
	require.Len(t, metadata.Stages, 1)
	require.NotZero(
		t,
		metadata.Stages[0].Latency,
	)
}

func TestRuntimeConcurrentPublish(
	t *testing.T,
) {
	cfg := DefaultConfig()

	cfg.Workers = 8

	runtime, sink := createTestRuntime(cfg)

	ctx := context.Background()

	runtime.Start(ctx)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func(worker int) {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				_ = runtime.Publish(

					TestStruct{
						Age: worker*100 + j,
					},
				)
			}
		}(i)
	}

	wg.Wait()

	runtime.Stop()

	require.Equal(
		t,
		10000,
		sink.Count(),
	)
}

func TestRuntimeProcessesEvents(
	t *testing.T,
) {
	cfg := DefaultConfig()

	cfg.Workers = 4

	runtime, sink := createTestRuntime(cfg)

	ctx := context.Background()

	runtime.Start(ctx)

	for i := 0; i < 1000; i++ {
		require.NoError(
			t,
			runtime.Publish(
				TestStruct{
					Age: i,
				},
			),
		)
	}

	runtime.Stop()

	require.Equal(
		t,
		1000,
		sink.Count(),
	)
}

func TestRuntimePublish(t *testing.T) {
	rt, sink := createTestRuntime(DefaultConfig())

	ctx := context.Background()

	rt.Start(ctx)

	item := TestStruct{
		ID: "1",

		Name: "John",

		Active: true,
	}

	require.NoError(
		t,
		rt.Publish(item),
	)

	time.Sleep(200 * time.Millisecond)

	rt.Stop()

	assert.Equal(
		t,
		sink.Count(),
		1,
	)
}

func TestRuntimeBatchProcessing(t *testing.T) {
	rt, sink := createTestRuntime(DefaultConfig())

	ctx := context.Background()

	rt.Start(ctx)

	for i := 0; i < 100; i++ {
		err := rt.Publish(
			TestStruct{
				ID:     strconv.Itoa(i),
				Active: true,
			},
		)

		require.NoError(t, err)
	}

	time.Sleep(500 * time.Millisecond)
	rt.Stop()

	assert.Equal(
		t,
		sink.Count(),
		100,
	)
}

func TestRuntimeConcurrentPublishTestStruct(t *testing.T) {
	rt, sink := createTestRuntime(DefaultConfig())

	ctx := context.Background()

	rt.Start(ctx)

	var wg sync.WaitGroup

	producers := 20

	perProducer := 500

	for p := 0; p < producers; p++ {

		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for i := 0; i < perProducer; i++ {

				err := rt.Publish(

					TestStruct{
						ID: strconv.Itoa(
							id*perProducer + i,
						),

						Active: true,
					},
				)

				require.NoError(
					t,
					err,
				)
			}
		}(p)
	}

	wg.Wait()

	rt.Stop()

	assert.Equal(
		t,
		sink.Count(),
		producers*perProducer,
	)
}

func TestRuntimeEmpty(t *testing.T) {
	rt, sink := createTestRuntime(DefaultConfig())

	rt.Start(context.Background())

	rt.Stop()

	assert.Empty(

		t,

		sink.Count(),
	)
}

func TestRuntimeGracefulShutdown(t *testing.T) {
	rt, sink := createTestRuntime(DefaultConfig())

	rt.Start(context.Background())

	for i := 0; i < 500; i++ {

		err := rt.Publish(

			TestStruct{
				ID: strconv.Itoa(i),

				Active: true,
			},
		)

		require.NoError(
			t,
			err,
		)
	}

	rt.Stop()

	assert.Equal(
		t,
		sink.Count(),
		500,
	)
}

func TestRuntimeFilter(t *testing.T) {
	sink := &mockSink[TestStruct]{}

	pipe := New[TestStruct](DefaultConfig()).
		Filter(
			"active",
			func(
				ctx context.Context,
				t TestStruct,
			) bool {
				return t.Active
			},
		)

	rt := NewRuntime(
		DefaultConfig(),
		pipe.Executor(),
		sink,
	)

	rt.Start(context.Background())

	_ = rt.Publish(
		TestStruct{
			ID:     "1",
			Active: true,
		},
	)

	_ = rt.Publish(
		TestStruct{
			ID:     "2",
			Active: false,
		},
	)

	time.Sleep(200 * time.Millisecond)
	rt.Stop()
	assert.Equal(
		t,
		sink.Count(),
		1,
	)
}

func BenchmarkRuntime1M_4W500B(b *testing.B) {
	runtime, _ := createTestRuntime(DefaultConfig())

	runtime.Start(context.Background())

	events := generateEvents(1000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}

func BenchmarkRuntime10M_4W500B(b *testing.B) {
	runtime, _ := createTestRuntime(DefaultConfig())

	runtime.Start(context.Background())

	events := generateEvents(10000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}

func BenchmarkRuntime1M_1W128B(b *testing.B) {
	runtime, _ := createTestRuntime(Config{
		Workers:    1,
		BatchSize:  128,
		BufferSize: 10000,
		FlushEvery: time.Second,
	})

	runtime.Start(context.Background())

	events := generateEvents(1000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}

func BenchmarkRuntime1M_8W500B(b *testing.B) {
	runtime, _ := createTestRuntime(Config{
		Workers:    8,
		BatchSize:  500,
		BufferSize: 10000,
		FlushEvery: time.Second,
	})

	runtime.Start(context.Background())

	events := generateEvents(1000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}

func BenchmarkRuntime1M_8W1024B(b *testing.B) {
	runtime, _ := createTestRuntime(Config{
		Workers:    8,
		BatchSize:  1024,
		BufferSize: 10000,
		FlushEvery: time.Second,
	})

	runtime.Start(context.Background())

	events := generateEvents(1000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}

func BenchmarkRuntime1M_16W500B(b *testing.B) {
	runtime, _ := createTestRuntime(Config{
		Workers:    16,
		BatchSize:  500,
		BufferSize: 10000,
		FlushEvery: time.Second,
	})

	runtime.Start(context.Background())

	events := generateEvents(1000000, b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range events {
			_ = runtime.Publish(TestStruct{Age: e.Age})
		}
	}

	runtime.Stop()
}
