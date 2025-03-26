package gsheetgo

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/v2/constant"
	"github.com/lk153/gsheet-go/v2/internal/service"
)

type ClientV2 struct {
	ConfigFromJSONFunc func(jsonKey []byte, scope ...string) (*oauth2.Config, error)
}

func GetClientV2() *ClientV2 {
	return &ClientV2{
		ConfigFromJSONFunc: google.ConfigFromJSON,
	}
}

func (cli *ClientV2) NewGsheetServiceV2(opts ...option.ClientOption) (gsrv *service.GSheetService, err error) {
	ctx := context.Background()
	b := []byte(os.Getenv(constant.GsheetCredential))
	config, err := cli.ConfigFromJSONFunc(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Default().Println("Unable to parse client secret file to config: ", err)
		return
	}

	opts = append(opts, option.WithHTTPClient(
		getClientV2(ctx, config),
	))
	srv, err := sheets.NewService(ctx, opts...)
	if err != nil {
		log.Default().Println("Unable to retrieve Sheets client: ", err)
		return
	}

	log.Default().Println("Initialized GSheet client.........")
	gsrv = &service.GSheetService{}
	gsrv.SetService(srv)
	return
}

func getClientV2(ctx context.Context, config *oauth2.Config) *http.Client {
	oauth2Token, err := tokenFromEnv()
	if err != nil {
		return nil
	}

	return config.Client(ctx, oauth2Token)
}

func tokenFromEnv() (oauth2Token *oauth2.Token, err error) {
	token := os.Getenv(constant.GsheetToken)
	oauth2Token = &oauth2.Token{}
	err = json.Unmarshal([]byte(token), oauth2Token)
	return oauth2Token, err
}
