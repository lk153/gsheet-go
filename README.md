# Gsheet Utils
Go Packages for GSheet Utils

### 1. Setup ENV

    export GSHEET_CREDENTIAL='...'
    export GSHEET_TOKEN='...'

### 2. Additional Information

* The Spreadsheet ID is the last string of characters in the URL for your spreadsheet. For example, in the URL https://docs.google.com/spreadsheets/d/1qpyC0XzvTcKT6EISywvqESX3A0MwQoFDE8p-Bll4hps/edit#gid=0 , the spreadsheet ID is **``1qpyC0XzvTcKT6EISywvqESX3A0MwQoFDE8p-Bll4hps``** .


### 3. Get Gsheet Credential

[Reference](https://developers.google.com/sheets/api/quickstart/go#authorize_credentials_for_a_desktop_application)


### 4. Get Gsheet Token

1. Input URL on web browser to get Authorization Code firstly:

    https://accounts.google.com/o/oauth2/auth?state-token=offline&redirect_uri=http://localhost&response_type=code&client_id=**{Client ID}**&scope=https://www.googleapis.com/auth/spreadsheets

2. Call POST API:

    `curl --location 'https://oauth2.googleapis.com/token' \\`<br>
    `--header 'Content-Type: application/x-www-form-urlencoded' \\`<br>
    `--data-urlencode 'grant_type=authorization_code' \\`<br>
    `--data-urlencode 'code={Authorization Code}' \\`<br>
    `--data-urlencode 'client_id={Client ID}' \\`<br>
    `--data-urlencode 'client_secret={Client Secret}' \\`<br>
    `--data-urlencode 'redirect_uri=http://localhost'`
