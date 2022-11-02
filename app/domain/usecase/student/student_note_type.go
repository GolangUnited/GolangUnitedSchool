package student

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentNoteTypeUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewStudentNoteType(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *StudentNoteTypeUseCase {
	return &StudentNoteTypeUseCase{lg: lg, repo: repo}
}
