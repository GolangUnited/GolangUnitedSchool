package mentor

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorNoteUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewMentorNote(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *MentorNoteUseCase {
	return &MentorNoteUseCase{lg: lg, repo: repo}
}

func (k *MentorUseCase) GetMentorNotes(ctx context.Context) ([]model.MentorNote, error) {
	panic("not implemented")
}
func (k *MentorUseCase) GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error) {
	panic("not implemented")
}
func (k *MentorUseCase) GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error) {
	panic("not implemented")
}
func (k *MentorUseCase) AddMentorNote(ctx context.Context, data *model.NewMentorNote) (int64, error) {
	panic("not implemented")
}
func (k *MentorUseCase) UpdateMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error {
	panic("not implemented")
}
func (k *MentorUseCase) DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error {
	panic("not implemented")
}
func (k *MentorUseCase) PutMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error {
	panic("not implemented")
}
