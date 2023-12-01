package lib

import (
	"log"

	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/constant"
)

type ISheetService interface {
	ReadSheet()
}

type GSheetService struct {
	*sheets.Service
}

/*
  - Example:
    spreadsheetId := "1AGHH6abcXzBmfC5e9r50t3wXKhlUs5XIE-fj1U4fV0Q"
    readRange := "Log1!A2:B"
    *
*/
func (srv *GSheetService) ReadSheet(spreadsheetId, readRange string) (values [][]string) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Default().Println("Unable to retrieve data from sheet: ", err)
		return
	}

	if len(resp.Values) == 0 {
		log.Default().Println("No data found.")
	} else {
		for _, row := range resp.Values {
			colsStr := []string{}
			for _, col := range row {
				colsStr = append(colsStr, col.(string))
			}
			values = append(values, colsStr)
		}
	}

	return
}

/*
API: POST https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}/values/{range}:append
DOC: https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/append
*/
func (srv *GSheetService) Append(
	spreadsheetId string, range_ string, data []string,
) (resp *sheets.AppendValuesResponse, err error) {
	values := [][]interface{}{}
	var interfaceSlice []interface{}
	for _, s := range data {
		interfaceSlice = append(interfaceSlice, s)
	}
	values = append(values, interfaceSlice)
	valuerange := &sheets.ValueRange{
		MajorDimension: constant.MajorDimension_ROWS.String(),
		Range:          range_,
		Values:         values,
	}

	caller := srv.Spreadsheets.Values.Append(spreadsheetId, range_, valuerange)
	caller.ValueInputOption(constant.ValueInputOption_RAW.String())
	caller.InsertDataOption(constant.InsertDataOption_INSERT_ROWS.String())
	caller.ResponseValueRenderOption(constant.ValueRenderOption_FORMATTED_VALUE.String())

	resp, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to append data to sheet: ", err)
	}

	respJson, _ := resp.MarshalJSON()
	log.Default().Println("Appended: json result:", string(respJson))
	return
}

func (srv *GSheetService) Update(
	spreadsheetId, updateRange string, values *sheets.ValueRange,
) (result *sheets.UpdateValuesResponse, err error) {
	caller := srv.Spreadsheets.Values.Update(spreadsheetId, updateRange, values)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}

func (srv *GSheetService) Search(
	spreadsheetId string,
) (result *sheets.SearchDeveloperMetadataResponse, err error) {
	searchdevelopermetadatarequest := &sheets.SearchDeveloperMetadataRequest{
		DataFilters: []*sheets.DataFilter{
			{
				DeveloperMetadataLookup: &sheets.DeveloperMetadataLookup{
					LocationType:     constant.LocationType_ROW.String(),
					MetadataLocation: &sheets.DeveloperMetadataLocation{},
				},
			},
		},
	}
	caller := srv.Spreadsheets.DeveloperMetadata.Search(spreadsheetId, searchdevelopermetadatarequest)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}
