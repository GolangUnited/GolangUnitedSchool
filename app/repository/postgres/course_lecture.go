package postgres

import (
	"fmt"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

var (
	ErrorCourseLectureDoesntExists  = "course lecture doesn't exists"
	ErrorCourseLectureAlreadyExists = "course lecture already exists"
)

func (r *PostgresRepository) GetCourseLectures(ctx context.Context, params model.CourseLectureListParams) ([]model.CourseLecture, error) {
	var args []interface{}
	qWhere := "WHERE 1=1"
	if params.CourseId != nil {
		args = append(args, params.CourseId)
		qWhere += fmt.Sprintf(" AND course_id=&%d", len(args))
	}
	if params.LectureId != nil {
		args = append(args, params.LectureId)
		qWhere += fmt.Sprintf(" AND lecture_id=&%d", len(args))
	}

	query := fmt.Sprintf(`SELECT course_id, lecture_id 
				FROM course_lecture
				%s`, qWhere)

	r.lg.Debug(query)

	var courses []model.CourseLecture
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get course lectures query error")
	}
	defer rows.Close()

	for rows.Next() {
		var course model.CourseLecture
		err := rows.Scan(&course.CourseId, &course.LectureId)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan course lectures")
		}

		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "couldn't get course lectures")
	}

	return courses, nil
}

func (r *PostgresRepository) GetCourseLectureById(ctx context.Context, course_id, lecture_id int64) (*model.CourseLecture, error) {
	var data model.CourseLecture
	err := r.pool.QueryRow(ctx,
		`SELECT course_id, lecture_id 
			FROM course_lecture 
			WHERE course_id=$1 AND lecture_id=$2`,
		course_id, lecture_id).
		Scan(&data.CourseId, &data.LectureId)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get course lecture relation by id")
	}

	return &data, nil
}

func (r *PostgresRepository) AddCourseLecture(ctx context.Context, data *model.CourseLecture) error {
	exists, err := r.IsExistsCourseLecture(ctx, data)
	if err != nil {
		return errors.Wrap(err, "couldn't create course lecture relation")
	}
	if exists {
		return errors.Wrap(errors.New(ErrorCourseLectureAlreadyExists),
			"couldn't create course lecture relation")
	}

	_, err = r.pool.Exec(ctx,
		`INSERT INTO course_lecture 
			(course_id, lecture_id) 
			VALUES ($1, $2)`)
	if err != nil {
		errors.Wrap(err, "couldn't create course lecture relation")
	}

	return nil
}

func (r *PostgresRepository) UpdateCourseLecture(ctx context.Context, course_id, lecture_id int64, data *model.CourseLecture) error {
	pct, err := r.pool.Exec(ctx,
		`UPDATE course_lecture 
			SET course_id=$1, lecture_id=$2
			WHERE course_id=$3, lecture_id=$4`,
		data.CourseId, data.LectureId,
		course_id, lecture_id,
	)
	if err != nil {
		return errors.Wrap(err, "couldn't update course lecture relation")
	}
	if pct.RowsAffected() != 1 {
		return errors.Wrap(errors.New(ErrorCourseLectureDoesntExists),
			"couldn't update course lecture relation")
	}

	return nil

}

func (r *PostgresRepository) DeleteCourseLecture(ctx context.Context, data *model.CourseLecture) error {
	pct, err := r.pool.Exec(ctx,
		`DELETE FROM course_lecture 
			WHERE course_id=$1 and lecture_id=$2`,
		data.CourseId, data.LectureId)
	if err != nil {
		return errors.Wrap(err, "couldn'r delete course lecture relation")
	}

	if pct.RowsAffected() != 1 {
		return errors.Wrap(errors.New(ErrorCourseLectureDoesntExists),
			"couldn't delete course lecture relation")
	}

	return nil
}

func (r *PostgresRepository) IsExistsCourseLecture(ctx context.Context, data *model.CourseLecture) (bool, error) {
	var count int
	err := r.pool.QueryRow(ctx,
		`SELECT count(*) FROM course_lecture 
			WHERE course_id=$1 AND lecture_id=$2`,
		data.CourseId, data.LectureId).
		Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}
