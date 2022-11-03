package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetLectures(ctx context.Context) ([]model.Lecture, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetLectureById(ctx context.Context, id int64) (*model.Lecture, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddLecture(ctx context.Context, data *model.Lecture) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateLecture(ctx context.Context, id int64, data *model.Lecture) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteLecture(ctx context.Context, id int64) error {
	panic("not implemented")
}
