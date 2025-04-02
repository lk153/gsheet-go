package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/v2/internal/service"
	serviceMock "github.com/lk153/gsheet-go/v2/mocks/service"
	utilMock "github.com/lk153/gsheet-go/v2/util/mock"
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

func (suite *GSheetServiceTestSuite) TestReadSheet() {
	s := &service.GSheetService{}
	caller := serviceMock.NewMockV4SpreadsheetsValuesGetCall(suite.T())
	caller.On("Do").Return(&sheets.ValueRange{Values: [][]interface{}{
		{0, 1},
		{2, 3},
	}})

	valueService := serviceMock.NewMockV4SpreadsheetsValuesService(suite.T())
	valueService.On("Get", utilMock.GenMocksParams(2)...).Return(caller)

	developerMetadata := serviceMock.NewMockV4DeveloperMetadata(suite.T())
	s.DeveloperMetadata = developerMetadata
	s.Values = valueService

	result := s.ReadSheet(suite.spreadsheetID, suite.readRange)
	suite.Nil(result)
}
