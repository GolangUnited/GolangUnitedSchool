package group

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type GroupContactUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewGroupContact(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *GroupContactUseCase {
	return &GroupContactUseCase{lg: lg, repo: repo}
}

func (k *GroupContactUseCase) GetAllGroupContacts(ctx context.Context) ([]model.GroupContact, error) {
	panic("not implemented")
}
func (k *GroupContactUseCase) GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error) {
	panic("not implemented")
}
func (k *GroupContactUseCase) AddGroupContact(ctx context.Context, data *model.GroupContactCU) (int64, error) {
	panic("not implemented")
}
func (k *GroupContactUseCase) PutGroupContactById(ctx context.Context, id int64, data *model.GroupContactCU) error {
	panic("not implemented")
}
func (k *GroupContactUseCase) UpdateGroupContactById(ctx context.Context, id int64, data *model.GroupContactUpdate) error {
	panic("not implemented")
}
func (k *GroupContactUseCase) DeleteGroupContactById(ctx context.Context, id int64) error {
	panic("not implemented")
}
func (k *GroupContactUseCase) GetGroupContacts(ctx context.Context, id int64) ([]model.GroupContact, error) {
	panic("not implemented")
}
