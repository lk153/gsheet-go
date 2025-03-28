package lib

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
)

func NewGsheetServiceV2() (gsrv *GSheetService, err error) {
	ctx := context.Background()
	b := []byte(os.Getenv(constant.GsheetCredential))
	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Default().Println("Unable to parse client secret file to config: ", err)
		return
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(
		getClientV2(ctx, config),
	))
	if err != nil {
		log.Default().Println("Unable to retrieve Sheets client: ", err)
		return
	}

	log.Default().Println("Initialized GSheet client.........")
	gsrv = &GSheetService{srv}
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
