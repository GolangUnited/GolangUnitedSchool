package operation

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationLogUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewOperationLog(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *OperationLogUsecase {
	return &OperationLogUsecase{lg: lg, repo: repo}
}

func (u *OperationLogUsecase) GetOperationLogById(
	ctx context.Context,
	id int64) (*model.OperationLog, error) {
	panic("not implemented")
}

func (u *OperationLogUsecase) AddOperationLog(
	ctx context.Context,
	data *model.OperationLog) error {
	panic("not implemented")
}

func (u *OperationLogUsecase) DeleteOperationLog(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
