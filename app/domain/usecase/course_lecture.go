package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseLectureseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewCourseLecture(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseLectureseCase {
	return &CourseLectureseCase{lg: lg, repo: repo}
}
