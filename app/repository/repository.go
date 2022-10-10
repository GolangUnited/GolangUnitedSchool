package repository

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"

	"github.com/lozovoya/GolangUnitedSchool/app/domain"
)

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*domain.Person, error)
	PostNewPerson(ctx context.Context, n handlers.NewPersonQuery) (*domain.Person, error)
}
