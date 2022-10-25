package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
}
