package postgres

import "github.com/lozovoya/GolangUnitedSchool/app/model"

func DBPersonToPerson(p *DBPerson) *model.Person {
	return &model.Person{
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
