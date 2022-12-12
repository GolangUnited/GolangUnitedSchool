package student

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudent(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentUseCase {
	return &StudentUseCase{lg: lg, repo: repo}
}

func (k *StudentUseCase) GetStudents(ctx context.Context) ([]model.Student, error) {
	panic("not implemented")
}
func (k *StudentUseCase) GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error) {
	panic("not implemented")
}
func (k *StudentUseCase) AddStudent(ctx context.Context, data *model.Student) (int64, error) {
	panic("not implemented")
}
func (k *StudentUseCase) UpdateStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error {
	panic("not implemented")
}
func (k *StudentUseCase) DeleteStudentByStudentId(ctx context.Context, id int64) error {
	panic("not implemented")
}
func (k *StudentUseCase) PutStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error {
	panic("not implemented")
}
