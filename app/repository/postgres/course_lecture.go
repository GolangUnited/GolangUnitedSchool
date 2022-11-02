package postgres

import (
	m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

type cntxt context.Context

func (r *PostgresRepository) GetCourseLectures(ctx cntxt) ([]m.CourseLecture, error) {
	panic("empty")
}
func (r *PostgresRepository) GetCourseLectureById(ctx cntxt, id int64) (*m.CourseLecture, error) {
	panic("empty")
}
func (r *PostgresRepository) AddCourseLecture(ctx cntxt, data *m.CourseLecture) error {
	panic("empty")
}
func (r *PostgresRepository) EditCourseLectureById(ctx cntxt, id int64, data *m.CourseLecture) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteCourseLectureById(ctx cntxt, id int64) error {
	panic("empty")
}
