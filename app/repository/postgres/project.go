package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type ProjectRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewProject(lg *zap.Logger, pool *pgxpool.Pool) *ProjectRepo {
	return &ProjectRepo{
		lg:   lg,
		pool: pool,
	}
}
