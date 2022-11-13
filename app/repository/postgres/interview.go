package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
)

func (r *PostgresRepository) AddInterview(ctx context.Context, interview *model.Interview) (int64, error) {
	var id int64

	return id, nil
}

func (r *PostgresRepository) GetInterviewById(ctx context.Context, id int64) (*model.Interview, error) {
	var interview model.Interview
	query := `SELECT 
					id, 
					student_id, 
					mentor_id,
					note, 
					score, 
					created_at
			FROM interview
			where id=$1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		interview.ID,
		interview.StudentID,
		interview.MentorID,
		interview.Note,
		interview.Score,
		interview.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get interview")
	}

	return &interview, nil
}

func (r *PostgresRepository) GetInterviews(ctx context.Context) ([]model.Interview, error) {
	return nil, nil
}

func (r *PostgresRepository) UpdateInterviewById(ctx context.Context, id int64, data *model.Interview) error {
	return nil
}

func (r *PostgresRepository) DeleteInterviewById(ctx context.Context, id int64) error {
	return nil
}
