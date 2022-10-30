package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type repository struct {
	lg   logger.Logger
	pool *pgxpool.Pool
}

func NewRepository(lg logger.Logger, pool *pgxpool.Pool) *repository {
	return &repository{
		lg:   lg,
		pool: pool,
	}
}
