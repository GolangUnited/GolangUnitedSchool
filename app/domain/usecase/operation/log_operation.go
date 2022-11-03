package operation

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type LogOperationUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewLogOperation(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *LogOperationUsecase {
	return &LogOperationUsecase{lg: lg, repo: repo}
}

func (u *LogOperationUsecase) GetLogOperationById(
	ctx context.Context,
	id int64) (*model.LogOperation, error) {
	panic("not implemented")
}

func (u *LogOperationUsecase) AddLogOperation(
	ctx context.Context,
	data *model.LogOperation) error {
	panic("not implemented")
}

func (u *LogOperationUsecase) DeleteLogOperation(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
