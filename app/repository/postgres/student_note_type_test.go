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

type StudentNoteTypeTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_StudentNoteTypeSuite(t *testing.T) {
	suite.Run(t, &StudentNoteTypeTestSuite{})
}

func (s *StudentNoteTypeTestSuite) SetupSuite() {
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
func (s *StudentNoteTypeTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *StudentNoteTypeTestSuite) SetupTest() {
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
func (s *StudentNoteTypeTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *StudentNoteTypeTestSuite) TestGetStudentNoteTypeById() {
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
		expectedResult *model.StudentNoteType
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
			expectedResult: &model.StudentNoteType{
				StudentNoteTypeId: 1,
				Title:             "work",
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  4,
			},
			expectedError: "couldn't get student note type by id: 4: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetStudentNoteTypeById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTypeTestSuite) TestGetStudentNoteTypes() {
	//kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.StudentNoteType
	}{
		{

			name: "get group contacts",
			ctx:  context.Background(),
			expectedResult: []model.StudentNoteType{
				{1, "work"},
				{2, "exp"},
				{3, "self"},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of student note types: closed pool"),
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

			result, err := s.testRepo.GetStudentNoteTypes(tt.ctx)
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
func (s *StudentNoteTypeTestSuite) TestAddStudentNoteType() {
	type args struct {
		ctx    context.Context
		course *model.NewStudentNoteType
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
				course: &model.NewStudentNoteType{
					Title: "kokok",
				},
			},
			ExpectedResult: 4,
		},
		{
			name: "error: duplicate title",
			args: args{
				ctx: context.Background(),
				course: &model.NewStudentNoteType{
					Title: "work",
				},
			},
			expectedError: "couldn't create a group contact: ERROR: duplicate key value violates unique constraint \"student_note_types_title_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddStudentNoteType(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *StudentNoteTypeTestSuite) TestPutStudentNoteTypeById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateStudentNoteType
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.StudentNoteType
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("lol"),
				},
			},
			expectedResult: &model.StudentNoteType{
				StudentNoteTypeId: 1,
				Title:             "lol",
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("ke"),
				},
			},
			expectedError: "couldn't update student note type by id 4: student note type doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("lo"),
				},
			},
			expectedError: "couldn't update student note type by id 4: closed pool",
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
		err := s.testRepo.PutStudentNoteTypeById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetStudentNoteTypeById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTypeTestSuite) TestUpdateStudentNoteTypeById() {
	//boolkek := true
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateStudentNoteType
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.StudentNoteType
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("llll"),
				},
			},
			expectedResult: &model.StudentNoteType{
				StudentNoteTypeId: 1,
				Title:             "llll",
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  4,
				ctx: ctxB,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("llll"),
				},
			},
			expectedError: "couldn't update student note type by id 4: student note type doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				group: &model.UpdateStudentNoteType{
					Title: utils.GetStrPtr("llll"),
				},
			},
			expectedError: "couldn't update student note type by id 2: closed pool",
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
		err := s.testRepo.UpdateStudentNoteTypeById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetStudentNoteTypeById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *StudentNoteTypeTestSuite) TestDeleteStudentNoteTypeById() {
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
			expectedError: "couldn't delete student note type by id 4: student note type doesn't exist",
		},
		{
			name: "error nothing to delete",
			args: args{
				ctx: context.Background(),
				id:  10,
			},
			expectedError: "couldn't delete student note type by id 10: student note type doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete student note type by id 0: closed pool",
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
		err := s.testRepo.DeleteStudentNoteTypeById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
