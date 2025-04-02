package gsheetgo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/v2/internal/service"
)

const (
	FileMode = os.FileMode(0600)
)

type Client struct {
	ConfigFromJSONFunc  func(jsonKey []byte, scope ...string) (*oauth2.Config, error)
	ReadFileFunc        func(name string) ([]byte, error)
	GetOauth2ClientFunc func(ctx context.Context, config *oauth2.Config) (client *http.Client, err error)
}

func GetClient() *Client {
	return &Client{
		ConfigFromJSONFunc:  google.ConfigFromJSON,
		ReadFileFunc:        os.ReadFile,
		GetOauth2ClientFunc: getOauth2HttpClient,
	}
}

func (cli *Client) NewGsheetService(credentialFilePath string, opts ...option.ClientOption) (gsrv *service.GSheetService, err error) {
	ctx := context.Background()
	b, err := cli.ReadFileFunc(filepath.Clean(credentialFilePath))
	if err != nil {
		log.Default().Println("Unable to read client secret file: ", err)
		return
	}

	config, err := cli.ConfigFromJSONFunc(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Default().Println("Unable to parse client secret file to config: ", err)
		return
	}

	client, err := cli.GetOauth2ClientFunc(ctx, config)
	if err != nil {
		log.Default().Println("Unable to init client: ", err)
		return
	}

	opts = append(opts, option.WithHTTPClient(
		client,
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

func getOauth2HttpClient(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return
		}

		saveToken(tokenFile, tok)
	}

	client = config.Client(ctx, tok)
	return
}

func getTokenFromWeb(config *oauth2.Config) (tok *oauth2.Token, err error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Default().Printf(`Go to the following link in your browser then type the authorization code: \n%v\n`, authURL)

	var authCode string
	if _, err = fmt.Scan(&authCode); err != nil {
		err = fmt.Errorf("unable to read authorization code: %w", err)
		return
	}

	tok, err = config.Exchange(context.TODO(), authCode)
	if err != nil {
		err = fmt.Errorf("unable to retrieve token from web: %w", err)
	}

	return
}

func tokenFromFile(file string) (tok *oauth2.Token, err error) {
	f, err := os.Open(filepath.Clean(file))
	if err != nil {
		return
	}
	defer f.Close()

	tok = &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(filepath.Clean(path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, FileMode)
	if err != nil {
		log.Panicf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Panicf("Unable to encode oauth token: %v", err)
	}
}
