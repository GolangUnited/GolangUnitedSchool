package person

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type PersonUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewPerson(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *PersonUseCase {
	return &PersonUseCase{lg: lg, repo: repo}
}
