package lib_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lk153/gsheet-go/constant"
	"github.com/lk153/gsheet-go/lib"
)

type GsheetServiceTestSuite struct {
	suite.Suite
	credentialFilePath string
}

func TestGsheetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GsheetServiceTestSuite))
}

func (suite *GsheetServiceTestSuite) SetupTest() {
	suite.credentialFilePath = "client_secret.json"
	suite.T().Setenv(constant.GsheetCredential, `{
		"web": {
			"client_id": "client_id",
			"project_id": "project_id",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_secret": "client_secret",
			"redirect_uris": [
				"https://redirect_uris.com"
			]
		}
	}`)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService() {
	gsrv, err := lib.NewGsheetService(suite.credentialFilePath)
	suite.Nil(gsrv)
	suite.Require().Error(err)
	suite.Contains(err.Error(), "unable to read authorization code")
}

func (suite *GsheetServiceTestSuite) TestNewGsheetServiceV2() {
	gsrv, err := lib.NewGsheetServiceV2()
	suite.NotNil(gsrv)
	suite.Require().NoError(err)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetServiceV2_has_GSHEET_TOKEN() {
	gsrv, err := lib.NewGsheetServiceV2()
	suite.NotNil(gsrv)
	suite.Require().NoError(err)
}
