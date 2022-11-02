package course

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type HomeworkUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewHomework(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *HomeworkUsecase {
	return &HomeworkUsecase{lg: lg, repo: repo}
}

func (u *HomeworkUsecase) GetHomeworks(
	ctx context.Context) ([]model.Homework, error) {
	panic("not implemented")
}

func (u *HomeworkUsecase) GetHomeworksByLectureId(
	ctx context.Context,
	lectureId int64) ([]model.Homework, error) {
	panic("not implemented")
}

func (u *HomeworkUsecase) GetHomeworkById(
	ctx context.Context,
	id int64) (*model.Homework, error) {
	panic("not implemented")
}

func (u *HomeworkUsecase) AddHomework(
	ctx context.Context,
	data *model.Homework) error {
	panic("not implemented")
}

func (u *HomeworkUsecase) UpdateHomework(
	ctx context.Context,
	id int64,
	data *model.Homework) error {
	panic("not implemented")
}

func (u *HomeworkUsecase) DeleteHomework(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
