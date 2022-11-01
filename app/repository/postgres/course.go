package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) CreateCourse(ctx context.Context, course *model.CourseCreate) (int64, error)
func (r *PostgresRepository) GetCourseByID(context.Context, int64) (model.Course, error)
func (r *PostgresRepository) GetCourses(context.Context) ([]model.Course, error)
func (r *PostgresRepository) UpdateCourseByID(context.Context, int64, *model.CourseUpdate) error
func (r *PostgresRepository) AddCourse(context.Context, *model.Course) (int64, error)
func (r *PostgresRepository) DeleteCourseByID(context.Context, int64) error
