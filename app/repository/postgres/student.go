package postgres

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudents(ctx context.Context) ([]model.Student, error) {
	var students []model.Student
	rows, err := r.pool.Query(ctx, `SELECT * FROM students`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of students")
	}

	for rows.Next() {
		var c model.Student
		err := rows.Scan(&c.Id, &c.StudentId)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan students")
		}

		students = append(students, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "student rows error")
	}

	return students, nil
}
func (r *PostgresRepository) GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error) {
	var student model.Student
	err := r.pool.QueryRow(ctx, `
		SELECT
			*
		FROM students
		WHERE id = $1
	`, id).Scan(
		&student.Id,
		&student.StudentId,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get student by id: %d", id)
	}

	return &student, nil

}
func (r *PostgresRepository) AddStudent(ctx context.Context, data *model.Student) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.StudentId)
	keys = append(keys, "student_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO students (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't add new student")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error {
	_, err := r.pool.Exec(
		ctx,
		`UPDATE students 
			SET student_id=$1 
			WHERE id=$2`,
		data.StudentId, id)
	if err != nil {
		return errors.Wrap(err, "couldn't update course status")
	}

	return nil
}
func (r *PostgresRepository) DeleteStudentByStudentId(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM students WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete student by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student doesn't exist"),
			"couldn't delete student by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error {
	query := `UPDATE students 
				SET 
					student_id = $1
				WHERE id = $2`

	cmn, err := r.pool.Exec(ctx, query,
		&data.StudentId,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put student by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student doesn't exist"),
			"couldn't put student by id: %d", id)
	}
	return nil
}

//GetStudents(ctx context.Context) ([]model.Student, error)
//GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error)
//AddStudent(ctx context.Context, data *model.Student) (int64, error)
//UpdateStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error
//DeleteStudentByStudentId(ctx context.Context, id int64) error
//PutStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error
