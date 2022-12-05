package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/utils"
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
	suite.Run(t, &PersonTestSuite{})
}

func (s *PersonTestSuite) SetupSuite() {
	fmt.Println("setup all suite")
	var err error
	ctxB := context.Background()
	s.testRepo.lg, err = zap.NewLogger("debug", "json")
	if err != nil {
		s.FailNowf("setup failed: unable to create new logger", err.Error())
	}
	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
	if err != nil {
		s.FailNowf("setup failed: unable to connect to database", err.Error())
	}
	s.Data, err = loadTestDataFromYaml("person_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *PersonTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *PersonTestSuite) SetupTest() {
	fmt.Println("setup test course")
	ctxB := context.Background()
	for i, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Errorf(err, "case [%d]", i)
			return
		}
	}
}
func (s *PersonTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

//func (s *PersonTestSuite) SetupTest() {
//	fmt.Println("setup test")
//	var err error
//	ctxB := context.Background()
//	s.testRepo.lg, err = zap.NewLogger("debug", "json")
//	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
//	if err != nil {
//		s.Error(err)
//		s.Fail("setup failed")
//		return
//	}
//	s.Data, err = loadTestDataFromYaml("person_test.yml")
//	if err != nil {
//		s.Error(err)
//		s.Fail("setup failed")
//		return
//	}
//	for _, r := range s.Data.Conf.Setup.Requests {
//		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
//		if err != nil {
//			s.Error(err)
//			return
//		}
//	}
//}
//func (s *PersonTestSuite) TearDownTest() {
//	fmt.Println("cleaning up")
//	var err error
//	for _, r := range s.Data.Conf.Teardown.Requests {
//		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
//		if err != nil {
//			s.Error(err)
//			s.Fail("cleaning failed", "requet", r.Resuest)
//		}
//	}
//}

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
			expectedErr:  errors.New("couldn't get list of persons: closed pool"),
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
func (s *PersonTestSuite) TestAddNewPerson() {
	type args struct {
		ctx    context.Context
		course *model.NewPersonDto
	}

	tests := []struct {
		name           string
		args           args
		expectedError  string
		ExpectedResult int64
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				course: &model.NewPersonDto{
					FirstName:  "lkj",
					LastName:   "lkj",
					Patronymic: "lkj",
					Login:      "098",
					Email:      "oiu",
					Birthday:   "oiu",
					RoleId:     2,
					Passwd:     "oiu",
				},
			},
			ExpectedResult: 4,
		},
		{
			name: "error: duplicate login",
			args: args{
				ctx: context.Background(),
				course: &model.NewPersonDto{
					FirstName:  "lkj",
					LastName:   "lkj",
					Patronymic: "lkj",
					Login:      "3",
					Email:      "lkj",
					Birthday:   "lkj",
					RoleId:     7,
					Passwd:     "lkj",
				},
			},
			expectedError: "couldn't add new person: ERROR: duplicate key value violates unique constraint \"person_login_key\" (SQLSTATE 23505)",
		},
		{
			name: "error: no required field",
			args: args{
				ctx: context.Background(),
				course: &model.NewPersonDto{
					FirstName:  "",
					LastName:   "lkj",
					Patronymic: "lkj",
					Login:      "kjhlkjh",
					Email:      "lkj",
					Birthday:   "lkj",
					RoleId:     0,
					Passwd:     "lkj",
				},
			},
			expectedError: "couldn't add new person: ERROR: null value in column \"first_name\" of relation \"person\" violates not-null constraint (SQLSTATE 23502)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddNewPerson(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *PersonTestSuite) TestUpdatePersonById() {
	type args struct {
		ctx    context.Context
		id     int64
		person *model.UpdatePerson
	}
	ctxB := context.Background()
	//startDate, err := time.Parse(time.RFC3339, "2022-11-22T12:00:00Z")
	//s.Nil(err)
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T12:00:00Z")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Person
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				person: &model.UpdatePerson{
					FirstName:  utils.GetStrPtr("11"),
					LastName:   utils.GetStrPtr("11"),
					Patronymic: utils.GetStrPtr("11"),
					Login:      utils.GetStrPtr("11"),
					Email:      utils.GetStrPtr("11"),
					Birthday:   utils.GetStrPtr("11"),
					RoleId:     utils.GetIntPtr(11),
					Passwd:     utils.GetStrPtr("11"),
					UpdatedAt:  utils.GetTimePtr(endDate),
				},
			},

			expectedResult: &model.Person{
				PersonId:   1,
				FirstName:  "11",
				LastName:   "11",
				Patronymic: "11",
				Login:      "11",
				Email:      "11",
				Birthday:   "11",
				RoleId:     11,
				Passwd:     "11",
				UpdatedAt:  endDate,
				Deleted:    false,
				CreatedAt:  kek,
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  6,
				ctx: ctxB,
				person: &model.UpdatePerson{
					FirstName:  utils.GetStrPtr("11"),
					LastName:   utils.GetStrPtr("11"),
					Patronymic: utils.GetStrPtr("11"),
					Login:      utils.GetStrPtr("11"),
					Email:      utils.GetStrPtr("11"),
					Birthday:   utils.GetStrPtr("11"),
					RoleId:     utils.GetIntPtr(11),
					Passwd:     utils.GetStrPtr("11"),
					UpdatedAt:  utils.GetTimePtr(endDate),
				},
			},
			expectedError: "couldn't update person by id 6: person doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				person: &model.UpdatePerson{
					FirstName:  nil,
					LastName:   nil,
					Patronymic: nil,
					Login:      nil,
					Email:      nil,
					Birthday:   nil,
					RoleId:     nil,
					Passwd:     nil,
					UpdatedAt:  nil,
					Deleted:    true,
				},
			},
			expectedError: "couldn't update person by id 2: closed pool",
		},
	}
	for i, tt := range tests {
		if tt.isConnClosed {
			s.testRepo.pool.Close()
			defer func() {
				// reconnect
				s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
			}()
		}
		err := s.testRepo.UpdatePersonByID(tt.args.ctx, tt.args.id, tt.args.person)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			person, err := s.testRepo.GetPersonById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, person, "case [%d]", i)
		}
	}
}
func (s *PersonTestSuite) TestPutPersonById() {
	type args struct {
		ctx    context.Context
		id     int64
		person *model.UpdatePerson
	}
	ctxB := context.Background()
	//startDate, err := time.Parse(time.RFC3339, "2022-11-22T12:00:00Z")
	//s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T12:00:00Z")
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Person
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				person: &model.UpdatePerson{
					FirstName:  utils.GetStrPtr("11"),
					LastName:   utils.GetStrPtr("11"),
					Patronymic: utils.GetStrPtr("11"),
					Login:      utils.GetStrPtr("11"),
					Email:      utils.GetStrPtr("11"),
					Birthday:   utils.GetStrPtr("11"),
					RoleId:     utils.GetIntPtr(11),
					Passwd:     utils.GetStrPtr("11"),
					UpdatedAt:  utils.GetTimePtr(endDate),
					Deleted:    true,
				},
			},
			expectedResult: &model.Person{
				PersonId:   1,
				FirstName:  "11",
				LastName:   "11",
				Patronymic: "11",
				Login:      "11",
				Email:      "11",
				Birthday:   "11",
				RoleId:     11,
				Passwd:     "11",
				UpdatedAt:  endDate,
				CreatedAt:  kek,
				Deleted:    true,
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  5,
				person: &model.UpdatePerson{
					FirstName:  utils.GetStrPtr("11"),
					LastName:   utils.GetStrPtr("11"),
					Patronymic: utils.GetStrPtr("11"),
					Login:      utils.GetStrPtr("11"),
					Email:      utils.GetStrPtr("11"),
					Birthday:   utils.GetStrPtr("11"),
					RoleId:     utils.GetIntPtr(11),
					Passwd:     utils.GetStrPtr("11"),
					UpdatedAt:  utils.GetTimePtr(endDate),
					Deleted:    true,
				},
			},
			expectedError: "couldn't put person by id: 5: person doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  2,
				person: &model.UpdatePerson{
					FirstName:  utils.GetStrPtr("11"),
					LastName:   utils.GetStrPtr("11"),
					Patronymic: utils.GetStrPtr("11"),
					Login:      utils.GetStrPtr("11"),
					Email:      utils.GetStrPtr("11"),
					Birthday:   utils.GetStrPtr("11"),
					RoleId:     utils.GetIntPtr(11),
					Passwd:     utils.GetStrPtr("11"),
					UpdatedAt:  utils.GetTimePtr(endDate),
					Deleted:    true,
				},
			},
			expectedError: "couldn't put person by id: 2: closed pool",
		},
	}
	for i, tt := range tests {
		if tt.isConnClosed {
			s.testRepo.pool.Close()
			defer func() {
				// reconnect
				s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
			}()
		}
		err := s.testRepo.PutPersonByID(tt.args.ctx, tt.args.id, tt.args.person)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetPersonById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *PersonTestSuite) TestDeletePersonById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	ctxB := context.Background()
	tests := []struct {
		name          string
		args          args
		isConnClosed  bool
		expectedError string
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
		},
		{
			name: "error nothing to delete",
			args: args{
				ctx: ctxB,
				id:  10,
			},
			expectedError: "couldn't delete person by id 10: person doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: ctxB,
			},
			isConnClosed:  true,
			expectedError: "couldn't delete person by id 0: closed pool",
		},
	}
	for i, tt := range tests {
		if tt.isConnClosed {
			s.testRepo.pool.Close()
			defer func() {
				// reconnect
				s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
			}()
		}
		err := s.testRepo.DeletePersonById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
