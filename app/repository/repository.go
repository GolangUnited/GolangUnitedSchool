package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

type PersonRepositoryInterface interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
}

type CourseRepositoryInterface interface {
}

type ContactRepository interface {
	GetContactsByPersonId(personId int64) (model.Contact, error)
	AddNewPersonContact(model.Contact) error
	DeletePersonContact(contactId int64) error
	UpdatePersonContact(model.Contact) error
}
