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
)

type MentorTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_MentorSuite(t *testing.T) {
	suite.Run(t, &MentorTestSuite{})
}

func (s *MentorTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("student_note_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *MentorTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *MentorTestSuite) SetupTest() {
	fmt.Println("setup test group contact")
	ctxB := context.Background()
	for i, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Errorf(err, "case [%d]", i)
			return
		}
	}
}
func (s *MentorTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *MentorTestSuite) TestGetMentorById() {
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
		expectedResult *model.Mentor
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  2,
			},
			expectedResult: &model.Mentor{
				MentorId: 2,
				UserId:   2,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  5,
			},
			expectedError: "couldn't get mentor by id: 5: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetMentorById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *MentorTestSuite) TestGetMentors() {
	//kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.Mentor
	}{
		{

			name: "get group contacts",
			ctx:  context.Background(),
			expectedResult: []model.Mentor{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of mentors: closed pool"),
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

			result, err := s.testRepo.GetMentors(tt.ctx)
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
func (s *MentorTestSuite) TestAddMentor() {
	type args struct {
		ctx    context.Context
		course *model.UpdateMentor
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
				course: &model.UpdateMentor{
					UserId: utils.GetIntPtr(7),
				},
			},
			ExpectedResult: 5,
		},
		{
			name: "error: duplicate user id",
			args: args{
				ctx: context.Background(),
				course: &model.UpdateMentor{
					UserId: utils.GetIntPtr(2),
				},
			},
			expectedError: "couldn't create a group contact: ERROR: duplicate key value violates unique constraint \"mentors_user_id_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddMentor(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *MentorTestSuite) TestPutMentorById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateMentor
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Mentor
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(10),
				},
			},
			expectedResult: &model.Mentor{
				MentorId: 1,
				UserId:   10,
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  10,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(10),
				},
			},
			expectedError: "couldn't put mentor by id: 10: mentor doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(10),
				},
			},
			expectedError: "couldn't put mentor by id: 4: closed pool",
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
		err := s.testRepo.PutMentorById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetMentorById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *MentorTestSuite) TestUpdateMentorById() {
	//boolkek := true
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateMentor
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Mentor
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(11),
				},
			},
			expectedResult: &model.Mentor{
				UserId:   11,
				MentorId: 1,
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  14,
				ctx: ctxB,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(11),
				},
			},
			expectedError: "couldn't update mentor by id 14: mentor doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				group: &model.UpdateMentor{
					UserId: utils.GetIntPtr(11),
				},
			},
			expectedError: "couldn't update mentor by id 2: closed pool",
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
		err := s.testRepo.UpdateMentorByMentorId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetMentorById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *MentorTestSuite) TestDeleteMentorById() {
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
				id:  4,
			},
			//expectedError: "couldn't delete mentor by id 1: ERROR: update or delete on table \"mentors\" violates foreign key constraint \"mentor_notes_mentor_id_fkey\" on table \"mentor_notes\" (SQLSTATE 23503)",
		},
		{
			name: "error nothing to delete",
			args: args{
				ctx: context.Background(),
				id:  10,
			},
			expectedError: "couldn't delete mentor by id 10: mentor doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete mentor by id 0: closed pool",
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
		err := s.testRepo.DeleteMentorByMentorId(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
