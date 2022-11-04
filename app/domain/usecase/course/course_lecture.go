package course

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseLectureseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewCourseLecture(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *CourseLectureseCase {
	return &CourseLectureseCase{lg: lg, repo: repo}
}
