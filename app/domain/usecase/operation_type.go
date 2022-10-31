package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationTypeUseCase struct {
	lg   *zap.Logger
	repo repository.OperationRepositoryInterface // CourseRepositoryInterface
}

func NewOperationType(lg *zap.Logger, repo repository.OperationTypeRepositoryInterface) *OperationTypeUseCase {
	return &OperationTypeUseCase{
		lg:   lg,
		repo: repo,
	}
}
