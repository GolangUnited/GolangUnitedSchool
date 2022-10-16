package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/internal/models"
)

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*models.Person, error)
}
