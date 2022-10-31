package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationUseCase struct {
	lg   *zap.Logger
	repo repository.OperationRepositoryInterface // CourseRepositoryInterface
}

func NewOperation(lg *zap.Logger, repo repository.OperationRepositoryInterface) *OperationUseCase {
	return &OperationUseCase{
		lg:   lg,
		repo: repo,
	}
}
