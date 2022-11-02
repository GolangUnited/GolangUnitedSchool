package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetPersons(ctx context.Context) ([]model.Person, error) {
	panic("empty")
}
func (r *PostgresRepository) GetPersonById(ctx context.Context, id int64) (*model.Person, error) {
	panic("empty")
}
func (r *PostgresRepository) AddNewPerson(ctx context.Context, data *model.Person) error {
	panic("empty")
}
func (r *PostgresRepository) EditPersonById(ctx context.Context, id int64, data *model.Person) error {
	panic("empty")
}
func (r *PostgresRepository) DeletePersonById(ctx context.Context, id int64) error {
	panic("empty")
}
