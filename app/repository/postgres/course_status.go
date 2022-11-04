package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetCourseStatuses(ctx context.Context) ([]model.CourseStatus, error) {
	panic("empty")
}
func (r *PostgresRepository) GetCourseStatusById(ctx context.Context, id int64) (*model.CourseStatus, error) {
	panic("empty")
}
func (r *PostgresRepository) AddCourseStatus(ctx context.Context, data *model.CourseStatus) error {
	panic("empty")
}
func (r *PostgresRepository) UpdateCourseStatusById(ctx context.Context, id int64, data *model.CourseStatus) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteCourseStatusById(ctx context.Context, id int64) error {
	panic("empty")
}
