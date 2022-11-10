package postgres

import (
	"context"
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
			s.Fail("setup failed")
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
	}
	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddCourseStatus(tt.args.ctx, tt.args.courseStatus)
			if tt.expectedErr == nil {
				s.Nil(err, "AddCourseStatus, case[%d]: expected nil, got %v", i, err)
				s.Equalf(tt.expectedResult, id,
					"AddCourseStatus, case[%d]: expected %v, got %v",
					i, tt.expectedResult, id)
			} else {
				s.EqualErrorf(err, tt.expectedErr.Error(),
					"AddCourseStatus, case[%d]: expected %v, got %v", i, err, tt.expectedErr)
			}
		})
	}
}
