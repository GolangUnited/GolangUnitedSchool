package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetStudentNotes(ctx context.Context) ([]model.StudentNote, error) {
	var studentNote []model.StudentNote
	rows, err := r.pool.Query(ctx, `SELECT * FROM student_notes`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of student notes")
	}

	for rows.Next() {
		var c model.StudentNote
		err := rows.Scan(&c.Id,
			&c.StudentId,
			&c.StudentNoteTypeId,
			&c.Note,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan student notes")
		}

		studentNote = append(studentNote, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "group rows error")
	}

	return studentNote, nil
}
func (r *PostgresRepository) GetStudentNoteById(ctx context.Context, id int64) (*model.StudentNote, error) {
	var studentNote model.StudentNote
	err := r.pool.QueryRow(ctx, `
		SELECT 
		 	id,
			student_id, 
			student_note_type_id,   
			note,
			created_at
		FROM student_notes
		WHERE id = $1 
	`, id).Scan(
		&studentNote.Id,
		&studentNote.StudentId,
		&studentNote.StudentNoteTypeId,
		&studentNote.Note,
		&studentNote.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get student note by id: %d", id)
	}

	return &studentNote, nil
}
func (r *PostgresRepository) AddStudentNote(ctx context.Context, data *model.NewStudentNote) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.StudentId)
	keys = append(keys, "student_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.StudentNoteTypeId)
	keys = append(keys, "student_note_type_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.Note)
	keys = append(keys, "note")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.CreatedAt)
	keys = append(keys, "created_at")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO student_notes (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a student note")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateStudentNoteByStudentId(ctx context.Context, id int64, data *model.UpdateStudentNote) error {
	var args []interface{}
	var keys []string

	if data.StudentId != nil {
		args = append(args, data.StudentId)
		keys = append(keys, fmt.Sprintf("student_id = $%d", len(args)))
	}
	if data.StudentNoteTypeId != nil {
		args = append(args, data.StudentNoteTypeId)
		keys = append(keys, fmt.Sprintf("student_note_type_id = $%d", len(args)))
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

	query := fmt.Sprintf(`UPDATE student_notes
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update student note by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note doesn't exist"),
			"couldn't update student note by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutStudentNoteById(ctx context.Context, id int64, data *model.UpdateStudentNote) error {
	query := `UPDATE student_notes 
				SET 
					student_id = $1, 
					student_note_type_id = $2,
					note = $3,
					created_at = $4
				WHERE id=$5`

	cmn, err := r.pool.Exec(ctx, query,
		&data.StudentId,
		&data.StudentNoteTypeId,
		&data.Note,
		&data.CreatedAt,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put student note by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note doesn't exist"),
			"couldn't put student note by id: %d", id)
	}

	return nil
}
func (r *PostgresRepository) DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM student_notes WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete student note by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note doesn't exist"),
			"couldn't delete student note by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) GetStudentsNotesByStudentId(ctx context.Context, id int64) ([]model.StudentNote, error) {
	var studentNote []model.StudentNote
	rows, err := r.pool.Query(ctx, `SELECT * FROM student_notes WHERE student_id = $1 `, id)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of student's notes")
	}

	for rows.Next() {
		var c model.StudentNote
		err := rows.Scan(&c.Id,
			&c.StudentId,
			&c.StudentNoteTypeId,
			&c.Note,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan list of student's notes")
		}

		studentNote = append(studentNote, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "student's notes rows error")
	}

	return studentNote, nil
}
