package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetMentorNotes(ctx cntxt) ([]m.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) GetMentorNotesByMentorId(ctx cntxt, id int64) ([]m.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) GetMentorNoteByMentorNoteId(ctx cntxt, id int64) (*m.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) AddMentorNote(ctx cntxt, data *m.MentorNote) error {
	panic("Empty")
}
func (r *PostgresRepository) EditMentorNoteByMentorNoteId(ctx cntxt, id int64, data *m.MentorNote) error {
	panic("Empty")
}
func (r *PostgresRepository) DeleteMentorNoteByMentorNoteId(ctx cntxt, id int64) error {
	panic("Empty")
}
