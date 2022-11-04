package group

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type GroupContactUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewGroupContact(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *GroupContactUseCase {
	return &GroupContactUseCase{lg: lg, repo: repo}
}
