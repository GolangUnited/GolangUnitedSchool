package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (u *repository) GetLectures(ctx context.Context) ([]model.Lecture, error) {
	panic("not implemented")
}

func (u *repository) GetLectureById(ctx context.Context, id int64) (*model.Lecture, error) {
	panic("not implemented")
}

func (u *repository) AddLecture(ctx context.Context, data *model.Lecture) error {
	panic("not implemented")
}

func (u *repository) UpdateLecture(ctx context.Context, id int64, data *model.Lecture) error {
	panic("not implemented")
}

func (u *repository) DeleteLecture(ctx context.Context, id int64) error {
	panic("not implemented")
}
