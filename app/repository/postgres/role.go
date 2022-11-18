package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetRoleById(ctx context.Context, id int64) (*model.Role, error) {
	return nil, nil
}

func (r *PostgresRepository) GetRoles(ctx context.Context) ([]model.Role, error) {
	return nil, nil
}

func (r *PostgresRepository) AddRole(ctx context.Context, role *model.RoleCU) (int64, error) {
	return 0, nil
}

func (r *PostgresRepository) PutRoleById(ctx context.Context, id int64, role *model.Role) error {
	return nil
}

func (r *PostgresRepository) UpadateRoleById(ctx context.Context, id int64, role *model.RoleCU) error {
	return nil
}

func (r *PostgresRepository) DeleteRoleById(ctx context.Context, id int64) error {
	return nil
}
