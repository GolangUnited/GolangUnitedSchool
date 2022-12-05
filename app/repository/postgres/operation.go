package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

var (
	ErrorOperationDoesntExists = "operation doesn't exists"
)

func (r *PostgresRepository) GetOperationById(ctx context.Context, id int64) (*model.Operation, error) {
	var op model.Operation
	err := r.pool.QueryRow(ctx,
		`SELECT 
				op.id, 
				op.title, 
				op.is_logging,
				opt.id,
				opt.title
		FROM operation AS op
			INNER JOIN operation_type AS opt 
				ON op.id = opt.id
			WHERE id=$1`, id).
		Scan(
			&op.Id,
			&op.Title,
			&op.IsLogging,
			&op.OperationType.Id,
			&op.OperationType.Title,
		)
	if err != nil {
		return nil, errors.Wrapf(err,
			"couldn't get operation by id %d", id)
	}

	return &op, nil
}

func (r *PostgresRepository) GetOperations(ctx context.Context) (*model.OperationList, error) {
	var opsList model.OperationList

	return &opsList, nil
}

func (r *PostgresRepository) CreateOperation(ctx context.Context, op *model.OperationCreate) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// Title
	args = append(args, op.Title)
	keys = append(keys, "title")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// operation type id
	if op.OperationTypeId != nil {
		args = append(args, op.OperationTypeId)
		keys = append(keys, "operation_type_id")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	// is logging
	if op.IsLogging != nil {
		args = append(args, op.IsLogging)
		keys = append(keys, "is_logging")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	query := fmt.Sprintf(`INSERT INTO operation (%s)
				VALUES (%s)
				RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "))

	var id int64
	err := r.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "couldn't create operation")
	}

	return id, nil
}

func (r *PostgresRepository) UpdateOperationById(ctx context.Context, id int64, op *model.OperationUpdate) error {
	var args []interface{}
	var keys []string

	// operation type id
	if op.OperationTypeId != nil {
		args = append(args, op.OperationTypeId)
		keys = append(keys, fmt.Sprintf("operation_type_id = $%d", len(args)))
	}

	// title
	if op.Title != nil {
		args = append(args, op.Title)
		keys = append(keys, fmt.Sprintf("title = $%d", len(args)))
	}

	// is logging
	if op.IsLogging != nil {
		args = append(args, op.IsLogging)
		keys = append(keys, fmt.Sprintf("is_logging = $%d", len(args)))
	}

	// id
	args = append(args, id)

	query := fmt.Sprintf(`UPDATE operation 
				SET %s 
				WHERE id=$%d`,
		strings.Join(keys, ", "), len(args))
	r.lg.Debug(query)

	pct, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update operation by id %d", id)
	}
	if pct.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorOperationDoesntExists),
			"couldn't update operation by id %d", id)
	}

	return nil
}

func (r *PostgresRepository) DeleteOperation(ctx context.Context, id int64) error {
	pct, err := r.pool.Exec(ctx,
		`DELETE FROM operation
			WHERE id=$1`, id)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't delete operation by id %d", id)
	}
	if pct.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorOperationDoesntExists),
			"couldn't delete operation by id %d", id)
	}

	return nil
}
