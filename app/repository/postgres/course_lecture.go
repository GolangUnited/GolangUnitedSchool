package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type CourseLectureRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewCourseLecture(lg *zap.Logger, pool *pgxpool.Pool) *CourseLectureRepo {
	return &CourseLectureRepo{
		lg:   lg,
		pool: pool,
	}
}
