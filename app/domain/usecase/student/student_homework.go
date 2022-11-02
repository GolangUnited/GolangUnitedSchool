package student

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentHomeworkUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentHomework(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentHomeworkUsecase {
	return &StudentHomeworkUsecase{lg: lg, repo: repo}
}

func (u *StudentHomeworkUsecase) GetStudentHomeworks(
	ctx context.Context) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (u *StudentHomeworkUsecase) GetStudentHomeworksByStudentId(
	ctx context.Context,
	studentId int64) ([]model.StudentHomework, error) {
	panic("not implemented")
}

func (u *StudentHomeworkUsecase) GetStudentHomeworkById(
	ctx context.Context,
	id int64) (*model.StudentHomework, error) {
	panic("not implemented")
}

func (u *StudentHomeworkUsecase) AddStudentHomework(
	ctx context.Context,
	data *model.StudentHomework) error {
	panic("not implemented")
}

func (u *StudentHomeworkUsecase) UpdateStudentHomework(
	ctx context.Context,
	id int64,
	data *model.StudentHomework) error {
	panic("not implemented")
}

func (u *StudentHomeworkUsecase) DeleteStudentHomework(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
