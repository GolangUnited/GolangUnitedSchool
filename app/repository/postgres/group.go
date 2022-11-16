package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetGroupById(ctx context.Context, id int64) (*model.Group, error)

func (r *PostgresRepository) GetGroups(ctx context.Context) ([]model.Group, error)

func (r *PostgresRepository) CreateGroup(ctx context.Context, group *model.Group) (int64, error)

func (r *PostgresRepository) UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error

func (r *PostgresRepository) PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error

func (r *PostgresRepository) DeleteGroupById(ctx context.Context, id int64) error
