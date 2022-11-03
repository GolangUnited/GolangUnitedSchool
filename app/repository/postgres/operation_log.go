package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetOperationLogById(
	ctx context.Context, id int64) (*model.OperationLog, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddOperationLog(
	ctx context.Context, data *model.OperationLog) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteOperationLog(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
