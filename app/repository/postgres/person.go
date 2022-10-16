package postgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lozovoya/GolangUnitedSchool/app/domain"
)

type DBPerson struct {
	ID        int64
	FirstName *string
	LastName  *string
	SurName   *string
	Login     *string
}

func (r *PostgreSQLRepository) GetPersonById(ctx context.Context, id int64) (*domain.Person, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	personQuery := psql.Select("*").From("person").Where(sq.Eq{"person_id": id})
	sql, args, err := personQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetPersonById: %w", err)
	}

	var person DBPerson
	err = r.conn.QueryRow(ctx, sql, args...).Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.SurName)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQLRepository.GetPersonById: %w", err)
	}

	return DBPersonToPerson(&person), nil
}

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
