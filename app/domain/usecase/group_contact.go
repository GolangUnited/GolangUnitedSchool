package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type GroupContactUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewGroupContact(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *GroupContactUseCase {
	return &GroupContactUseCase{lg: lg, repo: repo}
}
