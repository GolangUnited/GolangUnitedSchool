package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (u *repository) GetHomeworks(ctx context.Context) ([]model.Homework, error) {
	panic("not implemented")
}

func (u *repository) GetHomeworksByLectureId(ctx context.Context, lectureId int64) ([]model.Homework, error) {
	panic("not implemented")
}

func (u *repository) GetHomeworkById(ctx context.Context, id int64) (*model.Homework, error) {
	panic("not implemented")
}

func (u *repository) AddHomework(ctx context.Context, data *model.Homework) error {
	panic("not implemented")
}

func (u *repository) UpdateHomework(ctx context.Context, id int64, data *model.Homework) error {
	panic("not implemented")
}

func (u *repository) DeleteHomework(ctx context.Context, id int64) error {
	panic("not implemented")
}
