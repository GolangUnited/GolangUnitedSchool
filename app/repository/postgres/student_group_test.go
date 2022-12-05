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

type StudentGroupTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_StudentGroupSuite(t *testing.T) {
	suite.Run(t, &StudentGroupTestSuite{})
}
func (s *StudentGroupTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("student_group_test.yaml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *StudentGroupTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *StudentGroupTestSuite) SetupTest() {
	fmt.Println("setup test for student_group")
	ctxB := context.Background()
	for i, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Errorf(err, "case [%d]", i)
			return
		}
	}
}
func (s *StudentGroupTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *StudentGroupTestSuite) TestGetStudentGroups() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	fmt.Println(kek)
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.StudentGroup
	}{
		{

			name: "get student groups",
			ctx:  context.Background(),
			expectedResult: []model.StudentGroup{
				{1, 1, 1},
				{2, 2, 2},
				{3, 1, 2},
				{4, 2, 1},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get student groups: closed pool"),
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

			result, err := s.testRepo.GetStudentGroups(tt.ctx)
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
func (s *StudentGroupTestSuite) TestGetStudentGroupById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name           string
		args           args
		expectedErr    error
		expectedResult *model.StudentGroup
	}{
		{
			name: "get course status success",
			args: args{
				ctx: context.Background(),
				id:  4,
			},
			expectedResult: &model.StudentGroup{
				Id:        4,
				StudentId: 2,
				GroupId:   1,
			},
		},
		{
			name: "course status doesn't exists",
			args: args{
				ctx: context.Background(),
				id:  7,
			},
			expectedErr: errors.New("couldn't get student group by id: no rows in result set"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			result, err := s.testRepo.GetStudentGroupById(tt.args.ctx, tt.args.id)
			fmt.Println(result)
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
func (s *StudentGroupTestSuite) TestCreateStudentGroup() {
	type args struct {
		ctx    context.Context
		course *model.StudentGroup
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
				course: &model.StudentGroup{
					StudentId: 3,
					GroupId:   2,
				},
			},
			ExpectedResult: 5,
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.CreateStudentGroup(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *StudentGroupTestSuite) TestPutStudentGroupById() {
	type args struct {
		ctx    context.Context
		id     int64
		course *model.UpdateStudentGroup
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.StudentGroup
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				course: &model.UpdateStudentGroup{
					StudentId: utils.GetIntPtr(1),
					GroupId:   utils.GetIntPtr(2),
				},
			},
			expectedResult: &model.StudentGroup{
				Id:        1,
				StudentId: 1,
				GroupId:   2,
			},
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
		err := s.testRepo.PutStudentGroupById(tt.args.ctx, tt.args.id, tt.args.course)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetStudentGroupById(tt.args.ctx, tt.args.id)
			fmt.Print(s.testRepo.GetStudentGroupById(ctxB, 4))
			fmt.Print(s.testRepo.GetStudentGroups(ctxB))
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentGroupTestSuite) TestDeleteStudentGroupById() {
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
			expectedError: "couldn't delete student group by id 10: student group doesn't not exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: ctxB,
			},
			isConnClosed:  true,
			expectedError: "couldn't delete student group by id 0: closed pool",
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
		err := s.testRepo.DeleteStudentGroupById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
