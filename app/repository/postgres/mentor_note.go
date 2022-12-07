package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetMentorNotes(ctx context.Context) ([]model.MentorNote, error) {
	var mentorNotes []model.MentorNote
	rows, err := r.pool.Query(ctx, `SELECT * FROM mentor_notes`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of mentor notes")
	}

	for rows.Next() {
		var c model.MentorNote
		err := rows.Scan(&c.MentorNoteId,
			&c.StudentId,
			&c.MentorId,
			&c.Note,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan mentor notes")
		}

		mentorNotes = append(mentorNotes, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "mentor note rows error")
	}

	return mentorNotes, nil
}
func (r *PostgresRepository) GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error) {
	panic("empty")
}
func (r *PostgresRepository) GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error) {
	var mentorNote model.MentorNote
	err := r.pool.QueryRow(ctx, `
		SELECT 
			mentor_note_id, 
			student_id,
			mentor_id,
			note,
			created_at
		FROM mentor_notes
		WHERE mentor_note_id = $1 
	`, id).Scan(
		&mentorNote.MentorNoteId,
		&mentorNote.StudentId,
		&mentorNote.MentorId,
		&mentorNote.Note,
		&mentorNote.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get mentor note by id: %d", id)
	}

	return &mentorNote, nil
}
func (r *PostgresRepository) AddMentorNote(ctx context.Context, data *model.NewMentorNote) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.StudentId)
	keys = append(keys, "student_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.MentorId)
	keys = append(keys, "mentor_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.Note)
	keys = append(keys, "note")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.CreatedAt)
	keys = append(keys, "created_at")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO mentor_notes (%s) VALUES (%s) RETURNING mentor_note_id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a mentor note")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error {

	var args []interface{}
	var keys []string

	if data.StudentId != nil {
		args = append(args, data.StudentId)
		keys = append(keys, fmt.Sprintf("student_id = $%d", len(args)))
	}
	if data.MentorId != nil {
		args = append(args, data.MentorId)
		keys = append(keys, fmt.Sprintf("mentor_id = $%d", len(args)))
	}
	if data.Note != nil {
		args = append(args, data.Note)
		keys = append(keys, fmt.Sprintf("note = $%d", len(args)))
	}
	if data.CreatedAt != nil {
		args = append(args, data.CreatedAt)
		keys = append(keys, fmt.Sprintf("created_at = $%d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE mentor_notes
				SET %s
				WHERE mentor_note_id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update mentor note by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor note doesn't exist"),
			"couldn't update mentor note by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM mentor_notes WHERE mentor_note_id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete mentor note by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor note doesn't exist"),
			"couldn't delete mentor note by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error {
	query := `UPDATE mentor_notes
				SET 
					student_id = $1, 
					mentor_id = $2,
					note = $3,
					created_at = $4
				WHERE mentor_note_id =$5`

	cmn, err := r.pool.Exec(ctx, query,
		&data.StudentId,
		&data.MentorId,
		&data.Note,
		&data.CreatedAt,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put mentor note by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor note doesn't exist"),
			"couldn't put mentor note by id: %d", id)
	}

	return nil
}
