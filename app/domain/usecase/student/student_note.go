package student

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentNoteUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentNote(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentNoteUseCase {
	return &StudentNoteUseCase{lg: lg, repo: repo}
}

func (k *StudentNoteUseCase) GetStudentNotes(ctx context.Context) ([]model.StudentNote, error) {
	panic("not implemented")
}
func (k *StudentNoteUseCase) GetStudentNoteById(ctx context.Context, id int64) (*model.StudentNote, error) {
	panic("not implemented")
}
func (k *StudentNoteUseCase) AddStudentNote(ctx context.Context, data *model.NewStudentNote) (int64, error) {
	panic("not implemented")
}
func (k *StudentNoteUseCase) UpdateStudentNoteByStudentId(ctx context.Context, id int64, data *model.UpdateStudentNote) error {
	panic("not implemented")
}
func (k *StudentNoteUseCase) PutStudentNoteById(ctx context.Context, id int64, data *model.UpdateStudentNote) error {
	panic("not implemented")
}
func (k *StudentNoteUseCase) DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error {
	panic("not implemented")
}
func (k *StudentNoteUseCase) GetStudentsNotesByStudentId(ctx context.Context, id int64) ([]model.StudentNote, error) {
	panic("not implemented")
}
