package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/avast/retry-go/v4"
	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"
)

const maxDelay = 10 * time.Second

type DB struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
}

func New(ctx context.Context, dsn string, attempts int, timeout time.Duration) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to parse config: %w", err)
	}

	cfg.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   zapadapter.NewLogger(zap.L()),
		LogLevel: tracelog.LogLevelTrace,
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: can't create connection pool: %w", err)
	}

	err = retry.Do(func() error {
		if err = pool.Ping(ctx); err != nil {
			return fmt.Errorf("postgres: failed to ping database: %w", err)
		}
		return nil
	},
		retry.Context(ctx),
		retry.Attempts(uint(attempts)),
		retry.MaxDelay(maxDelay),
		retry.Delay(timeout),
		retry.DelayType(retry.BackOffDelay),
		retry.OnRetry(func(attempt uint, err error) {
			zap.L().Info("trying to connect postgres...",
				zap.Uint("attempt", attempt), zap.Error(err))
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to connect to database: %w", err)
	}

	return &DB{
		Pool:    pool,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar), // PostgreSQL placeholder format
	}, err
}
