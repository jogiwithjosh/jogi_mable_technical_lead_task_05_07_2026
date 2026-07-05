package pipeline

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBatcherFlushBySize(t *testing.T) {
	batcher := NewBatcher[int](
		3,
		time.Minute,
	)

	in := make(chan int)
	out := make(chan []int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go batcher.Run(ctx, in, out)

	in <- 1
	in <- 2
	in <- 3

	batch := <-out

	require.Len(t, batch, 3)
}
