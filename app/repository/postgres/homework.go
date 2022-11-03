package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetHomeworks(ctx context.Context) ([]model.Homework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetHomeworksByLectureId(ctx context.Context, lectureId int64) ([]model.Homework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetHomeworkById(ctx context.Context, id int64) (*model.Homework, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddHomework(ctx context.Context, data *model.Homework) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateHomework(ctx context.Context, id int64, data *model.Homework) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteHomework(ctx context.Context, id int64) error {
	panic("not implemented")
}
