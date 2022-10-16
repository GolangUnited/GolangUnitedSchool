package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/internal/domain"
)

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*domain.Person, error)
}
