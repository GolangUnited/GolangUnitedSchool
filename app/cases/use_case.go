package cases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/model"
)

type User_case interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	DeletePerson(ctx context.Context, id int64) error
}
