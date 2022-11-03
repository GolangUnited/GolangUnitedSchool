package course

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type ProjectUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewProject(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *ProjectUsecase {
	return &ProjectUsecase{lg: lg, repo: repo}
}

func (u *ProjectUsecase) GetProjects(
	ctx context.Context) ([]model.Project, error) {
	panic("not implemented")
}

func (u *ProjectUsecase) GetProjectById(
	ctx context.Context,
	id int64) (*model.Project, error) {
	panic("not implemented")
}

func (u *ProjectUsecase) AddProject(
	ctx context.Context,
	data *model.Project) error {
	panic("not implemented")
}

func (u *ProjectUsecase) UpdateProject(
	ctx context.Context,
	id int64,
	data *model.Project) error {
	panic("not implemented")
}

func (u *ProjectUsecase) DeleteProject(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
