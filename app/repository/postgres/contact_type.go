package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetContactTypes(
	ctx context.Context) (*model.ContactType, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetContactTypeById(
	ctx context.Context, id int64) (*model.ContactType, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddContactType(
	ctx context.Context) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateContactType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteContactType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
