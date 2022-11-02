package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetStudentNotes(ctx cntxt) ([]m.StudentNote, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentNoteByStudentId(ctx cntxt, id int64) (*m.StudentNote, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudentNote(ctx cntxt, data *m.StudentNote) error {
	panic("empty")
}
func (r *PostgresRepository) EditStudentNoteByStudentId(ctx cntxt, id int64, data *m.StudentNote) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentNoteByStudentNoteId(ctx cntxt, id int64) error {
	panic("empty")
}
