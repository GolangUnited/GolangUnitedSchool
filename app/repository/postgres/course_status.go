package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetCourseStatuses(ctx cntxt) ([]m.CourseStatus, error) {
	panic("empty")
}
func (r *PostgresRepository) GetCourseStatusById(ctx cntxt, id int64) (*m.CourseStatus, error) {
	panic("empty")
}
func (r *PostgresRepository) AddCourseStatus(ctx cntxt, data *m.CourseStatus) error {
	panic("empty")
}
func (r *PostgresRepository) EditCourseStatusById(ctx cntxt, id int64, data *m.CourseStatus) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteCourseStatusById(ctx cntxt, id int64) error {
	panic("empty")
}
