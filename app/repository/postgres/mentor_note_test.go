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

type MentorNoteTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_MentorNoteSuite(t *testing.T) {
	suite.Run(t, &MentorNoteTestSuite{})
}

func (s *MentorNoteTestSuite) SetupSuite() {
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
func (s *MentorNoteTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *MentorNoteTestSuite) SetupTest() {
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
func (s *MentorNoteTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *MentorNoteTestSuite) TestAddMentorNote() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx    context.Context
		course *model.NewMentorNote
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
				course: &model.NewMentorNote{
					StudentId: 1,
					MentorId:  1,
					Note:      "kukuku",
					CreatedAt: kek,
				},
			},
			ExpectedResult: 4,
		},
		{
			name: "error: duplicate note",
			args: args{
				ctx: context.Background(),
				course: &model.NewMentorNote{
					StudentId: 1,
					MentorId:  1,
					Note:      "kekeke",
					CreatedAt: time.Time{},
				},
			},
			expectedError: "couldn't create a mentor note: ERROR: duplicate key value violates unique constraint \"mentor_notes_note_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddMentorNote(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *MentorNoteTestSuite) TestGetMentorNotes() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.MentorNote
	}{
		{

			name: "get mentor notes",
			ctx:  context.Background(),
			expectedResult: []model.MentorNote{
				{1, 1, 2, "kekeke", kek},
				{2, 2, 2, "kokoko", kek},
				{3, 3, 3, "kakaka", kek},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of mentor notes: closed pool"),
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

			result, err := s.testRepo.GetMentorNotes(tt.ctx)
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
func (s *MentorNoteTestSuite) TestUpdateMentorNoteById() {
	//boolkek := true
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateMentorNote
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.MentorNote
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("kekw"),
					CreatedAt: &kek,
				},
			},
			expectedResult: &model.MentorNote{
				MentorNoteId: 1,
				StudentId:    2,
				MentorId:     2,
				Note:         "kekw",
				CreatedAt:    kek,
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  4,
				ctx: ctxB,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("kekw"),
					CreatedAt: &kek,
				},
			},
			expectedError: "couldn't update mentor note by id 4: mentor note doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("kekw"),
					CreatedAt: &kek,
				},
			},
			expectedError: "couldn't update mentor note by id 2: closed pool",
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
		err := s.testRepo.UpdateMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *MentorNoteTestSuite) TestGetMentorNoteById() {
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
		expectedResult *model.MentorNote
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
			expectedResult: &model.MentorNote{
				MentorNoteId: 1,
				StudentId:    1,
				MentorId:     2,
				Note:         "kekeke",
				CreatedAt:    kek,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  4,
			},
			expectedError: "couldn't get mentor note by id: 4: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *MentorNoteTestSuite) TestDeleteMentorNoteById() {
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
			expectedError: "couldn't delete mentor note by id 10: mentor note doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete mentor note by id 0: closed pool",
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
		err := s.testRepo.DeleteMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
func (s *MentorNoteTestSuite) TestPutMentorNoteById() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	type args struct {
		ctx   context.Context
		id    int64
		group *model.UpdateMentorNote
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.MentorNote
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("klkl"),
					CreatedAt: &kek,
				},
			},
			expectedResult: &model.MentorNote{
				MentorNoteId: 1,
				StudentId:    2,
				MentorId:     2,
				Note:         "klkl",
				CreatedAt:    kek,
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("klkl"),
					CreatedAt: &kek,
				},
			},
			expectedError: "couldn't put mentor note by id: 4: mentor note doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.UpdateMentorNote{
					StudentId: utils.GetIntPtr(2),
					MentorId:  utils.GetIntPtr(2),
					Note:      utils.GetStrPtr("klkl"),
					CreatedAt: &kek,
				},
			},
			expectedError: "couldn't put mentor note by id: 4: closed pool",
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
		err := s.testRepo.PutMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetMentorNoteByMentorNoteId(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *MentorNoteTestSuite) TestGetMentorNotesByMentorId() {
	kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		mentorId       int64
		expectedErr    error
		expectedResult []model.MentorNote
	}{
		{

			name:     "get group contacts",
			ctx:      context.Background(),
			mentorId: 2,
			expectedResult: []model.MentorNote{
				{1, 1, 2, "kekeke", kek},
				{2, 2, 2, "kokoko", kek},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of mentor's notes: closed pool"),
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

			result, err := s.testRepo.GetMentorNotesByMentorId(tt.ctx, tt.mentorId)
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
