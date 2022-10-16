package usecases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/internal/models"
	"github.com/lozovoya/GolangUnitedSchool/internal/repository"
)

type Cases struct {
	Auth   Auth
	User   User
	Course Course
}

func NewUseCases(r repository.Repository) *Cases {
	return &Cases{
		User:   NewUserCases(r),
		Course: NewCourseCases(r),
	}
}

type User interface {
	GetPersonById(ctx context.Context, id int64) (*models.Person, error)
}

type Course interface {
}

type Auth interface {
}
