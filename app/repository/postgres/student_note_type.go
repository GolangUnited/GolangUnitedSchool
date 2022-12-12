package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error) {
	var types []model.StudentNoteType
	rows, err := r.pool.Query(ctx, `SELECT * FROM student_note_types`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of student note types")
	}

	for rows.Next() {
		var c model.StudentNoteType
		err := rows.Scan(&c.StudentNoteTypeId,
			&c.Title,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan student note types")
		}

		types = append(types, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "group rows error")
	}

	return types, nil
}
func (r *PostgresRepository) GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error) {
	var noteType model.StudentNoteType
	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			title
		FROM student_note_types
		WHERE id = $1 
	`, id).Scan(
		&noteType.StudentNoteTypeId,
		&noteType.Title,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get student note type by id: %d", id)
	}

	return &noteType, nil
}
func (r *PostgresRepository) AddStudentNoteType(ctx context.Context, data *model.NewStudentNoteType) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.Title)
	keys = append(keys, "title")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO student_note_types (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a group contact")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error {
	var args []interface{}
	var keys []string

	if data.Title != nil {
		args = append(args, data.Title)
		keys = append(keys, fmt.Sprintf("title = $%d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE student_note_types
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update student note type by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note type doesn't exist"),
			"couldn't update student note type by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error {
	var args []interface{}
	var keys []string

	if data.Title != nil {
		args = append(args, data.Title)
		keys = append(keys, fmt.Sprintf("title = $%d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE student_note_types
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update student note type by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note type doesn't exist"),
			"couldn't update student note type by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) DeleteStudentNoteTypeById(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM student_note_types WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete student note type by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student note type doesn't exist"),
			"couldn't delete student note type by id %d", id)
	}
	return nil
}

//GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error)
//GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error)
//AddStudentNoteType(ctx context.Context, data *model.NewStudentNoteType) (int64, error)
//UpdateStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error
//PutStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error
//DeleteStudentNoteTypeById(ctx context.Context, id int64) error
