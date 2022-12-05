package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GroupTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_GroupSuite(t *testing.T) {
	suite.Run(t, &GroupTestSuite{})
}

func (s *GroupTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("group_test.yaml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *GroupTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *GroupTestSuite) SetupTest() {
	fmt.Println("setup test group")
	ctxB := context.Background()
	for i, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Errorf(err, "case [%d]", i)
			return
		}
	}
}
func (s *GroupTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *GroupTestSuite) TestCreateGroup() {
	type args struct {
		ctx    context.Context
		course *model.GroupCU
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
				course: &model.GroupCU{
					MentorId: 2,
					Title:    "kekeke",
				},
			},
			ExpectedResult: 3,
		},
		{
			name: "error: duplicate title",
			args: args{
				ctx: context.Background(),
				course: &model.GroupCU{
					MentorId: 2,
					Title:    "kek",
				},
			},
			expectedError: "couldn't create a group: ERROR: duplicate key value violates unique constraint \"groups_title_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.CreateGroup(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *GroupTestSuite) TestGetGroupById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Group
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
			expectedResult: &model.Group{
				ID:       1,
				CourseId: 1,
				MentorId: 1,
				Title:    "kek",
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  3,
			},
			expectedError: "couldn't get group by id: 3: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetGroupById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *GroupTestSuite) TestGetGroups() {
	//kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.Group
	}{
		{

			name: "get course statuses",
			ctx:  context.Background(),
			expectedResult: []model.Group{
				{1, 1, 1, "kek"},
				{2, 2, 2, "kok"},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of groups: closed pool"),
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

			result, err := s.testRepo.GetGroups(tt.ctx)
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
func (s *GroupTestSuite) TestDeleteCourseById() {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name          string
		args          args
		isConnClosed  bool
		expectedError string
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
		{
			name: "error nothing to delete",
			args: args{
				ctx: context.Background(),
				id:  10,
			},
			expectedError: "couldn't delete group by id 10: group doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete group by id 0: closed pool",
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
		err := s.testRepo.DeleteGroupById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
func (s *GroupTestSuite) TestUpdateGroupById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.GroupCU
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Group
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.GroupCU{
					MentorId: 1,
					Title:    "kkk",
				},
			},
			expectedResult: &model.Group{
				ID:       1,
				CourseId: 1,
				MentorId: 1,
				Title:    "kkk",
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  4,
				ctx: ctxB,
				group: &model.GroupCU{
					MentorId: 1,
					Title:    "kkk",
				},
			},
			expectedError: "couldn't update group by id 4: group doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				group: &model.GroupCU{
					Title: "kakaka",
				},
			},
			expectedError: "couldn't update group by id 2: closed pool",
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
		err := s.testRepo.UpdateGroupById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetGroupById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *GroupTestSuite) TestPutGroupById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.GroupCU
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Group
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.GroupCU{
					MentorId: 2,
					Title:    "kkk",
				},
			},
			expectedResult: &model.Group{
				CourseId: 1,
				ID:       1,
				MentorId: 2,
				Title:    "kkk",
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.GroupCU{
					MentorId: 0,
					Title:    "",
				},
			},
			expectedError: "couldn't put group by id: 4: group doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.GroupCU{
					MentorId: 0,
					Title:    "",
				},
			},
			expectedError: "couldn't put group by id: 4: closed pool",
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
		err := s.testRepo.PutGroupById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetGroupById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
