package postgres

import "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *repository) GetContactsByPersonId(personId int64) (model.Contact, error) {
	return model.Contact{}, nil
}
func (r *repository) AddNewPersonContact(model.Contact) error {
	return nil
}
func (r *repository) DeletePersonContact(contactId int64) error {
	return nil
}
func (r *repository) UpdatePersonContact(model.Contact) error {
	return nil
}
