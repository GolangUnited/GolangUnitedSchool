package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type MentorRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewMentor(lg *zap.Logger, pool *pgxpool.Pool) *MentorRepo {
	return &MentorRepo{
		lg:   lg,
		pool: pool,
	}
}
