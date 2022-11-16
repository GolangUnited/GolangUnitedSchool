package interview

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"github.com/pkg/errors"
)

type InterviewUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewInterviewUsecase(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *InterviewUsecase {
	return &InterviewUsecase{
		lg:   lg,
		repo: repo,
	}
}

func (u *InterviewUsecase) GetInterviewById(ctx context.Context, id int64) (*model.Interview, error) {
	interview, err := u.repo.GetInterviewById(ctx, id)
	if err != nil {
		return nil, err
	}

	return interview, err
}

func (u *InterviewUsecase) GetInterviews(ctx context.Context) ([]model.Interview, error) {
	interviews, err := u.repo.GetInterviews(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get interviews")
	}

	return interviews, nil
}

func (u *InterviewUsecase) AddInterview(context.Context, *model.Interview) (int64, error)
func (u *InterviewUsecase) UpdateInterviewById(context.Context, int64, *model.Interview) error
func (u *InterviewUsecase) DeleteInterviewById(context.Context, int64) error
