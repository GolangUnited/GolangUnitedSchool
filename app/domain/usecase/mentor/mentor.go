package mentor

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewMentor(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *MentorUseCase {
	return &MentorUseCase{lg: lg, repo: repo}
}
