package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetGroupContacts(ctx context.Context) ([]model.GroupContact, error) {
	panic("empty")
}
func (r *PostgresRepository) GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error) {
	panic("empty")
}
func (r *PostgresRepository) AddGroupContact(ctx context.Context, data *model.GroupContact) error {
	panic("empty")
}
func (r *PostgresRepository) EditGroupContactById(ctx context.Context, id int64, data *model.GroupContact) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteGroupContactById(ctx context.Context, id int64) error {
	panic("empty")
}
