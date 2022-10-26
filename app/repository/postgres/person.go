package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type PersonRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewPerson(lg *zap.Logger, pool *pgxpool.Pool) *PersonRepo {
	return &PersonRepo{
		lg:   lg,
		pool: pool,
	}
}
