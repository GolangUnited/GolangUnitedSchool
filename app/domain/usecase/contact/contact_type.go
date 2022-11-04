package contact

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type ContactTypeUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewContactType(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *ContactTypeUsecase {
	return &ContactTypeUsecase{lg: lg, repo: repo}
}

func (u *ContactTypeUsecase) GetContactTypes(
	ctx context.Context) ([]model.ContactType, error) {
	panic("not implemented")
}
func (u *ContactTypeUsecase) GetContactTypeById(
	ctx context.Context, id int64) (*model.ContactType, error) {
	panic("not implemented")
}
func (u *ContactTypeUsecase) AddContactType(
	ctx context.Context) error {
	panic("not implemented")
}
func (u *ContactTypeUsecase) UpdateContactType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
func (u *ContactTypeUsecase) DeleteContactType(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
