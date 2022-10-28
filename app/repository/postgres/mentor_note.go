package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type MentorNoteRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewMentorNote(lg *zap.Logger, pool *pgxpool.Pool) *MentorNoteRepo {
	return &MentorNoteRepo{
		lg:   lg,
		pool: pool,
	}
}
