package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type OperationTypeRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewOperationType(lg *zap.Logger, pool *pgxpool.Pool) *OperationTypeRepo {
	return &OperationTypeRepo{
		lg:   lg,
		pool: pool,
	}
}
