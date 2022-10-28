package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewMentor(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *MentorUseCase {
	return &MentorUseCase{lg: lg, repo: repo}
}
