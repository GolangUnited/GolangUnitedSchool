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

type GroupContactTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_GroupContactSuite(t *testing.T) {
	suite.Run(t, &GroupContactTestSuite{})
}

func (s *GroupContactTestSuite) SetupSuite() {
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
	s.Data, err = loadTestDataFromYaml("group_contact_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
}
func (s *GroupContactTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
}
func (s *GroupContactTestSuite) SetupTest() {
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
func (s *GroupContactTestSuite) TearDownTest() {
	fmt.Println("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *GroupContactTestSuite) TestGetGroupContactById() {
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
		expectedResult *model.GroupContact
	}{
		{
			name: "success",
			args: args{
				ctx: ctxB,
				id:  1,
			},
			expectedResult: &model.GroupContact{
				GroupId:        1,
				ContactValue:   "one",
				IsPrimary:      false,
				ContactTypeId:  1,
				GroupContactId: 1,
			},
		},
		{
			name: "doesn't exists",
			args: args{
				ctx: ctxB,
				id:  3,
			},
			expectedError: "couldn't get group contact by id: 3: no rows in result set",
		},
	}

	for i, tt := range tests {

		result, err := s.testRepo.GetGroupContactById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.Nil(result)
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nil(err)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *GroupContactTestSuite) TestGetGroupContacts() {
	//kek, _ := time.Parse("2006-01-02", "2006-01-02")
	tests := []struct {
		name           string
		ctx            context.Context
		isConnClosed   bool
		expectedErr    error
		expectedResult []model.GroupContact
	}{
		{

			name: "get group contacts",
			ctx:  context.Background(),
			expectedResult: []model.GroupContact{
				{1, 1, 1, false, "one"},
				{2, 2, 2, true, "two"},
			},
		},
		{
			name:         "db pool is closed",
			ctx:          context.Background(),
			isConnClosed: true,
			expectedErr:  errors.New("couldn't get list of group contacts: closed pool"),
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

			result, err := s.testRepo.GetGroupContacts(tt.ctx)
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
func (s *GroupContactTestSuite) TestAddGroupContact() {
	type args struct {
		ctx    context.Context
		course *model.GroupContactCU
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
				course: &model.GroupContactCU{
					GroupId:       4,
					ContactTypeId: 4,
					IsPrimary:     false,
					ContactValue:  "kek",
				},
			},
			ExpectedResult: 3,
		},
		{
			name: "error: duplicate title",
			args: args{
				ctx: context.Background(),
				course: &model.GroupContactCU{
					GroupId:       1,
					ContactTypeId: 1,
					IsPrimary:     false,
					ContactValue:  "one",
				},
			},
			expectedError: "couldn't create a group contact: ERROR: duplicate key value violates unique constraint \"group_contacts_contact_value_key\" (SQLSTATE 23505)",
		},
	}

	for i, tt := range tests {
		s.Run(tt.name, func() {
			id, err := s.testRepo.AddGroupContact(tt.args.ctx, tt.args.course)
			if tt.expectedError != "" {
				s.EqualErrorf(err, tt.expectedError, "case[%d]", i)
			} else {
				s.Nil(err)
				s.Equalf(tt.ExpectedResult, id, "case[%d]", i)
			}
		})
	}
}
func (s *GroupContactTestSuite) TestPutGroupContactById() {
	type args struct {
		ctx   context.Context
		id    int64
		group *model.GroupContactCU
	}
	ctxB := context.Background()
	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.GroupContact
	}{
		{
			name: "success full put",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.GroupContactCU{
					GroupId:       2,
					ContactTypeId: 2,
					IsPrimary:     false,
					ContactValue:  "kok",
				},
			},
			expectedResult: &model.GroupContact{
				GroupContactId: 1,
				GroupId:        2,
				ContactTypeId:  2,
				IsPrimary:      false,
				ContactValue:   "kok",
			},
		},
		{
			name: "nothing to update",
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.GroupContactCU{
					GroupId:       0,
					ContactTypeId: 0,
					IsPrimary:     false,
					ContactValue:  "",
				},
			},
			expectedError: "couldn't put group contact by id: 4: group contact doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				ctx: ctxB,
				id:  4,
				group: &model.GroupContactCU{
					GroupId:       0,
					ContactTypeId: 0,
					IsPrimary:     false,
					ContactValue:  "",
				},
			},
			expectedError: "couldn't put group contact by id: 4: closed pool",
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
		err := s.testRepo.PutGroupContactById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
			result, err := s.testRepo.GetGroupContactById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, result, "case [%d]", i)
		}
	}
}
func (s *GroupContactTestSuite) TestUpdateGroupContactById() {
	boolkek := true
	type args struct {
		ctx   context.Context
		id    int64
		group *model.GroupContactUpdate
	}
	ctxB := context.Background()

	tests := []struct {
		name           string
		args           args
		isConnClosed   bool
		expectedError  string
		expectedResult *model.GroupContact
	}{
		{
			name: "success full update",
			args: args{
				ctx: ctxB,
				id:  1,
				group: &model.GroupContactUpdate{
					GroupId:       utils.GetIntPtr(3),
					ContactTypeId: nil,
					IsPrimary:     &boolkek,
					ContactValue:  nil,
				},
			},
			expectedResult: &model.GroupContact{
				GroupContactId: 1,
				GroupId:        3,
				ContactTypeId:  1,
				IsPrimary:      true,
				ContactValue:   "one",
			},
		},
		{
			name: "nothing to update",
			args: args{
				id:  4,
				ctx: ctxB,
				group: &model.GroupContactUpdate{
					GroupId:       utils.GetIntPtr(2),
					ContactTypeId: utils.GetIntPtr(2),
					IsPrimary:     &boolkek,
					ContactValue:  utils.GetStrPtr("two"),
				},
			},
			expectedError: "couldn't update group contact by id 4: group contact doesn't exist",
		},
		{
			name:         "closed pool",
			isConnClosed: true,
			args: args{
				id:  2,
				ctx: ctxB,
				group: &model.GroupContactUpdate{
					GroupId:       nil,
					ContactTypeId: nil,
					IsPrimary:     nil,
					ContactValue:  nil,
				},
			},
			expectedError: "couldn't update group contact by id 2: closed pool",
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
		err := s.testRepo.UpdateGroupContactById(tt.args.ctx, tt.args.id, tt.args.group)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)

			course, err := s.testRepo.GetGroupContactById(tt.args.ctx, tt.args.id)
			s.Nilf(err, "case [%d]", i)
			s.Equalf(tt.expectedResult, course, "case [%d]", i)
		}
	}
}
func (s *GroupContactTestSuite) TestDeleteGroupContactById() {
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
			expectedError: "couldn't delete group contact by id 10: group contact doesn't exist",
		},
		{
			name: "error bad request",
			args: args{
				ctx: context.Background(),
			},
			isConnClosed:  true,
			expectedError: "couldn't delete group contact by id 0: closed pool",
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
		err := s.testRepo.DeleteGroupContactById(tt.args.ctx, tt.args.id)
		if tt.expectedError != "" {
			s.EqualErrorf(err, tt.expectedError, "case [%d]", i)
		} else {
			s.Nilf(err, "case [%d]", i)
		}
	}
}
