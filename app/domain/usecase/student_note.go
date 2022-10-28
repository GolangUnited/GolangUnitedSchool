package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentNoteUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewStudentNote(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *StudentNoteUseCase {
	return &StudentNoteUseCase{lg: lg, repo: repo}
}
