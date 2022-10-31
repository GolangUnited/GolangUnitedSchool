package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type ContactTypeUseCase struct {
	lg   *zap.Logger
	repo repository.ContactTypeRepositoryInterface // CourseRepositoryInterface
}

func NewContactType(lg *zap.Logger, repo repository.ContactTypeRepositoryInterface) *ContactTypeUseCase {
	return &ContactTypeUseCase{
		lg:   lg,
		repo: repo,
	}
}
