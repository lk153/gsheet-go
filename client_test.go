package gsheetgo_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"

	gsheetgo "github.com/lk153/gsheet-go/v2"
)

type GsheetServiceTestSuite struct {
	suite.Suite
	credentialFilePath string
}

func TestGsheetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GsheetServiceTestSuite))
}

func (suite *GsheetServiceTestSuite) SetupTest() {
	// suite.credentialFilePath = "client_secret.json"
	// suite.T().Setenv(constant.GsheetCredential, `{
	// 	"installed": {
	// 		"client_id": "600779649356-8uu1iknp22orljn6rgu3oumn34reh8om.apps.googleusercontent.com",
	// 		"project_id": "gsheet-454308",
	// 		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
	// 		"token_uri": "https://oauth2.googleapis.com/token",
	// 		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	// 		"client_secret": "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",
	// 		"redirect_uris": [
	// 		"http://localhost"
	// 		]
	// 	}
	// }`)
	// suite.T().Setenv(constant.GsheetToken, `{
	// 	"access_token": "ya29.a0AeXRPp7QFiElmmsMklWEQx3PPJR8_-Kmp99GtQW_MnZMDBq3d3WlqOz8NKsVlphaVEpq4XumnudlpnH1h-4ygOKDKlz7y8jTtX0WQyYhDf2mwQPQm1MpBBGYPqDGST0Lln4OLmKGieibmMB7vCe6J-PkaH5csyuuAm-QNRWCaCgYKAX0SARISFQHGX2MiDV96ZQjutINICZBfY7vkPg0175",
	// 	"expires_in": 3599,
	// 	"refresh_token": "1//0gJ_lmDXanmtCCgYIARAAGBASNwF-L9IrlaP8knUO7mJ5Sp6oreVzUVsNVatsMhR08-detfVOzbXYhw1yA57iH-bCqUcQiHFyZY0",
	// 	"scope": "https://www.googleapis.com/auth/spreadsheets",
	// 	"token_type": "Bearer",
	// 	"refresh_token_expires_in": 604799
	// }`)
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
