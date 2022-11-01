package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

// RepositoryInterface is main interface for all repo implementations
type RepositoryInterface interface {
	CreateCourse(context.Context, *model.CourseCreate) (int64, error)
	GetCourseByID(context.Context, int64) (model.Course, error)
	GetCourses(context.Context) ([]model.Course, error)
	UpdateCourseByID(context.Context, int64, *model.CourseUpdate) error
	AddCourse(context.Context, *model.Course) (int64, error)
	DeleteCourseByID(context.Context, int64) error
}
