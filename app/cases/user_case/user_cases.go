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

// UpdatePerson update person data by id
func (c *UserCases) UpdatePerson(ctx context.Context, id int64, newPerson *model.Person) error {
	person, err := c.r.GetPersonById(ctx, id)
	if err != nil {
		return fmt.Errorf("UserCases.UpdatePerson: %w", err)
	}

	if person == nil {
		var err = fmt.Errorf("PERSON WITH ID= %d NOT FOUND", id)
		return fmt.Errorf("UserCases.UpdatePerson: %w", err)
	}

	err = c.r.UpdatePerson(ctx, id, newPerson)
	if err != nil {
		return fmt.Errorf("UserCases.UpdatePerson: %w", err)
	}

	return nil
}

// NewUserCases construct UserCases
func NewUserCases(r repository.Repository) *UserCases {
	return &UserCases{
		r: r,
	}
}
