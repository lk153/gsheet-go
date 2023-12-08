package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewGsheetService(credentialFilePath string) (gsrv *GSheetService, err error) {
	ctx := context.Background()
	b, err := os.ReadFile(credentialFilePath)
	if err != nil {
		log.Default().Println("Unable to read client secret file: ", err)
		return
	}

	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Default().Println("Unable to parse client secret file to config: ", err)
		return
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(
		getClient(ctx, config),
	))
	if err != nil {
		log.Default().Println("Unable to retrieve Sheets client: ", err)
		return
	}

	log.Default().Println("Initialized GSheet client.........")
	gsrv = &GSheetService{srv}
	return
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(ctx, tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf(`Go to the following link in your browser then type the 
	authorization code: \n%v\n`, authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (tok *oauth2.Token, err error) {
	f, err := os.Open(file)
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
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
