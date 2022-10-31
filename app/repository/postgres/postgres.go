package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/pkg/errors"
)

type PostgresRepository struct {
	lg   logger.Logger
	pool *pgxpool.Pool
}

// NewPostgresRepository is impl RepositoryInterface for postgres
func NewPostgresRepository(lg logger.Logger, pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		lg:   lg,
		pool: pool,
	}
}

func NewDbPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: parse config")
	}

	dbConfig.MaxConns = 100
	dbConfig.MinConns = 10
	dbConfig.MaxConnLifetime = 30 * time.Minute
	dbConfig.MaxConnIdleTime = 15 * time.Minute
	dbConfig.HealthCheckPeriod = 20 * time.Second

	// dbConfig.ConnConfig.RuntimeParams["timezone"] = "Asia/Almaty"

	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: new pool with config")
	}

	return dbPool, nil
}
