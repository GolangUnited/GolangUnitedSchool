package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetOperationTypes(
	ctx context.Context) ([]model.OperationType, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetOperationTypeById(
	ctx context.Context, id int64) (*model.OperationType, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddOperationType(
	ctx context.Context) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateOperationType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteOperationType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
