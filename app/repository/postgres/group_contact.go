package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetGroupContacts(ctx cntxt) ([]m.GroupContact, error) {
	panic("empty")
}
func (r *PostgresRepository) GetGroupContactById(ctx cntxt, id int64) (*m.GroupContact, error) {
	panic("empty")
}
func (r *PostgresRepository) AddGroupContact(ctx, data *m.GroupContact) error {
	panic("empty")
}
func (r *PostgresRepository) EditGroupContactById(ctx cntxt, id int64, data *m.GroupContact) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteGroupContactById(ctx, id int64) error {
	panic("empty")
}
