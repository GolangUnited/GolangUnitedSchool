package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) AddInterview() {

}

func (r *PostgresRepository) GetInterviewById(ctx context.Context, id int64) (*model.Interview, error) {
	return nil, nil
}

func (r *PostgresRepository) GetInterviews(ctx context.Context) ([]*model.Interview, error) {
	return nil, nil
}

func (r *PostgresRepository) UpfateInterviewById(ctx context.Context, id int64, data *model.Interview) error {
	return nil
}

func (r *PostgresRepository) DeleteInterviewByID(ctx context.Context, id int64) error {
	return nil
}
