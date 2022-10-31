package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type ContactTypeRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewContactType(lg *zap.Logger, pool *pgxpool.Pool) *ContactTypeRepo {
	return &ContactTypeRepo{
		lg:   lg,
		pool: pool,
	}
}
