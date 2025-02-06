// Copyright (c) 2025 Sigma-Firma
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package gsheet

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func (a *Access) Sheets() *sheets.Service {
	service, err := sheets.NewService(context.Background(), option.WithHTTPClient(a.GetClient()))
	if err != nil {
		log.Println(err)
	}
	a.SheetsAPI = service
	return a.SheetsAPI
}

// Write is used to write to a spreadsheet
func (a *Access) Write(srv *sheets.Service, spreadsheetId string, writeRange string, vals []interface{}) (*sheets.AppendValuesResponse, error) {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, vals)

	return srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
}

// Read is used to read from a spreadsheet
func (a *Access) Read(srv *sheets.Service, spreadsheetId string, readRange string) (*sheets.ValueRange, error) {
	return srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
}
