package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PersonTestSuite struct {
	suite.Suite
	testRepo PostgresRepository
	Data     TestData
}

func Test_PersonSuite(t *testing.T) {
	suite.Run(t, new(PersonTestSuite))
}

func (s *PersonTestSuite) SetupTest() {
	fmt.Println("setup test")
	var err error
	ctxB := context.Background()
	s.testRepo.pool, err = NewDbPool(ctxB, testDSN)
	if err != nil {
		s.Error(err)
		s.Fail("setup failed")
		return
	}
	s.Data, err = loadTestDataFromYaml("person_test.yml")
	if err != nil {
		s.Error(err)
		s.Fail("setup failed")
		return
	}
	for _, r := range s.Data.Conf.Setup.Requests {
		_, err := s.testRepo.pool.Exec(ctxB, r.Resuest)
		if err != nil {
			s.Error(err)
			return
		}
	}
}

func (s *PersonTestSuite) TearDownTest() {
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

func (s *PersonTestSuite) TestAddNewPerson() {

}
