package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type PersonTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_PersonSuite(t *testing.T) {
	suite.Run(t, new(PersonTestSuite))
}

func (s *PersonTestSuite) SetupTest() {
	fmt.Println("setup test")
	var err error
	ctxB := context.Background()
	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
	if err != nil {
		s.Error(err)
		s.Fail("setup failed")
		return
	}
	s.Data, err = loadTestDataFromYaml("person_test.yml")
	if err != nil {
		s.Error(err)
		s.Fail("setup failed")
		return
	}
	for _, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Error(err)
			return
		}
	}
}

func (s *PersonTestSuite) TearDownTest() {
	fmt.Println("cleaning up")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
			s.Fail("cleaning failed", "requet", r.Resuest)
		}
	}
}

func (s *PersonTestSuite) TestPostgresRepository_GetPersons() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	fmt.Println(kek)
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.Person
	}{
		{

			name: "get course statuses",
			ctx:  context.Background(),
			expectedResult: []model.Person{
				{Passwd: "1", RoleId: 1, PersonId: 1, FirstName: "1", LastName: "1", Patronymic: "1", Login: "1", UpdatedAt: kek, CreatedAt: kek},
				{Passwd: "2", RoleId: 2, PersonId: 2, FirstName: "2", LastName: "2", Patronymic: "2", Login: "2", UpdatedAt: kek, CreatedAt: kek},
				{Passwd: "3", RoleId: 3, PersonId: 3, FirstName: "3", LastName: "3", Patronymic: "3", Login: "3", UpdatedAt: kek, CreatedAt: kek},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get course statuses: closed pool"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			if tt.isConnClosed {
				s.testRepo.pool.Close()
				defer func() {
					// reconnect
					s.testRepo.pool, _ = NewDbPool(tt.ctx, testDSN)
				}()
			}

			result, err := s.testRepo.GetPersons(tt.ctx)
			if tt.expectedErr == nil {
				s.Nil(err, "case [%d]", i)
				s.Equal(result, tt.expectedResult, "case [%d]", i)
			} else {
				s.Nil(result, "case [%d]", i)
				s.EqualError(err, tt.expectedErr.Error(), "case [%d]", i)
			}
		})
	}
}

func (s *PersonTestSuite) TestGetPersonById() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name           string
		args           args
		expectedErr    error
		expectedResult *model.Person
	}{
		{
			name: "get person by id  success",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			expectedResult: &model.Person{Passwd: "2", RoleId: 2, PersonId: 2, FirstName: "2", LastName: "2", Patronymic: "2", Login: "2", UpdatedAt: kek, CreatedAt: kek},
		},
		{
			name: "person with such id doesn't exists",
			args: args{
				ctx: context.Background(),
				id:  7,
			},
			expectedErr: errors.New("couldn't get person by id: no rows in result set"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			result, err := s.testRepo.GetPersonById(tt.args.ctx, tt.args.id)
			if tt.expectedErr == nil {
				s.Nil(err, "case[%d]", i)
				s.Equal(tt.expectedResult, result, "case[%d]", i)
			} else {
				s.EqualError(err, tt.expectedErr.Error(), "case[%d]", i)
				s.Nil(result, "case[%d]", i)
			}
		})
	}
}
