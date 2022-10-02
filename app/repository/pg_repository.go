package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type PostgreSQLRepository struct {
	conn   *pgx.Conn
	logger *zap.SugaredLogger
}

func (r *PostgreSQLRepository) GetPersonById(ctx context.Context, id int64) (*DBPerson, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	personQuery := psql.Select("*").From("person").Where(sq.Eq{"person_id": id})
	sql, args, err := personQuery.ToSql()
	if err != nil {
		return nil, err
	}
	var person DBPerson
	err = r.conn.QueryRow(ctx, sql, args...).Scan(&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.SurName)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func NewPostgreSQLRepository(ctx context.Context,
	connectionString string,
	logger *zap.SugaredLogger) *PostgreSQLRepository {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	return &PostgreSQLRepository{
		conn:   conn,
		logger: logger,
	}
}
