package postgres

import "github.com/lozovoya/GolangUnitedSchool/app/domain"

func DBPersonToPerson(p *DBPerson) *domain.Person {
	return &domain.Person{
		ID:         p.ID,
		FirstName:  StringPointerToString(p.FirstName),
		LastName:   StringPointerToString(p.LastName),
		Patronymic: StringPointerToString(p.Patronymic),
		Login:      StringPointerToString(p.Login),
		RoleId:     p.RoleId,
		Passwd:     StringPointerToString(p.Passwd),
		UpdatedAt:  p.UpdatedAt,
		Deleted:    p.Deleted,
	}
}

func StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
