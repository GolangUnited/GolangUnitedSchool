package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type CourseRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewCourse(lg *zap.Logger, pool *pgxpool.Pool) *CourseRepo {
	return &CourseRepo{
		lg:   lg,
		pool: pool,
	}
}
