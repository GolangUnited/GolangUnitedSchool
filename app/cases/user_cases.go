package cases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"go.uber.org/zap"
)

type UserCases struct {
	r      repository.Repository
	logger *zap.SugaredLogger
}

func (c *UserCases) GetPersonById(ctx context.Context, id int64) (*domain.Person, error) {
	dbRow, err := c.r.GetPersonById(ctx, id)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return repository.DBPersonToPerson(dbRow), nil
}

func NewUserCases(
	r repository.Repository,
	logger *zap.SugaredLogger) *UserCases {
	return &UserCases{
		r:      r,
		logger: logger,
	}
}
