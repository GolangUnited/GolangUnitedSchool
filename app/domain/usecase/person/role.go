package person

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type RoleUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func (u *RoleUsecase) GetRoleById(ctx context.Context, id int64) (*model.Role, error) {
	return nil, nil
}
func (u *RoleUsecase) GetRoles(ctx context.Context) ([]model.Role, error) {
	return nil, nil
}
func (u *RoleUsecase) AddRole(ctx context.Context, role *model.RoleCU) error {
	return nil
}
func (u *RoleUsecase) UpadateRoleById(ctx context.Context, role *model.RoleCU) error {
	return nil
}
func (u *RoleUsecase) PutRoleById(ctx context.Context, role *model.Role) error {
	return nil
}
func (u *RoleUsecase) DeleteRoleById(ctx context.Context, id int64) error {
	return nil
}
