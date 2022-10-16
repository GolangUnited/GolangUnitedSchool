package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgreSQLRepository struct {
	conn *pgx.Conn
}

func NewPostgreSQLRepository(
	ctx context.Context,
	connectionString string,
) (*PostgreSQLRepository, error) {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLRepository{
		conn: conn,
	}, nil
}
