package group

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentGroupUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentGroup(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentGroupUseCase {
	return &StudentGroupUseCase{lg: lg, repo: repo}
}

func (k *StudentGroupUseCase) GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error) {
	panic("not implemented")
}
func (k *StudentGroupUseCase) GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error) {
	panic("not implemented")
}
func (k *StudentGroupUseCase) CreateStudentGroup(ctx context.Context, data *model.StudentGroup) (int64, error) {
	panic("not implemented")
}
func (k *StudentGroupUseCase) PutStudentGroupById(ctx context.Context, id int64, data *model.UpdateStudentGroup) error {
	panic("not implemented")
}
func (k *StudentGroupUseCase) DeleteStudentGroupById(ctx context.Context, id int64) error {
	panic("not implemented")
}
