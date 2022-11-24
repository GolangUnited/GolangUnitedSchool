package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudents(ctx context.Context) (model.StudentsListDto, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudent(ctx context.Context, data *model.Student) error {
	panic("empty")
}
func (r *PostgresRepository) UpdateStudentByStudentId(ctx context.Context, id int64, data *model.Student) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentByStudentId(ctx context.Context, id int64) error {
	panic("empty")
}
func (r *PostgresRepository) PutStudentByStudentId(ctx context.Context, id int64, data *model.Student) error {
	panic("empty")
}
