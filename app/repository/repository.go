package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/model"
)

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	DeletePerson(ctx context.Context, id int64) error
}
