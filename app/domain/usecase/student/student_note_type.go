package student

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentNoteTypeUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewStudentNoteType(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *StudentNoteTypeUseCase {
	return &StudentNoteTypeUseCase{lg: lg, repo: repo}
}

func (k *StudentNoteTypeUseCase) GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error) {
	panic("not implemented")
}
func (k *StudentNoteTypeUseCase) GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error) {
	panic("not implemented")
}
func (k *StudentNoteTypeUseCase) AddStudentNoteType(ctx context.Context, data *model.NewStudentNoteType) (int64, error) {
	panic("not implemented")
}
func (k *StudentNoteTypeUseCase) UpdateStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error {
	panic("not implemented")
}
func (k *StudentNoteTypeUseCase) PutStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error {
	panic("not implemented")
}
func (k *StudentNoteTypeUseCase) DeleteStudentNoteTypeById(ctx context.Context, id int64) error {
	panic("not implemented")
}
