package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlowSipSummaryDataSource_Read_Happy exercises GetAllFlowSipSummaryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFlowSipSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFlowSipSummaryDataSource{client: newMockClientStatus(t, 200, "{\"flowSipReportsSummary\":[]}")}
	m := GetAllFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFlowSipSummaryDataSource_Read_NilClient exercises GetAllFlowSipSummaryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFlowSipSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFlowSipSummaryDataSource{}
	m := GetAllFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFlowSipSummaryDataSource_Read_BuildError exercises GetAllFlowSipSummaryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFlowSipSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFlowSipSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlowSipSummaryDataSource_Read_SendError exercises GetAllFlowSipSummaryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFlowSipSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFlowSipSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetAllFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlowSipSummaryDataSource_Read_InvalidJSON exercises GetAllFlowSipSummaryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFlowSipSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFlowSipSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
