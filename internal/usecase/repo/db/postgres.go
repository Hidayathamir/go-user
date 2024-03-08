// Package db contains func related to database.
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Postgres -.
type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

// NewPostgresPoolConnection return postgres pool connection.
func NewPostgresPoolConnection() (*Postgres, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.PG.Username, config.PG.Password, config.PG.Host, config.PG.Port, config.PG.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}
	poolConfig.MaxConns = int32(config.PG.PoolMax)

	pg := &Postgres{
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	for i := 0; i < 10; i++ {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			logrus.
				WithField("attempt count", i+1).
				Warnf("err create new conn pool: pgxpool.NewWithConfig: %v", err)

			time.Sleep(time.Second)

			continue
		}

		err = pg.Pool.Ping(context.Background())
		if err != nil {
			logrus.
				WithField("attempt count", i+1).
				Warnf("Postgres.Pool.Ping: %v", err)

			time.Sleep(time.Second)

			continue
		}

		break
	}

	if err != nil {
		return nil, fmt.Errorf("err 10 times when try to create conn pool: %w", err)
	}

	logrus.Info("success create db connection pool 🟢")

	return pg, nil
}
