package group

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type GroupUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewGroup(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *GroupUsecase {
	return &GroupUsecase{
		lg:   lg,
		repo: repo,
	}
}

func (u *GroupUsecase) GetGroupById(ctx context.Context, id int64) (*model.Group, error)
func (u *GroupUsecase) GetGroups(ctx context.Context) ([]model.Group, error)
func (u *GroupUsecase) CreateGroup(ctx context.Context, group *model.Group) (int64, error)
func (u *GroupUsecase) UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error
func (u *GroupUsecase) PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error
func (u *GroupUsecase) DeleteGroupById(ctx context.Context, id int64) error
