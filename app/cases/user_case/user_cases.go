package user_case

import (
	"context"
	"fmt"

	"github.com/lozovoya/GolangUnitedSchool/app/model"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type UserCases struct {
	r repository.Repository
}

// Realize User_case interface
// GetPersonById return person data by id
func (c *UserCases) GetPersonById(ctx context.Context, id int64) (*model.Person, error) {
	person, err := c.r.GetPersonById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserCases.GetPersonById: %w", err)
	}
	return person, nil
}

func (c *UserCases) DeletePerson(ctx context.Context, id int64) error {
	return c.r.DeletePerson(ctx, id)
}

// NewUserCases construct UserCases
func NewUserCases(r repository.Repository) *UserCases {
	return &UserCases{
		r: r,
	}
}
