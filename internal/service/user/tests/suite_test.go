package tests

import (
	"testing"

	repoMocks "github.com/Arturyus92/auth/internal/repository/mocks"
	dbMocks "github.com/Arturyus92/platform_common/pkg/db/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	userRepositoryMock *repoMocks.UserRepositoryMock
	logRepositoryMock  *repoMocks.LogRepositoryMock
	transactorMock     *dbMocks.TransactorMock
	mc                 *minimock.Controller
}

func (s *TestSuite) SetupTest() {
	mc := minimock.NewController(s.T())
	s.mc = mc
	s.userRepositoryMock = repoMocks.NewUserRepositoryMock(mc)
	s.logRepositoryMock = repoMocks.NewLogRepositoryMock(mc)
	s.transactorMock = dbMocks.NewTransactorMock(mc)
}

func TestApp(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
