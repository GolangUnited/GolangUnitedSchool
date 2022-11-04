package student

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentNoteUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentNote(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentNoteUseCase {
	return &StudentNoteUseCase{lg: lg, repo: repo}
}
