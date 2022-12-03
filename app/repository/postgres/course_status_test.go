package postgres

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/stretchr/testify/suite"
)

const (
	testDSN = "postgres://pguser:pguser@localhost:5432/pgdbtest"
)

type CourseStatusTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_CourseSatatusSuite(t *testing.T) {
	suite.Run(t, new(CourseStatusTestSuite))
}

func (s *CourseStatusTestSuite) SetupTest() {
	fmt.Println("setup test")
	var err error
	ctxB := context.Background()
	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
	if err != nil {
		s.Error(err)
		s.Fail("setup failed")
		return
	}
	s.Data, err = loadTestDataFromYaml("course_status_test.yaml")
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

func (s *CourseStatusTestSuite) TearDownTest() {
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

func (s *CourseStatusTestSuite) TestAddCourseStatus() {
	type args struct {
		ctx          context.Context
		courseStatus *model.CourseStatus
	}
	tests := []struct {
		name           string
		args           args
		expectedResult int64
		expectedErr    error
	}{
		{
			name: "create new course status",
			args: args{
				ctx:          context.Background(),
				courseStatus: &model.CourseStatus{Title: "some title"},
			},
			expectedResult: 5,
		},
		{
			name: "create with existing title",
			args: args{
				ctx:          context.Background(),
				courseStatus: &model.CourseStatus{Title: "some title"},
			},
			expectedErr: errors.New("couldn't add course status: ERROR: duplicate key value violates unique constraint \"course_status_title_key\" (SQLSTATE 23505)"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddCourseStatus(tt.args.ctx, tt.args.courseStatus)
			if tt.expectedErr == nil {
				s.Nil(err, "case[%d]", i)
				s.Equalf(tt.expectedResult, id, "case[%d]", i)
			} else {
				s.EqualErrorf(err, tt.expectedErr.Error(),
					"case[%d]", i)
			}
		})
	}
}

func (s *CourseStatusTestSuite) TestGetCourseStatusByID() {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name           string
		args           args
		expectedErr    error
		expectedResult *model.CourseStatus
	}{
		{
			name: "get course status success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			expectedResult: &model.CourseStatus{
				Id:    1,
				Title: "active",
			},
		},
		{
			name: "course status doesn't exists",
			args: args{
				ctx: context.Background(),
				id:  7,
			},
			expectedErr: errors.New("couldn't get course status: no rows in result set"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			result, err := s.testRepo.GetCourseStatusById(tt.args.ctx, tt.args.id)
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

func (s *CourseStatusTestSuite) TestGetCourseStatuses() {
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.CourseStatus
	}{
		{
			name: "get course statuses",
			ctx:  context.Background(),
			expectedResult: []model.CourseStatus{
				{Id: 1, Title: "active"},
				{Id: 2, Title: "done"},
				{Id: 3, Title: "inactive"},
				{Id: 4, Title: "created"},
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

			result, err := s.testRepo.GetCourseStatuses(tt.ctx)
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

func (s *CourseStatusTestSuite) TestDeleteCourseStatusById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name         string
		isConnClosed bool
		args         args
		expectedErr  error
	}{
		{
			name: "delete course status by id success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
		{
			name:         "db pool is closed",
			isConnClosed: true,
			args: args{
				ctx: context.Background(),
				id:  7,
			},
			expectedErr: errors.New("couldn't delete course status: closed pool"),
		},
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			if tt.isConnClosed {
				s.testRepo.pool.Close()
				defer func() {
					// reconnect
					s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
				}()
			}

			err := s.testRepo.DeleteCourseStatusById(tt.args.ctx, tt.args.id)
			if tt.expectedErr == nil {
				s.Nilf(err, "case[%d]", i)
			} else {
				s.EqualErrorf(err, tt.expectedErr.Error(), "case[%d]", i)
			}
		})
	}
}

func (s *CourseStatusTestSuite) TestDeleteCourseStatusByID() {
	type args struct {
		ctx  context.Context
		id   int64
		data *model.CourseStatus
	}
	tests := []struct {
		name         string
		isConnClosed bool
		args         args
		expectedErr  error
	}{
		{
			name: "course status updated successfully",
			args: args{
				ctx:  context.Background(),
				id:   2,
				data: &model.CourseStatus{Id: 2, Title: "done status [updated]"},
			},
		},
		{
			name:         "db pool is closed",
			isConnClosed: true,
			args: args{
				ctx:  context.Background(),
				data: &model.CourseStatus{},
			},
			expectedErr: errors.New("couldn't update course status: closed pool"),
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

		err := s.testRepo.UpdateCourseStatusById(tt.args.ctx, tt.args.id, tt.args.data)
		if tt.expectedErr == nil {
			s.Nilf(err, "case[%d]", i)
			res, err := s.testRepo.GetCourseStatusById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case[%d]", i)
			s.Equalf(tt.args.data, res, "case[%d]", i)
		} else {
			s.EqualErrorf(err, tt.expectedErr.Error(), "case[%d]", i)
		}
	}
}
