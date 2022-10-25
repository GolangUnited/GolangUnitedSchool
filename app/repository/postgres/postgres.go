package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgreRepository {
	return &PostgreRepository{pool: pool}
}
