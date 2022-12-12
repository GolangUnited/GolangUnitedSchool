package person

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type PersonUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewPerson(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *PersonUseCase {
	return &PersonUseCase{lg: lg, repo: repo}
}

func (k *PersonUseCase) GetPersons(
	ctx context.Context) ([]model.Person, error) {
	panic("not implemented")
}
func (k *PersonUseCase) GetPersonById(
	ctx context.Context, id int64) (*model.Person, error) {
	panic("not implemented")
}
func (k *PersonUseCase) AddNewPerson(
	ctx context.Context, data *model.NewPersonDto) (int64, error) {
	panic("not implemented")
}
func (k *PersonUseCase) UpdatePersonByID(
	ctx context.Context, id int64, data *model.UpdatePerson) error {
	panic("not implemented")
}
func (k *PersonUseCase) PutPersonByID(
	ctx context.Context, id int64, data *model.UpdatePerson) error {
	panic("not implemented")
}
func (k *PersonUseCase) DeletePersonById(
	ctx context.Context, id int64) error {
	panic("not implemented")
}
