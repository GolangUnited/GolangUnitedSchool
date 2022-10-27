package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorUseCase struct {
	lg   *zap.Logger
	repo repository.CourseRepositoryInterface
}

func NewMentor(
	lg *zap.Logger,
	repo repository.MentorRepositoryInterface,
) *MentorUseCase {
	return &MentorUseCase{lg: lg, repo: repo}
}
