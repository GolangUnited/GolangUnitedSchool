package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetLogOperationById(
	ctx context.Context, id int64) (*model.LogOperation, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddLogOperation(
	ctx context.Context, data *model.LogOperation) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteLogOperation(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
