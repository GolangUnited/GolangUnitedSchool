package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetPersons(ctx context.Context) (model.PersonListDto, error) {
	panic("empty")
}
func (r *PostgresRepository) GetPersonById(ctx context.Context, id int64) (*model.Person, error) {
	panic("empty")
}
func (r *PostgresRepository) AddNewPerson(ctx context.Context, data *model.NewPersonDto) error {
	panic("empty")
}
func (r *PostgresRepository) UpdatePersonById(ctx context.Context, id int64, data model.UpdatePersonDto) error {
	panic("empty")
}
func (r *PostgresRepository) PutPersonById(ctx context.Context, id int64, data *model.Person) error {
	panic("empty")
}
func (r *PostgresRepository) DeletePersonById(ctx context.Context, id int64) error {
	panic("empty")
}
