package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudentNoteType(ctx context.Context, data *model.StudentNoteType) error {
	panic("empty")
}
func (r *PostgresRepository) EditStudentNoteTypeById(ctx context.Context, id int64, data *model.StudentNoteType) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentNoteTypeById(ctx context.Context, is int64) error {
	panic("empty")
}
