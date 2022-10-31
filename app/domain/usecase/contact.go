package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type ContactUsecase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewContact(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseUseCase {
	return &CourseUseCase{lg: lg, repo: repo}
}

func (u *ContactUsecase) GetContactsByPersonId(personId int64) (model.Contact, error) {
	return model.Contact{}, nil
}
func (u *ContactUsecase) AddNewPersonContact(model.Contact) error {
	return nil
}
func (u *ContactUsecase) DeletePersonContact(contactId int64) error {
	return nil
}
func (u *ContactUsecase) UpdatePersonContact(model.Contact) error {
	return nil
}
