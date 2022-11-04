package operation

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type OperationUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewOperation(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *OperationUsecase {
	return &OperationUsecase{lg: lg, repo: repo}
}

func (u *OperationUsecase) GetOperations(
	ctx context.Context) ([]model.Operation, error) {
	panic("not implemented")
}
func (u *OperationUsecase) GetOperationById(
	ctx context.Context, id int64) (*model.Operation, error) {
	panic("not implemented")
}
func (u *OperationUsecase) AddOperation(
	ctx context.Context) error {
	panic("not implemented")
}
func (u *OperationUsecase) UpdateOperation(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
func (u *OperationUsecase) DeleteOperation(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
