package repository

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

type RepositoryInterface interface {
	GetContactsByPersonId(personId int64) (model.Contact, error)
	AddNewPersonContact(model.Contact) error
	DeletePersonContact(contactId int64) error
	UpdatePersonContact(model.Contact) error
}
