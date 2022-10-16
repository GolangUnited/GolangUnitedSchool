package postgres

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/lozovoya/GolangUnitedSchool/app/model"
)

type PostgreSQLRepository struct {
	conn *pgx.Conn
}

func (r *PostgreSQLRepository) GetPersonById(ctx context.Context, id int64) (*model.Person, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	personQuery := psql.Select("*").From("person").Where(sq.Eq{"id": id})
	sql, args, err := personQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetPersonById: %w", err)
	}
	var person DBPerson
	err = r.conn.QueryRow(ctx, sql, args...).Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.Patronymic,
		&person.Login,
		&person.RoleId,
		&person.Passwd,
		&person.UpdatedAt,
		&person.Deleted,
	)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQLRepository.GetPersonById: %w", err)
	}
	return DBPersonToPerson(&person), nil
}

func (r *PostgreSQLRepository) UpdatePerson(ctx context.Context, id int64, person *model.Person) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Update("person").SetMap(sq.Eq{
		"first_name": person.FirstName,
		"last_name":  person.LastName,
		"patronymic": person.Patronymic,
		"role_id":    person.RoleId,
		"passwd":     person.Passwd,
		"updated_at": time.Now().UTC(),
	}).Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("UpdatePerson: %w", err)
	}

	_, err = r.conn.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgreSQLRepository.UpdatePerson: %w", err)
	}

	return nil
}

func NewPostgreSQLRepository(ctx context.Context,
	connectionString string) *PostgreSQLRepository {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	return &PostgreSQLRepository{
		conn: conn,
	}
}
