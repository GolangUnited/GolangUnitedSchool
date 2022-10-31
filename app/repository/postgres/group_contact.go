package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type GroupContactRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewGroupContact(lg *zap.Logger, pool *pgxpool.Pool) *GroupContactRepo {
	return &GroupContactRepo{
		lg:   lg,
		pool: pool,
	}
}
