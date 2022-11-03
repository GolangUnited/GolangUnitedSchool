package course

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type LectureUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewLecture(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *LectureUsecase {
	return &LectureUsecase{lg: lg, repo: repo}
}

func (u *LectureUsecase) GetLectures(
	ctx context.Context) ([]model.Lecture, error) {
	panic("not implemented")
}

func (u *LectureUsecase) GetLectureById(
	ctx context.Context,
	id int64) (*model.Lecture, error) {
	panic("not implemented")
}

func (u *LectureUsecase) AddLecture(
	ctx context.Context,
	data *model.Lecture) error {
	panic("not implemented")
}

func (u *LectureUsecase) UpdateLecture(
	ctx context.Context,
	id int64,
	data *model.Lecture) error {
	panic("not implemented")
}

func (u *LectureUsecase) DeleteLecture(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
