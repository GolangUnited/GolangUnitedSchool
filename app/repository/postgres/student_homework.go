package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (u *repository) GetStudentHomeworks(ctx context.Context) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (u *repository) GetStudentHomeworksByStudentId(ctx context.Context, studentId int64) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (u *repository) GetStudentHomeworkById(ctx context.Context, id int64) (*model.StudentHomework, error) {
	panic("not implemented")
}

func (u *repository) AddStudentHomework(ctx context.Context, data *model.StudentHomework) error {
	panic("not implemented")
}

func (u *repository) UpdateStudentHomework(ctx context.Context, id int64, data *model.StudentHomework) error {
	panic("not implemented")
}

func (u *repository) DeleteStudentHomework(ctx context.Context, id int64) error {
	panic("not implemented")
}
