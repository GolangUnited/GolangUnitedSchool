package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorNoteUseCase struct {
	lg   *zap.Logger
	repo repository.CourseRepositoryInterface
}

func NewMentorNote(
	lg *zap.Logger,
	repo repository.MentorNoteRepositoryInterface,
) *MentorNoteUseCase {
	return &MentorNoteUseCase{lg: lg, repo: repo}
}
