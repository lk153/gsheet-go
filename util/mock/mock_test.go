package mock_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lk153/gsheet-go/v2/util/mock"
)

type MockTestSuite struct {
	suite.Suite
}

func TestMockTestSuite(t *testing.T) {
	suite.Run(t, new(MockTestSuite))
}

func (suite *MockTestSuite) TestSetServiceGenMocksParams() {
	params := mock.GenMocksParams(2)
	suite.Equal(2, len(params))
}
