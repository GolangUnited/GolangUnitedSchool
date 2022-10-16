package postgres

import "github.com/lozovoya/GolangUnitedSchool/app/model"

func DBPersonToPerson(p *DBPerson) *model.Person {
	return &model.Person{
		ID:         p.ID,
		FirstName:  StringPointerToString(p.FirstName),
		LastName:   StringPointerToString(p.LastName),
		Patronymic: StringPointerToString(p.Patronymic),
		Login:      StringPointerToString(p.Login),
		RoleId:     p.RoleId,
		Passwd:     StringPointerToString(p.Passwd),
		UpdatedAt:  p.UpdatedAt,
	}
}

func StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
