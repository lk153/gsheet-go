package lib

import (
	"log"

	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/v2/constant"
)

type ISheetService interface {
	ReadSheet(spreadsheetID, readRange string) (values [][]string)
	Append(spreadsheetID string, range_ string, values [][]any) (resp *sheets.AppendValuesResponse, err error)
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
func (srv *GSheetService) ReadSheet(spreadsheetID, readRange string) (values [][]string) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
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
	spreadsheetID string, range_ string, values [][]any,
) (resp *sheets.AppendValuesResponse, err error) {
	valuerange := &sheets.ValueRange{
		MajorDimension: constant.MajorDimensionRows.String(),
		Range:          range_,
		Values:         values,
	}

	caller := srv.Spreadsheets.Values.Append(spreadsheetID, range_, valuerange)
	caller.ValueInputOption(constant.ValueInputOptionRaw.String())
	caller.InsertDataOption(constant.InsertDataOptionInsertRows.String())
	caller.ResponseValueRenderOption(constant.ValueRenderOptionFormattedValue.String())

	resp, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to append data to sheet: ", err)
	}

	respJSON, _ := resp.MarshalJSON()
	log.Default().Println("Appended: json result:", string(respJSON))
	return
}

func (srv *GSheetService) Update(
	spreadsheetID, updateRange string, values *sheets.ValueRange,
) (result *sheets.UpdateValuesResponse, err error) {
	caller := srv.Spreadsheets.Values.Update(spreadsheetID, updateRange, values)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}

func (srv *GSheetService) Search(
	spreadsheetID string,
) (result *sheets.SearchDeveloperMetadataResponse, err error) {
	searchdevelopermetadatarequest := &sheets.SearchDeveloperMetadataRequest{
		DataFilters: []*sheets.DataFilter{
			{
				DeveloperMetadataLookup: &sheets.DeveloperMetadataLookup{
					LocationType:     constant.LocationTypeRow.String(),
					MetadataLocation: &sheets.DeveloperMetadataLocation{},
				},
			},
		},
	}
	caller := srv.Spreadsheets.DeveloperMetadata.Search(spreadsheetID, searchdevelopermetadatarequest)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}
