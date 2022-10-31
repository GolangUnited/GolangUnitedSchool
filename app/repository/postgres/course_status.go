package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type CourseStatusRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewCourseStatus(lg *zap.Logger, pool *pgxpool.Pool) *CourseStatusRepo {
	return &CourseStatusRepo{
		lg:   lg,
		pool: pool,
	}
}
