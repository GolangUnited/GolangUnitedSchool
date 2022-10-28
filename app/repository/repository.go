package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

type PersonRepositoryInterface interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
}

type CourseRepositoryInterface interface {
}

type ProjectRepositoryInterface interface {
}
