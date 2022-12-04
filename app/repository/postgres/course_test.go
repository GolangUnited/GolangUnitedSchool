package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/utils"
	"github.com/stretchr/testify/suite"
)

type CourseTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_CourseSuite(t *testing.T) {
	suite.Run(t, &CourseTestSuite{})
}

func (s *CourseTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("course_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}

func (s *CourseTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}

func (s *CourseTestSuite) SetupTest() {
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

func (s *CourseTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *CourseTestSuite) TestCreateCourse() {
	type args struct {
		ctx    context.Context
		course *model.CourseCreate
	}
	startDate, err := time.Parse(time.RFC3339, "2022-11-22T12:00:00Z")
	s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T12:00:00Z")
	s.Nil(err)

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
				course: &model.CourseCreate{
					Title:     "some title",
					StartDate: &startDate,
					EndDate:   &endDate,
				},
			},
			ExpectedResult: 2,
		},
		{
			name: "error: duplicate title",
			args: args{
				ctx: context.Background(),
				course: &model.CourseCreate{
					Title: "some title 1",
				},
			},
			expectedError: "couldn't create a course: ERROR: duplicate key value violates unique constraint \"course_title_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.CreateCourse(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}

func (s *CourseTestSuite) TestGetCourses() {
	type args struct {
		ctx        context.Context
		pagination *model.PaginationParams
	}
	ctxB := context.Background()
	startDate, err := time.Parse(time.RFC3339, "2022-10-20T00:00:00Z")
	s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T00:00:00Z")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.CourseList
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:     1,
					PageSize: 1,
				},
			},
			expectedResult: &model.CourseList{
				Courses: []model.Course{
					{
						Id:        1,
						Title:     "some title 1",
						StatusId:  1,
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
			},
		},
		{
			name: "success with metadata",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:         1,
					PageSize:     10,
					WithMetadata: true,
				},
			},
			expectedResult: &model.CourseList{
				Courses: []model.Course{
					{
						Id:        1,
						Title:     "some title 1",
						StatusId:  1,
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
				Metadata: model.PaginationResponse{
					Page:       1,
					PageSize:   10,
					PageCount:  1,
					TotalCount: 1,
				},
			},
		},
		{
			name: "pagination error page number",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:     -1,
					PageSize: 10,
				},
			},
			expectedError: ErrorBadPaginationParams,
		},
		{
			name: "pagination error page size",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:     1,
					PageSize: -1,
				},
			},
			expectedError: ErrorBadPaginationParams,
		},
		{
			name: "page out of range",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:     2,
					PageSize: 10,
				},
			},
			expectedError: "get course: out of the pages range: page doesn't exists",
		},
		{
			name: "closed conn",
			args: args{
				ctx: ctxB,
				pagination: &model.PaginationParams{
					Page:     1,
					PageSize: 1,
				},
			},
			isConnClosed:  true,
			expectedError: "get course: couldn't get total count: couldn't get total count: closed pool",
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
		result, err := s.testRepo.GetCourses(tt.args.ctx, tt.args.pagination)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}

func (s *CourseTestSuite) TestGetCourseById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	ctxB := context.Background()
	startDate, err := time.Parse(time.RFC3339, "2022-10-20T00:00:00Z")
	s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T00:00:00Z")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Course
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
			expectedResult: &model.Course{
				Id:        1,
				Title:     "some title 1",
				StatusId:  1,
				StartDate: startDate,
				EndDate:   endDate,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  3,
			},
			expectedError: "couldn't get course by id: 3: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetCourseByID(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}

func (s *CourseTestSuite) TestUpdateCourseById() {
	type args struct {
		ctx    context.Context
		id     int64
		course *model.CourseUpdate
	}
	ctxB := context.Background()
	startDate, err := time.Parse(time.RFC3339, "2022-11-22T12:00:00Z")
	s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T12:00:00Z")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Course
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				course: &model.CourseUpdate{
					Title:     utils.GetStrPtr("some title updated"),
					StatusId:  utils.GetIntPtr(2),
					StartDate: utils.GetTimePtr(startDate),
					EndDate:   utils.GetTimePtr(endDate),
				},
			},
			expectedResult: &model.Course{
				Id:        1,
				Title:     "some title updated",
				StatusId:  2,
				StartDate: startDate,
				EndDate:   endDate,
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  2,
				ctx: ctxB,
				course: &model.CourseUpdate{
					Title: utils.GetStrPtr("some updated title"),
				},
			},
			expectedError: "couldn't update course by id 2: course doesn't exists",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				course: &model.CourseUpdate{
					Title: utils.GetStrPtr("some updated title"),
				},
			},
			expectedError: "couldn't update course by id 2: closed pool",
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
		err := s.testRepo.UpdateCourseByID(tt.args.ctx, tt.args.id, tt.args.course)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetCourseByID(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}

func (s *CourseTestSuite) TestPutCourseById() {
	type args struct {
		ctx    context.Context
		id     int64
		course *model.CourseUpdate
	}
	ctxB := context.Background()
	startDate, err := time.Parse(time.RFC3339, "2022-11-22T12:00:00Z")
	s.Nil(err)
	endDate, err := time.Parse(time.RFC3339, "2022-12-25T12:00:00Z")
	s.Nil(err)
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Course
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				course: &model.CourseUpdate{
					Title:     utils.GetStrPtr("some title put"),
					StatusId:  utils.GetIntPtr(2),
					StartDate: utils.GetTimePtr(startDate),
					EndDate:   utils.GetTimePtr(endDate),
				},
			},
			expectedResult: &model.Course{
				Id:        1,
				Title:     "some title put",
				StatusId:  2,
				StartDate: startDate,
				EndDate:   endDate,
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  2,
				course: &model.CourseUpdate{
					Title:     utils.GetStrPtr("some title put"),
					StatusId:  utils.GetIntPtr(2),
					StartDate: utils.GetTimePtr(startDate),
					EndDate:   utils.GetTimePtr(endDate),
				},
			},
			expectedError: "couldn't put course by id: 2: course doesn't exists",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  2,
				course: &model.CourseUpdate{
					Title:     utils.GetStrPtr("some title put"),
					StatusId:  utils.GetIntPtr(2),
					StartDate: utils.GetTimePtr(startDate),
					EndDate:   utils.GetTimePtr(endDate),
				},
			},
			expectedError: "couldn't put course by id: 2: closed pool",
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
		err := s.testRepo.PutCourseByID(tt.args.ctx, tt.args.id, tt.args.course)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetCourseByID(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}

func (s *CourseTestSuite) TestDeleteCourseById() {
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
			expectedError: "couldn't delete course by id 10: course doesn't exists",
		},
		{
			name: "error bad request",
			args: args{
				ctx: ctxB,
			},
			isConnClosed:  true,
			expectedError: "couldn't delete course by id 0: closed pool",
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
		err := s.testRepo.DeleteCourseByID(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
