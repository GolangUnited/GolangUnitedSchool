package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type OperationRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewOperation(lg *zap.Logger, pool *pgxpool.Pool) *OperationRepo {
	return &OperationRepo{
		lg:   lg,
		pool: pool,
	}
}
