package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type StudentNoteTypeRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewStudentNoteType(
	lg *zap.Logger,
	pool *pgxpool.Pool) *StudentNoteTypeRepo {
	return &StudentNoteTypeRepo{
		lg:   lg,
		pool: pool,
	}
}
