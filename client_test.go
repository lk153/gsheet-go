package gsheetgo_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"

	gsheetgo "github.com/lk153/gsheet-go/v2"
)

type GsheetServiceTestSuite struct {
	suite.Suite
	credentialFilePath string
}

func TestGsheetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GsheetServiceTestSuite))
}

func (suite *GsheetServiceTestSuite) TestGetClient() {
	cli := gsheetgo.GetClient()
	suite.NotNil(cli)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService() {
	cli := &gsheetgo.Client{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, nil
		},
		GetOauth2ClientFunc: func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetService(suite.credentialFilePath)
	suite.NotNil(gsrv)
	suite.Require().NoError(err)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService_ConfigFromJSONFunc_ERR() {
	cli := &gsheetgo.Client{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, errors.New("ConfigFromJSONFunc has error")
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, nil
		},
		GetOauth2ClientFunc: func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetService(suite.credentialFilePath)
	suite.Nil(gsrv)
	suite.Require().Error(err)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService_ReadFileFunc_ERR() {
	cli := &gsheetgo.Client{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, errors.New("ReadFileFunc has error")
		},
		GetOauth2ClientFunc: func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetService(suite.credentialFilePath)
	suite.Nil(gsrv)
	suite.Require().Error(err)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService_GetOauth2ClientFunc_ERR() {
	cli := &gsheetgo.Client{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, nil
		},
		GetOauth2ClientFunc: func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
			return nil, errors.New("GetOauth2ClientFunc has error")
		},
	}
	gsrv, err := cli.NewGsheetService(suite.credentialFilePath)
	suite.Nil(gsrv)
	suite.Require().Error(err)
}

func (suite *GsheetServiceTestSuite) TestNewGsheetService_SheetsNewService_ERR() {
	cli := &gsheetgo.Client{
		ConfigFromJSONFunc: func(jsonKey []byte, scope ...string) (*oauth2.Config, error) {
			return nil, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, nil
		},
		GetOauth2ClientFunc: func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
			return nil, nil
		},
	}
	gsrv, err := cli.NewGsheetService(suite.credentialFilePath, option.WithAPIKey("apikey"), option.WithoutAuthentication())
	suite.Nil(gsrv)
	suite.Require().Error(err)
}
