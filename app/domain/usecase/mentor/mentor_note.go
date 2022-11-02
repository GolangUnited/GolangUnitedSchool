package mentor

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorNoteUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewMentorNote(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *MentorNoteUseCase {
	return &MentorNoteUseCase{lg: lg, repo: repo}
}
