package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudentNotes(ctx context.Context) ([]model.StudentNote, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentNoteByStudentId(ctx context.Context, id int64) (*model.StudentNote, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudentNote(ctx context.Context, data *model.StudentNote) error {
	panic("empty")
}
func (r *PostgresRepository) EditStudentNoteByStudentId(ctx context.Context, id int64, data *model.StudentNote) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error {
	panic("empty")
}
