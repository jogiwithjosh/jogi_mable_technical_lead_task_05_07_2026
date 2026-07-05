package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
)

type TestStruct struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Age        int                    `json:"age"`
	Salary     float64                ``
	Active     bool                   ``
	SessionID  string                 ``
	Page       string                 ``
	Properties map[string]interface{} ``
}

func createBenchmarkPipeline() *Pipeline[TestStruct] {
	p := New[TestStruct](DefaultConfig())

	p.Map(
		"validate",
		func(
			ctx context.Context,
			t TestStruct,
		) (TestStruct, error) {
			return t, nil
		},
	)

	p.Map(
		"normalize",
		func(
			ctx context.Context,
			t TestStruct,
		) (TestStruct, error) {
			t.Name = strings.ToLower(t.Name)

			return t, nil
		},
	)

	p.Filter(
		"filter",

		func(
			ctx context.Context,
			t TestStruct,
		) bool {
			return t.Active
		},
	)

	return p
}

func generateTestData(
	n int,
) []TestStruct {
	items := make(
		[]TestStruct,
		n,
	)

	for i := range items {
		items[i] = TestStruct{
			ID:     strconv.Itoa(i),
			Name:   "JOHN",
			Age:    25,
			Salary: 50000,
			Active: true,

			SessionID: uuid.NewString(),
			Page:      "/",
		}
	}

	return items
}

func benchmarkExecutor(
	b *testing.B,
	count int,
) {
	executor := createBenchmarkPipeline().
		Executor()

	items := generateTestData(count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := executor.Execute(
			context.Background(),
			items,
		)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPipeline10(
	b *testing.B,
) {
	benchmarkExecutor(
		b,
		10,
	)
}

func BenchmarkPipeline1K(b *testing.B) {
	benchmarkExecutor(b, 1000)
}

func BenchmarkPipeline100K(
	b *testing.B,
) {
	benchmarkExecutor(
		b,
		100000,
	)
}

func BenchmarkPipeline1M(
	b *testing.B,
) {
	benchmarkExecutor(
		b,
		1000000,
	)
}

func loadSampleEvent(
	b *testing.B,
) TestStruct {
	data, err := os.ReadFile(

		"testdata/sample_event.json",
	)
	if err != nil {
		b.Fatal(err)
	}

	var event TestStruct

	if err := json.Unmarshal(

		data,

		&event,
	); err != nil {
		b.Fatal(err)
	}

	return event
}

func generateEvents(
	n int,
	b *testing.B,
) []TestStruct {
	template := loadSampleEvent(b)

	items := make(
		[]TestStruct,
		n,
	)

	for i := range items {
		items[i] = template
		items[i].ID = uuid.NewString()
	}

	return items
}

func benchmarkRealEvent(
	b *testing.B,
	count int,
) {
	pipe := createRealPipeline()
	executor := pipe.Executor()
	events := generateEvents(
		count,
		b,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := executor.Execute(
			context.Background(),
			events,
		)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func createRealPipeline() *Pipeline[TestStruct] {
	p := New[TestStruct](
		DefaultConfig(),
	)

	p.
		Map(
			"validate",
			validateTestStruct,
		).
		Map(
			"normalize",
			normalizeTestStruct,
		).
		Map(
			"enrich",
			enrichTestStruct,
		).
		Filter(
			"filter",
			isValidTestStruct,
		).
		Generate(
			"generate",
			generateTestStruct,
		)

	return p
}

func validateTestStruct(
	ctx context.Context,
	t TestStruct,
) (TestStruct, error) {
	if t.ID == "" {
		return t, errors.New("missing id")
	}

	return t, nil
}

func normalizeTestStruct(
	ctx context.Context,
	t TestStruct,
) (TestStruct, error) {
	t.Name = strings.ToLower(t.Name)

	return t, nil
}

func enrichTestStruct(
	ctx context.Context,
	t TestStruct,
) (TestStruct, error) {
	if t.Properties == nil {
		t.Properties = map[string]interface{}{}
	}

	//t.Properties["pipeline"] = "benchmark"

	return t, nil
}

func isValidTestStruct(
	ctx context.Context,
	t TestStruct,
) bool {
	return t.Active
}

func generateTestStruct(
	ctx context.Context,
	t TestStruct,
) ([]TestStruct, error) {
	child := t

	child.ID += "_copy"

	child.Name += "_generated"

	return []TestStruct{
		child,
	}, nil
}
