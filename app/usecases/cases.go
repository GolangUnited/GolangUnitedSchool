package usecases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type Cases struct {
	Auth   Auth
	User   User
	Course Course
}

func NewUseCases(r repository.Repository) *Cases {
	return &Cases{
		User:   NewUserCases(r),
		Course: NewCourseCases,
	}
}

type User interface {
	GetPersonById(ctx context.Context, id int64) (*domain.Person, error)
}

type Course interface {
}

type Auth interface {
}
