package database

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"

	"api/internal/config"
)

type ClickHouse struct {
	DB *sql.DB
}

func New(cfg config.ClickHouseConfig) (*ClickHouse, error) {
	options := &clickhouse.Options{

		Addr: []string{
			fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		},

		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},

		DialTimeout: time.Second * 5,

		MaxOpenConns: 20,
		MaxIdleConns: 10,

		ConnMaxLifetime: time.Hour,
	}

	if cfg.Secure {
		// options.Protocol = clickhouse.Native
		options.TLS = &tls.Config{}
	}

	conn := clickhouse.OpenDB(options)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	fmt.Printf("%+v", options)

	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return &ClickHouse{
		DB: conn,
	}, nil
}

func (c *ClickHouse) Close() error {
	return c.DB.Close()
}
