package postgres

import (
	"context"
	"github.com/pkg/errors"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetPersons(ctx context.Context) ([]model.Person, error) {
	var persons []model.Person
	rows, err := r.pool.Query(ctx, `SELECT first_name, last_name, updated_at FROM person`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of persons")
	}

	for rows.Next() {
		var c model.Person
		err := rows.Scan(&c.FirstName, &c.LastName, &c.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan person")
		}

		persons = append(persons, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "person rows error")
	}

	return persons, nil

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
