package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetPersons(ctx cntxt) ([]m.Person, error) {
	panic("empty")
}
func (r *PostgresRepository) GetPersonById(ctx cntxt, id int64) (*m.Person, error) {
	panic("empty")
}
func (r *PostgresRepository) AddNewPerson(ctx cntxt, data *m.Person) error {
	panic("empty")
}
func (r *PostgresRepository) EditPersonById(ctx cntxt, id int64, data *m.Person) error {
	panic("empty")
}
func (r *PostgresRepository) DeletePersonById(ctx cntxt, id int64) error {
	panic("empty")
}
