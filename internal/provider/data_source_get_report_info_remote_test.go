package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetReportInfoDataSource_Read_Happy exercises GetReportInfoDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetReportInfoDataSource_Read_Happy(t *testing.T) {
	r := &GetReportInfoDataSource{client: newMockClientStatus(t, 200, "{\"fmUserTokenEntities\":[]}")}
	m := GetReportInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetReportInfoDataSource_Read_NilClient exercises GetReportInfoDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetReportInfoDataSource_Read_NilClient(t *testing.T) {
	r := &GetReportInfoDataSource{}
	m := GetReportInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetReportInfoDataSource_Read_BuildError exercises GetReportInfoDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetReportInfoDataSource_Read_BuildError(t *testing.T) {
	r := &GetReportInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetReportInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetReportInfoDataSource_Read_SendError exercises GetReportInfoDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetReportInfoDataSource_Read_SendError(t *testing.T) {
	r := &GetReportInfoDataSource{client: newTransportErrorClient(t)}
	m := GetReportInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetReportInfoDataSource_Read_InvalidJSON exercises GetReportInfoDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetReportInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetReportInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetReportInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
