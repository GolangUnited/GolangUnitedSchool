package operation

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationTypeUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewOperationType(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *OperationTypeUsecase {
	return &OperationTypeUsecase{lg: lg, repo: repo}
}

func (u *OperationTypeUsecase) GetOperationTypes(
	ctx context.Context) (*model.OperationType, error) {
	panic("not implemented")
}
func (u *OperationTypeUsecase) GetOperationTypeById(
	ctx context.Context, id int64) (*model.OperationType, error) {
	panic("not implemented")
}
func (u *OperationTypeUsecase) AddOperationType(
	ctx context.Context) error {
	panic("not implemented")
}
func (u *OperationTypeUsecase) UpdateOperationType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
func (u *OperationTypeUsecase) DeleteOperationType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
