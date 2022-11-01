package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudentHomeworks(ctx context.Context) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetStudentHomeworksByStudentId(ctx context.Context, studentId int64) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetStudentHomeworkById(ctx context.Context, id int64) (*model.StudentHomework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddStudentHomework(ctx context.Context, data *model.StudentHomework) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateStudentHomework(ctx context.Context, id int64, data *model.StudentHomework) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteStudentHomework(ctx context.Context, id int64) error {
	panic("not implemented")
}
