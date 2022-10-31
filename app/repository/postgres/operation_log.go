package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type OperationLogRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewOperationLog(lg *zap.Logger, pool *pgxpool.Pool) *OperationLogRepo {
	return &OperationLogRepo{
		lg:   lg,
		pool: pool,
	}
}
