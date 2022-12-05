package postgres

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

var (
	ErrorOperationLogDoesntExists = "operation log doesn't exists"
)

func (r *PostgresRepository) GetOperationLogById(ctx context.Context, id int64) (*model.OperationLog, error) {
	panic("empty")
}

func (r *PostgresRepository) CreateOperationLog(ctx context.Context) (*model.OperationLog, error) {
	panic("empty")
}

func (r *PostgresRepository) UpdateOprationLogById(ct context.Context, id int64, data *model.OperationLog) {

}

func (r *PostgresRepository) DeleteOperationLog(ctx context.Context, id int64) error {
	pct, err := r.pool.Exec(ctx,
		`DELETE FROM operation_log 
			WHERE id=$1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete operation log by id %d", id)
	}
	if pct.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorOperationLogDoesntExists),
			"couldn't delete operation log by id %d", id)
	}

	return nil
}
