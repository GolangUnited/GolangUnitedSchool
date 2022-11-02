package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudentGroup(ctx context.Context, data *model.StudentGroup) error {
	panic("empty")
}
func (r *PostgresRepository) UpdateStudentGroupById(ctx context.Context, id int64, data *model.StudentGroup) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentGroupById(ctx context.Context, id int64) error {
	panic("empty")
}
