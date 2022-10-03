package repository

import "github.com/lozovoya/GolangUnitedSchool/app/domain"

func DBPersonToPerson(p *DBPerson) *domain.Person {
	return &domain.Person{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		SurName:   p.SurName,
	}
}
