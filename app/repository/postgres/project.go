package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetProjects(ctx context.Context) ([]model.Project, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetProjectById(ctx context.Context, id int64) (*model.Project, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddProject(ctx context.Context, data *model.Project) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateProject(ctx context.Context, id int64, data *model.Project) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteProject(ctx context.Context, id int64) error {
	panic("not implemented")
}