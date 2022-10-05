package postgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/lozovoya/GolangUnitedSchool/app/domain"
)

type PostgreSQLRepository struct {
	conn   *pgx.Conn
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

func NewPostgreSQLRepository(ctx context.Context,
	connectionString string) *PostgreSQLRepository {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	return &PostgreSQLRepository{
		conn:   conn,
	}
}
