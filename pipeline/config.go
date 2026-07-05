package pipeline

import "time"

type Config struct {
	Workers    int
	BatchSize  int
	BufferSize int
	FlushEvery time.Duration
}

func DefaultConfig() Config {
	return Config{
		Workers:    4,
		BatchSize:  500,
		BufferSize: 10000,
		FlushEvery: time.Second,
	}
}
