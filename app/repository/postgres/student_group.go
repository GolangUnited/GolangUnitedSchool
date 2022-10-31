package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type StudentGroupRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewStudentGroup(lg *zap.Logger, pool *pgxpool.Pool) *StudentGroupRepo {
	return &StudentGroupRepo{
		lg:   lg,
		pool: pool,
	}
}
