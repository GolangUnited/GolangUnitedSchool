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

type StudentTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_StudentTestSuite(t *testing.T) {
	suite.Run(t, &StudentTestSuite{})
}

func (s *StudentTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("student_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *StudentTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *StudentTestSuite) SetupTest() {
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
func (s *StudentTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *StudentTestSuite) TestDeleteStudentByStudentId() {
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
			expectedError: "couldn't delete student by id 10: student doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete student by id 0: closed pool",
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
		err := s.testRepo.DeleteStudentByStudentId(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
func (s *StudentTestSuite) TestGetStudents() {
	//kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.Student
	}{
		{

			name: "get group contacts",
			ctx:  context.Background(),
			expectedResult: []model.Student{
				{1, 1},
				{2, 2},
				{3, 3},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of students: closed pool"),
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

			result, err := s.testRepo.GetStudents(tt.ctx)
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
func (s *StudentTestSuite) TestAddStudent() {
	type args struct {
		ctx    context.Context
		course *model.Student
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
				course: &model.Student{
					StudentId: 4,
				},
			},
			ExpectedResult: 4,
		},
		{
			name: "error: duplicate title",
			args: args{
				ctx: context.Background(),
				course: &model.Student{
					StudentId: 1,
				},
			},
			expectedError: "couldn't add new student: ERROR: duplicate key value violates unique constraint \"students_student_id_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddStudent(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *StudentTestSuite) TestGetStudentByStudentId() {
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
		expectedResult *model.Student
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  3,
			},
			expectedResult: &model.Student{
				Id:        3,
				StudentId: 3,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  6,
			},
			expectedError: "couldn't get student by id: 6: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetStudentByStudentId(tt.args.ctx, tt.args.id)
		fmt.Println(result)
		fmt.Println(s.testRepo.GetStudentByStudentId(ctxB, 2))
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentTestSuite) TestUpdateStudentById() {
	//boolkek := true
	type args struct {
		ctx   context.Context
		id    int64
		group *model.StudentUpdate
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Student
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.StudentUpdate{
					StudentId: utils.GetIntPtr(4),
				},
			},
			expectedResult: &model.Student{
				Id:        1,
				StudentId: 4,
			},
		},
		//{
		//	name: "nothing to update",
		//	args: args{
		//		id:  4,
		//		ctx: ctxB,
		//		group: &model.GroupContactUpdate{
		//			GroupId:       utils.GetIntPtr(2),
		//			ContactTypeId: utils.GetIntPtr(2),
		//			IsPrimary:     &boolkek,
		//			ContactValue:  utils.GetStrPtr("two"),
		//		},
		//	},
		//	expectedError: "couldn't update group contact by id 4: group contact doesn't exist",
		//},
		//{
		//	name:         "closed pool",
		//	isConnClosed: true,
		//	args: args{
		//		id:  2,
		//		ctx: ctxB,
		//		group: &model.GroupContactUpdate{
		//			GroupId:       nil,
		//			ContactTypeId: nil,
		//			IsPrimary:     nil,
		//			ContactValue:  nil,
		//		},
		//	},
		//	expectedError: "couldn't update group contact by id 2: closed pool",
		//},
	}
	for i, tt := range tests {
		if tt.isConnClosed {
			s.testRepo.pool.Close()
			defer func() {
				// reconnect
				s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
			}()
		}
		err := s.testRepo.UpdateStudentByStudentId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetStudentByStudentId(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *StudentTestSuite) TestPutStudentById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.StudentUpdate
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.Student
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.StudentUpdate{
					StudentId: utils.GetIntPtr(10),
				},
			},
			expectedResult: &model.Student{
				Id:        1,
				StudentId: 10,
			},
		},
		//{
		//	name: "nothing to update",
		//	args: args{
		//		ctx: ctxB,
		//		id:  4,
		//		group: &model.GroupContactCU{
		//			GroupId:       0,
		//			ContactTypeId: 0,
		//			IsPrimary:     false,
		//			ContactValue:  "",
		//		},
		//	},
		//	expectedError: "couldn't put group contact by id: 4: group contact doesn't exist",
		//},
		//{
		//	name:         "closed pool",
		//	isConnClosed: true,
		//	args: args{
		//		ctx: ctxB,
		//		id:  4,
		//		group: &model.GroupContactCU{
		//			GroupId:       0,
		//			ContactTypeId: 0,
		//			IsPrimary:     false,
		//			ContactValue:  "",
		//		},
		//	},
		//	expectedError: "couldn't put group contact by id: 4: closed pool",
		//},
	}
	for i, tt := range tests {
		if tt.isConnClosed {
			s.testRepo.pool.Close()
			defer func() {
				// reconnect
				s.testRepo.pool, _ = NewDbPool(tt.args.ctx, testDSN)
			}()
		}
		err := s.testRepo.PutStudentByStudentId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetStudentByStudentId(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
