package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetOperationLogById(ctx context.Context, id int64) (*model.OperationLog, error) {
	panic("empty")
}

func (r *PostgresRepository) CreateOperationLog(ctx context.Context) (*model.OperationLog, error) {
	panic("empty")
}

func (r *PostgresRepository) UpdateOprationLog(ct context.Context) {

}

func (r *PostgresRepository) DeleteOperationLog(ctx context.Context) error {
	panic("empty")
}
