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

type OperationLogRepositoryInterface interface {
}

type OperationRepositoryInterface interface {
}

type OperationTypeRepositoryInterface interface {
}

type ContactTypeRepositoryInterface interface {
}
