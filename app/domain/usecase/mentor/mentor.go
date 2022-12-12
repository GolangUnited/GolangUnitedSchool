package mentor

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type MentorUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewMentor(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *MentorUseCase {
	return &MentorUseCase{lg: lg, repo: repo}
}

func (k *MentorUseCase) GetMentors(ctx context.Context) ([]model.Mentor, error) {
	panic("not implemented")
}
func (k *MentorUseCase) GetMentorById(ctx context.Context, id int64) (*model.Mentor, error) {
	panic("not implemented")
}
func (k *MentorUseCase) AddMentor(ctx context.Context, data *model.UpdateMentor) (int64, error) {
	panic("not implemented")
}
func (k *MentorUseCase) UpdateMentorByMentorId(ctx context.Context, id int64, data *model.UpdateMentor) error {
	panic("not implemented")
}
func (k *MentorUseCase) DeleteMentorByMentorId(ctx context.Context, id int64) error {
	panic("not implemented")
}
func (k *MentorUseCase) PutMentorById(ctx context.Context, id int64, data *model.UpdateMentor) error {
	panic("not implemented")
}
