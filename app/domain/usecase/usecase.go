package usecase

import "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

type CourseUsecaseInterface interface {
	// AddCourse()
	// EditCourse()
	// EdotCourseByID(id int64) error
	// DeleteCourse() error
	// DeleteCoursebyID(id int64) error
}

type ContactUsecaseInterface interface {
	GetContactsByPersonId(personId int64) (model.Contact, error)
	AddNewPersonContact(model.Contact) error
	DeletePersonContact(contactId int64) error
	UpdatePersonContact(model.Contact) error
}
