package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CourseTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_CourseSuite(t *testing.T) {
	suite.Run(t, new(CourseTestSuite))
}

func (s *CourseTestSuite) SetupSuite() {
	s.T().Log("setup all suite")
	var err error
	ctxB := context.Background()
	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
	if err != nil {
		s.FailNowf("setup failed: unable to connect to database", err.Error())
	}
	s.Data, err = loadTestDataFromYaml("course_test.yml")
	if err != nil {
		s.FailNowf("setup failed: unable to load data from yml file", err.Error())
	}
	for _, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Error(err)
			return
		}
	}
}

func (s *CourseTestSuite) TearDownSuite() {
	defer s.testRepo.pool.Close()
	s.T().Log("cleaning up database")
	var err error
	for _, r := range s.Data.Conf.Teardown.Requests {
		_, err = s.testRepo.pool.Exec(context.Background(), r.Resuest)
		if err != nil {
			s.Error(err)
		}
	}
}

func (s *CourseTestSuite) SetupTest() {
	fmt.Println("setup test")
	ctxB := context.Background()
	_, err := s.testRepo.pool.Exec(ctxB, s.Data.Conf.Setup.Requests[3].Resuest)
	if err != nil {
		s.FailNowf("couldn't insert data into course table", err.Error())
	}
}

func (s *CourseTestSuite) TearDownTest() {
	_, err := s.testRepo.pool.Exec(context.Background(), `DELETE FROM course`)
	if err != nil {
		s.FailNowf("couldn't clean table course", err.Error())
	}
}

func (s *CourseStatusTestSuite) TestCreateCourse() {

}
