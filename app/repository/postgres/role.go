package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

var (
	ErrorRoleDoesntExists = "role doesn't exists"
)

func (r *PostgresRepository) GetRoleById(ctx context.Context, id int64) (*model.Role, error) {
	query := `SELECT 
				id, 
				role_name, 
				is_read_inly 
			  FROM role
			  WHERE id = $1`
	var role model.Role
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&role.Id,
		&role.RoleName,
		&role.IsReadOnly,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get role by id %d", id)
	}

	return &role, nil
}

func (r *PostgresRepository) GetRoles(ctx context.Context) ([]model.Role, error) {
	query := `SELECT 
					id, 
					role_name, 
					is_read_only 
				FROM role`
	var roles []model.Role
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't query get all roles")
	}
	defer rows.Close()

	for rows.Next() {
		var r model.Role
		err := rows.Scan(&r.Id, &r.RoleName, &r.IsReadOnly)
		if err != nil {
			return nil, errors.Wrap(err, "get roles scan error")
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get roles rows error")
	}
	return roles, nil
}

func (r *PostgresRepository) AddRole(ctx context.Context, role *model.RoleCU) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// role name
	if role.RoleName != nil {
		args = append(args, role.RoleName)
	}

	query := fmt.Sprintf(`INSERT INTO role (%s) VALUES (%s) RETURBING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)

	var id int64
	err := r.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "couldn't create role")
	}

	return id, nil
}

func (r *PostgresRepository) PutRoleById(ctx context.Context, id int64, role *model.Role) error {
	ptg, err := r.pool.Exec(ctx, `UPDATE role SET (
		role_name=$1, 
		"is_read_only=$2"
		) WHERE id=$3`, role.RoleName, role.IsReadOnly, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't put role by id %d", id)
	}
	if ptg.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorRoleDoesntExists), "couldn't put role id %d", id)
	}

	return nil
}

func (r *PostgresRepository) UpdateRoleById(ctx context.Context, id int64, role *model.RoleCU) error {
	var args []interface{}
	var keys []string
	// role name
	if role.RoleName != nil {
		args = append(args, role.RoleName)
		keys = append(keys, fmt.Sprintf("role_name = %d", len(args)))
	}

	// is read only
	if role.IsReadOnly != nil {
		args = append(args, role.IsReadOnly)
		keys = append(keys, fmt.Sprintf("is_read_only = %d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE role 
				SET %s 
				WHERE id = &%d`, strings.Join(keys, ", "), id)

	ptg, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err, "couldn't update role id %d", id)
	}
	if ptg.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorRoleDoesntExists),
			"couldn't update role id %d", id)
	}
	return nil
}

func (r *PostgresRepository) DeleteRoleById(ctx context.Context, id int64) error {
	ptg, err := r.pool.Exec(ctx, `DELETE FROM role WHERE id=$1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete role id %d", id)
	}
	if ptg.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorRoleDoesntExists),
			"couldn't delete role id %d", id)
	}
	return nil
}
