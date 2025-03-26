[![Go Reference](https://pkg.go.dev/badge/lk153/gsheet-go/markdown.svg)](https://pkg.go.dev/github.com/lk153/gsheet-go/v2)
[![MultiPlatformUnitTest](https://github.com/lk153/gsheet-go/actions/workflows/unit_test.yml/badge.svg)](https://github.com/lk153/gsheet-go/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/lk153/gsheet-go/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/lk153/gsheet-go/actions/workflows/reviewdog.yml)
[![Gosec](https://github.com/lk153/gsheet-go/actions/workflows/gosec.yml/badge.svg)](https://github.com/lk153/gsheet-go/actions/workflows/gosec.yml)

[![Coverage](https://raw.githubusercontent.com/lk153/octocovs-central/main/badges/lk153/gsheet-go/coverage.svg)](https://github.com/lk153/octocovs-central)
[![Coverage Status](https://coveralls.io/repos/github/lk153/gsheet-go/badge.svg?branch=main)](https://coveralls.io/github/lk153/gsheet-go?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/lk153/gsheet-go/v2)](https://goreportcard.com/report/github.com/lk153/gsheet-go/v2)

# Gsheet Utils
Go Packages support for GSheet Integration

## Installation

#### 1. Setup ENV

    export GSHEET_CREDENTIAL='...'
    export GSHEET_TOKEN='...'

#### 2. Additional Information

* The Spreadsheet ID is the last string of characters in the URL for your spreadsheet. For example, in the URL https://docs.google.com/spreadsheets/d/1qpyC0XzvTcKT6EISywvqESX3A0MwQoFDE8p-Bll4hps/edit#gid=0 , the spreadsheet ID is **``1qpyC0XzvTcKT6EISywvqESX3A0MwQoFDE8p-Bll4hps``** .


#### 3. Get Gsheet Credential

[Reference](https://developers.google.com/sheets/api/quickstart/go#authorize_credentials_for_a_desktop_application)


#### 4. Get Gsheet Token

1. Input URL with replacement of param `client_id` on web browser to get **Authorization Code**:

    https://accounts.google.com/o/oauth2/auth?state-token=offline&redirect_uri=http://localhost&response_type=code&client_id={client_id}&scope=https://www.googleapis.com/auth/spreadsheets


	https://accounts.google.com/o/oauth2/auth?state-token=offline&redirect_uri=http://localhost&response_type=code&client_id=600779649356-8uu1iknp22orljn6rgu3oumn34reh8om.apps.googleusercontent.com&scope=https://www.googleapis.com/auth/spreadsheets


2. Make API call (POST method) to retrieve **access token**:

    `curl --location 'https://oauth2.googleapis.com/token' \\`<br>
    `--header 'Content-Type: application/x-www-form-urlencoded' \\`<br>
    `--data-urlencode 'grant_type=authorization_code' \\`<br>
    `--data-urlencode 'code={Authorization Code}' \\`<br>
    `--data-urlencode 'client_id={Client ID}' \\`<br>
    `--data-urlencode 'client_secret={Client Secret}' \\`<br>
    `--data-urlencode 'redirect_uri=http://localhost'`

## Quickstart:

```go
package main

import (
	"fmt"

	"github.com/lk153/gsheet-go/v2/lib"
)

func main() {
	runTest()
}

func runTest() {
	srv, err := lib.NewGsheetServiceV2()
	if err != nil {
		fmt.Println("Cannot connect Gsheet!")
		return
	}

	spreadsheetID := "01c-onQeYHmvc-EPkrJDU-WyAydbCAA1ng6hXCgdYiqqg"
	readRange := "'To Update on DB'!A3:AR3"
	values := srv.ReadSheet(spreadsheetID, readRange)
	for idx, row := range values {
            ...
            ...
	}
}
```
