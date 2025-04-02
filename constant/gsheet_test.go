package constant_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lk153/gsheet-go/v2/constant"
)

type GsheetConstantTestSuite struct {
	suite.Suite
}

func TestGsheetConstantTestSuite(t *testing.T) {
	suite.Run(t, new(GsheetConstantTestSuite))
}

func (suite *GsheetConstantTestSuite) TestMajorDimensionRows() {
	var s constant.MajorDimension = 1
	suite.NotNil(s)
	suite.Equal(s.String(), "ROWS")
}

func (suite *GsheetConstantTestSuite) TestMajorDimensionColumns() {
	var s constant.MajorDimension = 2
	suite.NotNil(s)
	suite.Equal(s.String(), "COLUMNS")
}

func (suite *GsheetConstantTestSuite) TestMajorDimensionUnknown() {
	var s constant.MajorDimension = 3
	suite.NotNil(s)
	suite.Equal(s.String(), "unknown")
}

func (suite *GsheetConstantTestSuite) TestValueInputOptionRaw() {
	var s constant.ValueInputOption = 0
	suite.NotNil(s)
	suite.Equal(s.String(), "RAW")
}

func (suite *GsheetConstantTestSuite) TestValueInputOptionUserEntered() {
	var s constant.ValueInputOption = 1
	suite.NotNil(s)
	suite.Equal(s.String(), "USER_ENTERED")
}

func (suite *GsheetConstantTestSuite) TestValueInputOptionUnknown() {
	var s constant.ValueInputOption = 2
	suite.NotNil(s)
	suite.Equal(s.String(), "unknown")
}

func (suite *GsheetConstantTestSuite) TestLocationTypeRow() {
	var s constant.LocationType = 0
	suite.NotNil(s)
	suite.Equal(s.String(), "ROW")
}

func (suite *GsheetConstantTestSuite) TestLocationTypeColumn() {
	var s constant.LocationType = 1
	suite.NotNil(s)
	suite.Equal(s.String(), "COLUMN")
}

func (suite *GsheetConstantTestSuite) TestLocationTypeUnknown() {
	var s constant.LocationType = 2
	suite.NotNil(s)
	suite.Equal(s.String(), "unknown")
}

func (suite *GsheetConstantTestSuite) TestInsertDataOptionOverwrite() {
	var s constant.InsertDataOption = 0
	suite.NotNil(s)
	suite.Equal(s.String(), "OVERWRITE")
}

func (suite *GsheetConstantTestSuite) TestInsertDataOptionInsertRows() {
	var s constant.InsertDataOption = 1
	suite.NotNil(s)
	suite.Equal(s.String(), "INSERT_ROWS")
}

func (suite *GsheetConstantTestSuite) TestInsertDataOptionUnknown() {
	var s constant.InsertDataOption = 2
	suite.NotNil(s)
	suite.Equal(s.String(), "unknown")
}

func (suite *GsheetConstantTestSuite) TestValueRenderOptionFormattedValue() {
	var s constant.ValueRenderOption = 0
	suite.NotNil(s)
	suite.Equal(s.String(), "FORMATTED_VALUE")
}

func (suite *GsheetConstantTestSuite) TestValueRenderOptionUnFormattedValue() {
	var s constant.ValueRenderOption = 1
	suite.NotNil(s)
	suite.Equal(s.String(), "UNFORMATTED_VALUE")
}

func (suite *GsheetConstantTestSuite) TestValueRenderOptionFormula() {
	var s constant.ValueRenderOption = 2
	suite.NotNil(s)
	suite.Equal(s.String(), "FORMULA")
}

func (suite *GsheetConstantTestSuite) TestValueRenderOptionUnknown() {
	var s constant.ValueRenderOption = 3
	suite.NotNil(s)
	suite.Equal(s.String(), "unknown")
}
