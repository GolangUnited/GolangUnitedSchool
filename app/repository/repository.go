package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

// RepositoryInterface is main interface for all repo implementations
type RepositoryInterface interface {

	// Course repo interfaces
	CreateCourse(context.Context, *model.CourseCreate) (int64, error)
	GetCourseByID(context.Context, int64) (*model.Course, error)
	GetCourses(context.Context, *model.PaginationParams) (*model.CourseList, error)
	UpdateCourseByID(context.Context, int64, *model.CourseUpdate) error
	PutCourseByID(context.Context, int64, *model.CourseUpdate) error
	DeleteCourseByID(context.Context, int64) error
}
