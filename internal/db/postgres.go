package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// IPgxPool use this as dependency instead of *pgxpool.Pool so we can mock for
// unit test. Add method when needed.
type IPgxPool interface {
	Ping(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

// Postgres -.
type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    IPgxPool // use IPgxPool instead *pgxpool.Pool
}

// NewPostgresPoolConnection return postgres pool connection.
func NewPostgresPoolConnection(cfg config.Config) (*Postgres, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.PG.Username, cfg.PG.Password, cfg.PG.Host, cfg.PG.Port, cfg.PG.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}
	poolConfig.MaxConns = int32(cfg.PG.PoolMax)

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
