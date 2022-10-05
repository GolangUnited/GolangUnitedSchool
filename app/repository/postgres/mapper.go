package postgres

import "github.com/lozovoya/GolangUnitedSchool/app/domain"

func DBPersonToPerson(p *DBPerson) *domain.Person {
	return &domain.Person{
		ID:        p.ID,
		FirstName: StringPointerToString(p.FirstName),
		LastName:  StringPointerToString(p.LastName),
		SurName:   StringPointerToString(p.SurName),
		Login:     StringPointerToString(p.Login),
	}
}

func StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
