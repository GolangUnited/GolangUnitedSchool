package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

var (
	ErrorOperationTypeDoesntExists = "operation type doesn't exists"
)

func (r *PostgresRepository) GetOperationTypes(ctx context.Context) ([]model.OperationType, error) {

	rows, err := r.pool.Query(ctx,
		`SELECT id, title
			FROM operation_type`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get operation types")
	}
	defer rows.Close()

	var ops []model.OperationType
	for rows.Next() {
		var op model.OperationType
		err := rows.Scan(
			&op.Id,
			&op.Title,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan operation type")
		}

		ops = append(ops, op)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "couldn't get operation types")
	}

	return ops, nil
}

func (r *PostgresRepository) GetOperationTypeById(ctx context.Context, id int64) (*model.OperationType, error) {
	var operationType model.OperationType
	err := r.pool.QueryRow(ctx,
		`SELECT id, title 
			FROM operation_type
			WHERE id=$1`, id).
		Scan(&operationType.Id, &operationType.Title)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get opertion type by id %d", id)
	}

	return &operationType, nil
}

func (r *PostgresRepository) CreateOperationType(ctx context.Context, data *model.OperationType) error {
	var id int64
	err := r.pool.QueryRow(ctx,
		`INSERT INTO operation_type 
			(title)	
			VALUES ($1)	
			RETURNING id`,
		data.Title).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "couldn't create operation type")
	}

	return nil
}

func (r *PostgresRepository) UpdateOperationTypeById(ctx context.Context, id int64, data *model.OperationType) error {
	pct, err := r.pool.Exec(ctx,
		`UPDATE operation_type 
			SET title=$1 
			WHERE id=$2`,
		data.Title, id)
	if err != nil {
		return errors.Wrap(err, "couldn't update operation type")
	}
	if pct.RowsAffected() != 1 {
		return errors.Wrap(errors.New(ErrorOperationTypeDoesntExists),
			"couldn't update operation type")
	}

	return nil
}

func (r *PostgresRepository) DeleteOpertaionTypeById(ctx context.Context, id int64) error {
	pct, err := r.pool.Exec(ctx,
		`DELETE FROOM operation_type 
			WHERE id=$1`, id)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if pct.RowsAffected() != 1 {
		return errors.Wrap(errors.New(""), "")
	}

	return nil
}
