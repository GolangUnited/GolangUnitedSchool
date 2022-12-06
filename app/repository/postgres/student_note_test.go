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

type StudentNoteTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_StudentNoteSuite(t *testing.T) {
	suite.Run(t, &StudentNoteTestSuite{})
}

func (s *StudentNoteTestSuite) SetupSuite() {
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
func (s *StudentNoteTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *StudentNoteTestSuite) SetupTest() {
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
func (s *StudentNoteTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *StudentNoteTestSuite) TestGetStudentNotes() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.StudentNote
	}{
		{

			name: "get student notes",
			ctx:  context.Background(),
			expectedResult: []model.StudentNote{
				{1, 1, 1, "kek", kek},
				{2, 2, 2, "kok", kek},
				{3, 3, 3, "kak", kek},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of student notes: closed pool"),
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

			result, err := s.testRepo.GetStudentNotes(tt.ctx)
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
func (s *StudentNoteTestSuite) TestGetStudentNoteById() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
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
		expectedResult *model.StudentNote
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  3,
			},
			expectedResult: &model.StudentNote{
				Id:                3,
				StudentId:         3,
				StudentNoteTypeId: 3,
				Note:              "kak",
				CreatedAt:         kek,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  4,
			},
			expectedError: "couldn't get student note by id: 4: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetStudentNoteById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTestSuite) TestAddStudentNote() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx    context.Context
		course *model.NewStudentNote
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
				course: &model.NewStudentNote{
					StudentId:         3,
					StudentNoteTypeId: 3,
					Note:              "koko",
					CreatedAt:         kek,
				},
			},
			ExpectedResult: 4,
		},
		//{
		//	name: "error: duplicate title",
		//	args: args{
		//		ctx: context.Background(),
		//		course: &model.NewStudentNote{
		//			StudentId:         0,
		//			StudentNoteTypeId: 0,
		//			Note:              "",
		//			CreatedAt:         time.Time{},
		//		},
		//	},
		//	expectedError: "couldn't create a group contact: ERROR: duplicate key value violates unique constraint \"group_contacts_contact_value_key\" (SQLSTATE 23505)",
		//},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddStudentNote(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *StudentNoteTestSuite) TestUpdateStudentNoteById() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	//boolkek := true
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateStudentNote
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.StudentNote
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateStudentNote{
					StudentId:         utils.GetIntPtr(2),
					StudentNoteTypeId: utils.GetIntPtr(2),
					Note:              utils.GetStrPtr("kekeke"),
					CreatedAt:         &kek,
				},
			},
			expectedResult: &model.StudentNote{
				Id:                1,
				StudentId:         2,
				StudentNoteTypeId: 2,
				Note:              "kekeke",
				CreatedAt:         kek,
			},
		},
		//{
		//	name: "nothing to update",
		//	args: args{
		//		id:  4,
		//		ctx: ctxB,
		//		group: &model.UpdateStudentNote{
		//			StudentId:         nil,
		//			StudentNoteTypeId: nil,
		//			Note:              nil,
		//			CreatedAt:         nil,
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
		//		group: &model.UpdateStudentNote{
		//			StudentId:         nil,
		//			StudentNoteTypeId: nil,
		//			Note:              nil,
		//			CreatedAt:         nil,
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
		err := s.testRepo.UpdateStudentNoteByStudentId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetStudentNoteById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTestSuite) TestPutStudentNoteById() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateStudentNote
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.StudentNote
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateStudentNote{
					StudentId:         utils.GetIntPtr(2),
					StudentNoteTypeId: utils.GetIntPtr(2),
					Note:              utils.GetStrPtr("kiki"),
					CreatedAt:         &kek,
				},
			},
			expectedResult: &model.StudentNote{
				Id:                1,
				StudentId:         2,
				StudentNoteTypeId: 2,
				Note:              "kiki",
				CreatedAt:         kek,
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateStudentNote{
					StudentId:         nil,
					StudentNoteTypeId: nil,
					Note:              nil,
					CreatedAt:         nil,
				},
			},
			expectedError: "couldn't put student note by id: 4: student note doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateStudentNote{
					StudentId:         nil,
					StudentNoteTypeId: nil,
					Note:              nil,
					CreatedAt:         nil,
				},
			},
			expectedError: "couldn't put student note by id: 4: closed pool",
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
		err := s.testRepo.PutStudentNoteById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetStudentNoteById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTestSuite) TestDeleteStudentNoteById() {
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
			expectedError: "couldn't delete student note by id 10: student note doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete student note by id 0: closed pool",
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
		err := s.testRepo.DeleteStudentNoteByStudentNoteId(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
