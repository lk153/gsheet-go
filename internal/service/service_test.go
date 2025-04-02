package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lk153/gsheet-go/v2/internal/service"
)

type GSheetServiceTestSuite struct {
	suite.Suite
	spreadsheetID string
	readRange     string
}

func (suite *GSheetServiceTestSuite) SetupTest() {
	suite.spreadsheetID = ""
	suite.readRange = ""
}

func TestGSheetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GSheetServiceTestSuite))
}

func (suite *GSheetServiceTestSuite) TestSetService() {
	s := &service.GSheetService{}
	s.SetService(nil)
	suite.NotNil(s)
}
