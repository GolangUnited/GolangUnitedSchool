package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationLogUseCase struct {
	lg   *zap.Logger
	repo repository.OperationLogRepositoryInterface
}

func NewOperationLog(lg *zap.Logger, repo repository.OperationLogRepositoryInterface) *OperationLogUseCase {
	return &OperationLogUseCase{
		lg:   lg,
		repo: repo,
	}
}
