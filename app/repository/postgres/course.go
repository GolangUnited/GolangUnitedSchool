package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

func (r *PostgresRepository) CreateCourse(ctx context.Context, course *model.CourseCreate) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// add title
	args = append(args, course.Title)
	keys = append(keys, "title")
	values = append(values, strconv.Itoa(len(args)))

	// add start date
	if course.StartDate != nil {
		args = append(args, course.StartDate)
		keys = append(keys, "start_date")
		values = append(values, strconv.Itoa(len(args)))
	}

	// add end date
	if course.EndDate != nil {
		args = append(args, course.EndDate)
		keys = append(keys, "end_date")
		values = append(values, strconv.Itoa(len(args)))
	}

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO courses(%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "create course")
	}

	return id, nil
}

func (r *PostgresRepository) GetCourseByID(ctx context.Context, id int64) (*model.Course, error) {
	var course model.Course
	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			title, 
			status,  
			start_date, 
			end_date 
		FROM course 
		WHERE id = $1 
	`, id).Scan(
		&course.ID,
		&course.Title,
		&course.Status,
		&course.StartDate,
		&course.EndDate,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "get course by id: %d", id)
	}

	return &course, nil
}

func (r *PostgresRepository) GetCourses(ctx context.Context, params *model.PaginationParams) (*model.CourseList, error) {

	offset := params.Page * params.PageSize
	limit := params.PageSize

	query := fmt.Sprintf(` SELECT  
					id, 
					title, 
					status, 
					start_date,
					end_date,
				FROM courses 
				WHERE offset = %d and limit %d`, offset, limit)
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get courses: could't make query")
	}

	courses := make([]model.Course, 0, params.PageSize)
	for rows.Next() {
		var course model.Course
		err := rows.Scan(
			&course.ID,
			&course.Status,
			&course.StartDate,
			&course.EndDate,
		)
		if err != nil {
			return nil, errors.Wrap(err, "get course: couldn't scan")
		}

		courses = append(courses, course)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "get course: rows error")
	}

	list := &model.CourseList{
		Courses: courses,
	}

	if params.WithMetadata {
		metadata := model.PaginationResponse{}
		metadata.TotalCount, err = r.getCoursesTotalCount(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get course")
		}
		metadata.Page = params.Page
		metadata.PageSize = params.PageSize
		metadata.PageCount = metadata.TotalCount / params.PageSize
		if metadata.TotalCount%params.PageSize != 0 {
			metadata.PageCount++
		}

		list.Metadata = metadata
	}

	return list, nil
}

func (r *PostgresRepository) UpdateCourseByID(
	ctx context.Context,
	id int64,
	course *model.CourseUpdate) error {
	return nil
}

func (r *PostgresRepository) AddCourse(ctx context.Context, course *model.Course) (int64, error) {
	query := ``

	var id int64
	err := r.pool.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "add course")
	}

	return id, nil
}

func (r *PostgresRepository) DeleteCourseByID(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM courses WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "delete course by id %d")
	}

	return nil
}

func (r *PostgresRepository) getCoursesTotalCount(ctx context.Context) (int64, error) {
	var count int64

	err := r.pool.QueryRow(ctx,
		`SELECT count(*) FROM courses`).
		Scan(&count)
	if err != nil {
		return count, errors.Wrap(err, "couldn't get total count")
	}

	return count, nil
}
