package health

import (
	"context"
	"time"

	"api/internal/database"
)

type Service struct {
	db *database.ClickHouse
}

func New(
	db *database.ClickHouse,
) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Ready() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	return s.db.DB.PingContext(ctx)
}

func (s *Service) Live() bool {
	return true
}
