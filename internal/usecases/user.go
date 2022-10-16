package usecases

import (
	"context"
	"fmt"

	"github.com/lozovoya/GolangUnitedSchool/internal/models"
	"github.com/lozovoya/GolangUnitedSchool/internal/repository"
)

type UserSt struct {
	r repository.Repository
}

// NewUserCases construct UserCases
func NewUserCases(r repository.Repository) User {
	return &UserSt{
		r: r,
	}
}

// Realize User_case interface
// GetPersonById return person data by id
func (c *UserSt) GetPersonById(ctx context.Context, id int64) (*models.Person, error) {
	person, err := c.r.GetPersonById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserCases.GetPersonById: %w", err)
	}
	return person, nil
}
