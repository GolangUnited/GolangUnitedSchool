package repository

import "context"

type Repository interface {
	GetPersonById(ctx context.Context, id int64) (*DBPerson, error)
}
