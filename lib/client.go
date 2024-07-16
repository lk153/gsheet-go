package lib

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
)

func NewGsheetService(credentialFilePath string) (gsrv *GSheetService, err error) {
	ctx := context.Background()
	b, err := os.ReadFile(filepath.Clean(credentialFilePath))
	if err != nil {
		log.Default().Println("Unable to read client secret file: ", err)
		return
	}

	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Default().Println("Unable to parse client secret file to config: ", err)
		return
	}

	client, err := getClient(ctx, config)
	if err != nil {
		log.Default().Println("Unable to init client: ", err)
		return
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(
		client,
	))
	if err != nil {
		log.Default().Println("Unable to retrieve Sheets client: ", err)
		return
	}

	log.Default().Println("Initialized GSheet client.........")
	gsrv = &GSheetService{srv}
	return
}

func getClient(ctx context.Context, config *oauth2.Config) (client *http.Client, err error) {
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
		err = fmt.Errorf("unable to read authorization code: %v", err)
		return
	}

	tok, err = config.Exchange(context.TODO(), authCode)
	if err != nil {
		err = fmt.Errorf("unable to retrieve token from web: %v", err)
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
	f, err := os.OpenFile(filepath.Clean(path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to encode oauth token: %v", err)
	}
}
