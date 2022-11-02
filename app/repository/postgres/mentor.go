package postgres

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetMentors(ctx context.Context) ([]model.Mentor, error) {
	panic("empty")
}
func (r *PostgresRepository) GetMentorById(ctx context.Context, id int64) (*model.Mentor, error) {
	panic("empty")
}
func (r *PostgresRepository) AddMentor(ctx context.Context, data *model.Mentor) error {
	panic("empty")
}
func (r *PostgresRepository) UpdateMentorByMentorId(ctx context.Context, id int64, data *model.Mentor) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteMentorByMentorId(ctx context.Context, id int64) error {
	panic("empty")
}
