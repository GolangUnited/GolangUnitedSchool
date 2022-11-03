package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetOperations(
	ctx context.Context) (*model.Operation, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetOperationById(
	ctx context.Context, id int64) (*model.Operation, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddOperation(
	ctx context.Context) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateOperation(
	ctx context.Context, id int64) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteOperation(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
