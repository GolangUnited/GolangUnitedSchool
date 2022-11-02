package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetMentorNotes(ctx context.Context) ([]model.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error) {
	panic("Empty")
}
func (r *PostgresRepository) AddMentorNote(ctx context.Context, data *model.MentorNote) error {
	panic("Empty")
}
func (r *PostgresRepository) EditMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.MentorNote) error {
	panic("Empty")
}
func (r *PostgresRepository) DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error {
	panic("Empty")
}
