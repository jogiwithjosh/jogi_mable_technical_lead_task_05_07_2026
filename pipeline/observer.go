package pipeline

import (
	"context"
	"time"
)

type Observer interface {
	OnExecutionStart(
		context.Context,
	)

	OnExecutionFinish(
		context.Context,
		ExecutionMetadata,
	)

	OnStageComplete(
		context.Context,
		StageMetrics,
	)
}

type StageMetrics struct {
	ExecutionID string
	Name        string
	StartedAt   time.Time
	FinishedAt  time.Time
	Latency     time.Duration
	Input       int
	Output      int
	Dropped     int
	Errors      int
	Throughput  float64
}

type ExecutionMetadata struct {
	ID         string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
	Stages     []StageMetrics
}

func CalculateThroughput(
	items int,
	duration time.Duration,
) float64 {
	if duration <= 0 {
		return 0
	}

	return float64(items) /
		duration.Seconds()
}
