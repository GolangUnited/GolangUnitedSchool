package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

var (
	ErrorBadPaginationParams = "bad pagination params"
	ErrorCourseDoesNotExists = "course doesn't exists"
	ErrorPageDoesNotExists   = "page doesn't exists"
)

func (r *PostgresRepository) CreateCourse(ctx context.Context, course *model.CourseCreate) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// add title
	args = append(args, course.Title)
	keys = append(keys, "title")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// add start date
	if course.StartDate != nil {
		args = append(args, course.StartDate)
		keys = append(keys, "start_date")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	// add end date
	if course.EndDate != nil {
		args = append(args, course.EndDate)
		keys = append(keys, "end_date")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO course (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a course")
	}

	return id, nil
}

func (r *PostgresRepository) GetCourseByID(ctx context.Context, id int64) (*model.Course, error) {
	var course model.Course
	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			title, 
			course_status_id,  
			start_date, 
			end_date 
		FROM course 
		WHERE id = $1 
	`, id).Scan(
		&course.Id,
		&course.Title,
		&course.StatusId,
		&course.StartDate,
		&course.EndDate,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get course by id: %d", id)
	}

	return &course, nil
}

func (r *PostgresRepository) GetCourses(ctx context.Context, params *model.PaginationParams) (*model.CourseList, error) {
	// check if page and page size ok
	if params == nil || params.Page <= 0 || params.PageSize <= 0 {
		return nil, errors.New(ErrorBadPaginationParams)
	}

	offset := (params.Page - 1) * params.PageSize
	limit := params.PageSize

	totalCount, err := r.getCoursesTotalCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get course: couldn't get total count")
	}

	// if page doesn't exists
	if offset >= totalCount {
		return nil, errors.Wrap(errors.New(ErrorPageDoesNotExists),
			"get course: out of the pages range")
	}

	query := fmt.Sprintf(`SELECT  
					id, 
					title, 
					course_status_id, 
					start_date,
					end_date
				FROM course 
				OFFSET %d LIMIT %d`, offset, limit)
	r.lg.Debug(query)
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get courses: could't make query")
	}

	courses := make([]model.Course, 0, params.PageSize)
	for rows.Next() {
		var course model.Course
		err := rows.Scan(
			&course.Id,
			&course.Title,
			&course.StatusId,
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
		metadata.TotalCount = totalCount

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
	var args []interface{}
	var keys []string

	// title
	if course.Title != nil {
		args = append(args, course.Title)
		keys = append(keys, fmt.Sprintf("title = $%d", len(args)))
	}

	// status
	if course.StatusId != nil {
		args = append(args, course.StatusId)
		keys = append(keys, fmt.Sprintf("course_status_id = $%d", len(args)))
	}

	// start date
	if course.StartDate != nil {
		args = append(args, course.StartDate)
		keys = append(keys, fmt.Sprintf("start_date = $%d", len(args)))
	}

	// end date
	if course.EndDate != nil {
		args = append(args, course.EndDate)
		keys = append(keys, fmt.Sprintf("end_date = $%d", len(args)))
	}

	// id
	args = append(args, id)

	query := fmt.Sprintf(`UPDATE course 
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update course by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorCourseDoesNotExists),
			"couldn't update course by id %d", id)
	}
	return nil
}

func (r *PostgresRepository) PutCourseByID(ctx context.Context, id int64, course *model.CourseUpdate) error {
	query := `UPDATE course 
				SET 
					title = $1, 
					course_status_id = $2, 
					start_date = $3, 
					end_date = $4
				WHERE id=$5`

	cmn, err := r.pool.Exec(ctx, query,
		course.Title,
		course.StatusId,
		course.StartDate,
		course.EndDate,
		id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put course by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorCourseDoesNotExists),
			"couldn't put course by id: %d", id)
	}

	return nil
}

func (r *PostgresRepository) DeleteCourseByID(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM course WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete course by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New(ErrorCourseDoesNotExists),
			"couldn't delete course by id %d", id)
	}
	return nil
}

func (r *PostgresRepository) getCoursesTotalCount(ctx context.Context) (int64, error) {
	var count int64

	err := r.pool.QueryRow(ctx,
		`SELECT count(*) FROM course`).
		Scan(&count)
	if err != nil {
		return count, errors.Wrap(err, "couldn't get total count")
	}

	return count, nil
}
