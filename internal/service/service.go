package service

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"

	"github.com/lk153/gsheet-go/v2/constant"
)

type ISheetService interface {
	ReadSheet(spreadsheetID, readRange string) (values [][]string)
	Append(spreadsheetID string, range_ string, values [][]any) (resp *sheets.AppendValuesResponse, err error)
}

var _ V4SpreadsheetsValuesService = &sheets.SpreadsheetsValuesService{}

type V4SpreadsheetsValuesService interface {
	Append(spreadsheetId string, range_ string, valuerange *sheets.ValueRange) *sheets.SpreadsheetsValuesAppendCall
	BatchClear(spreadsheetId string, batchclearvaluesrequest *sheets.BatchClearValuesRequest) *sheets.SpreadsheetsValuesBatchClearCall
	BatchClearByDataFilter(spreadsheetId string, batchclearvaluesbydatafilterrequest *sheets.BatchClearValuesByDataFilterRequest) *sheets.SpreadsheetsValuesBatchClearByDataFilterCall
	BatchGet(spreadsheetId string) *sheets.SpreadsheetsValuesBatchGetCall
	BatchGetByDataFilter(spreadsheetId string, batchgetvaluesbydatafilterrequest *sheets.BatchGetValuesByDataFilterRequest) *sheets.SpreadsheetsValuesBatchGetByDataFilterCall
	BatchUpdate(spreadsheetId string, batchupdatevaluesrequest *sheets.BatchUpdateValuesRequest) *sheets.SpreadsheetsValuesBatchUpdateCall
	BatchUpdateByDataFilter(spreadsheetId string, batchupdatevaluesbydatafilterrequest *sheets.BatchUpdateValuesByDataFilterRequest) *sheets.SpreadsheetsValuesBatchUpdateByDataFilterCall
	Clear(spreadsheetId string, range_ string, clearvaluesrequest *sheets.ClearValuesRequest) *sheets.SpreadsheetsValuesClearCall
	Get(spreadsheetId string, range_ string) *sheets.SpreadsheetsValuesGetCall
	Update(spreadsheetId string, range_ string, valuerange *sheets.ValueRange) *sheets.SpreadsheetsValuesUpdateCall
}

var _ V4DeveloperMetadata = &sheets.SpreadsheetsDeveloperMetadataService{}

type V4DeveloperMetadata interface {
	Get(spreadsheetId string, metadataId int64) *sheets.SpreadsheetsDeveloperMetadataGetCall
	Search(spreadsheetId string, searchdevelopermetadatarequest *sheets.SearchDeveloperMetadataRequest) *sheets.SpreadsheetsDeveloperMetadataSearchCall
}

var _ V4SpreadsheetsValuesGetCall = &sheets.SpreadsheetsValuesGetCall{}

type V4SpreadsheetsValuesGetCall interface {
	Context(ctx context.Context) *sheets.SpreadsheetsValuesGetCall
	DateTimeRenderOption(dateTimeRenderOption string) *sheets.SpreadsheetsValuesGetCall
	Do(opts ...googleapi.CallOption) (*sheets.ValueRange, error)
	Fields(s ...googleapi.Field) *sheets.SpreadsheetsValuesGetCall
	Header() http.Header
	IfNoneMatch(entityTag string) *sheets.SpreadsheetsValuesGetCall
	MajorDimension(majorDimension string) *sheets.SpreadsheetsValuesGetCall
	ValueRenderOption(valueRenderOption string) *sheets.SpreadsheetsValuesGetCall
}

type GSheetService struct {
	Values            V4SpreadsheetsValuesService
	DeveloperMetadata V4DeveloperMetadata
}

func (s *GSheetService) SetService(srv *sheets.Service) {
	if srv == nil {
		srv, _ = sheets.NewService(context.TODO())
	}

	s.Values = srv.Spreadsheets.Values
	s.DeveloperMetadata = srv.Spreadsheets.DeveloperMetadata
}

/*
  - Example:
    spreadsheetId := "1AGHH6abcXzBmfC5e9r50t3wXKhlUs5XIE-fj1U4fV0Q"
    readRange := "Log1!A2:B"
    *
*/
func (s *GSheetService) ReadSheet(spreadsheetID, readRange string) (values [][]string) {
	caller := s.Values.Get(spreadsheetID, readRange)
	resp, err := caller.Do()
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
func (s *GSheetService) Append(
	spreadsheetID string, range_ string, values [][]any,
) (resp *sheets.AppendValuesResponse, err error) {
	valuerange := &sheets.ValueRange{
		MajorDimension: constant.MajorDimensionRows.String(),
		Range:          range_,
		Values:         values,
	}

	caller := s.Values.Append(spreadsheetID, range_, valuerange)
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

func (s *GSheetService) Update(
	spreadsheetID, updateRange string, values *sheets.ValueRange,
) (result *sheets.UpdateValuesResponse, err error) {
	caller := s.Values.Update(spreadsheetID, updateRange, values)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}

func (s *GSheetService) Search(
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
	caller := s.DeveloperMetadata.Search(spreadsheetID, searchdevelopermetadatarequest)
	result, err = caller.Do()
	if err != nil {
		log.Default().Println("Unable to update data to sheet: ", err)
	}

	return
}
