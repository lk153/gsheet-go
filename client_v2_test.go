package gsheetgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"

	gsheetgo "github.com/lk153/gsheet-go/v2"
)

type GsheetServiceV2TestSuite struct {
	suite.Suite
	credentialFilePath string
}

func TestGsheetServiceV2TestSuite(t *testing.T) {
	suite.Run(t, new(GsheetServiceV2TestSuite))
}

func (suite *GsheetServiceV2TestSuite) TestGetClientV2() {
	cli := gsheetgo.GetClientV2()
	suite.NotNil(cli)
}

func (suite *GsheetServiceV2TestSuite) TestNewGsheetServiceV2() {
	cli := &gsheetgo.ClientV2{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetServiceV2()
	suite.NotNil(gsrv)
	suite.Require().NoError(err)
}

func (suite *GsheetServiceV2TestSuite) TestNewGsheetServiceV2_SheetsNewService_ERR() {
	cli := &gsheetgo.ClientV2{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetServiceV2(option.WithAPIKey("apikey"), option.WithoutAuthentication())
	suite.Nil(gsrv)
	suite.Require().Error(err)
}

func (suite *GsheetServiceV2TestSuite) TestNewGsheetServiceV2_ConfigFromJSONFunc_ERR() {
	cli := &gsheetgo.ClientV2{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, errors.New("ConfigFromJSONFunc has error")
		},
	}
	gsrv, err := cli.NewGsheetServiceV2()
	suite.Nil(gsrv)
	suite.Require().Error(err)
}
