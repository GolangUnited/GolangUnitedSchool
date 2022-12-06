package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

func (r *PostgresRepository) GetCourseStatuses(ctx context.Context) ([]model.CourseStatus, error) {
	var statuses []model.CourseStatus
	rows, err := r.pool.Query(
		ctx,
		`SELECT 
			id, 
			title 
		FROM course_status`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get course statuses")
	}
	defer rows.Close()

	for rows.Next() {
		var c model.CourseStatus
		err := rows.Scan(&c.Id, &c.Title)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan course status")
		}

		statuses = append(statuses, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "course statuses rows error")
	}

	return statuses, nil
}

func (r *PostgresRepository) GetCourseStatusById(ctx context.Context, id int64) (*model.CourseStatus, error) {
	var c model.CourseStatus
	err := r.pool.QueryRow(
		ctx,
		`SELECT 
			id, 
			title 
		FROM course_status 
		WHERE id=$1`, id).
		Scan(&c.Id, &c.Title)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get course status")
	}
	return &c, nil
}

func (r *PostgresRepository) AddCourseStatus(ctx context.Context, data *model.CourseStatus) (int64, error) {
	var id int64
	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO course_status 
			(title) VALUES ($1) 
		RETURNING id`, data.Title).
		Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "couldn't add course status")
	}

	return id, nil
}

func (r *PostgresRepository) UpdateCourseStatusById(ctx context.Context, id int64, data *model.CourseStatus) error {

	_, err := r.pool.Exec(
		ctx,
		`UPDATE course_status 
			SET title=$1 
			WHERE id=$2`,
		data.Title, id)
	if err != nil {
		return errors.Wrap(err, "couldn't update course status")
	}

	return nil
}
func (r *PostgresRepository) DeleteCourseStatusById(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(
		ctx,
		`DELETE 
			FROM course_status 
			WHERE id=$1`, id)
	if err != nil {
		return errors.Wrap(err, "couldn't delete course status")
	}

	return nil
}
