package cases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain"
)

type User_case interface {
	GetPersonById(ctx context.Context, id int64) (*domain.Person, error)
}
