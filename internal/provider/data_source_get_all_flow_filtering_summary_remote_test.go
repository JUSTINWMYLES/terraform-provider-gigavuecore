package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlowFilteringSummaryDataSource_Read_Happy exercises GetAllFlowFilteringSummaryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFlowFilteringSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 200, "{\"flowFilteringReportsSummary\":[]}")}
	m := GetAllFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFlowFilteringSummaryDataSource_Read_NilClient exercises GetAllFlowFilteringSummaryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFlowFilteringSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFlowFilteringSummaryDataSource{}
	m := GetAllFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFlowFilteringSummaryDataSource_Read_BuildError exercises GetAllFlowFilteringSummaryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFlowFilteringSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFlowFilteringSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlowFilteringSummaryDataSource_Read_SendError exercises GetAllFlowFilteringSummaryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFlowFilteringSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFlowFilteringSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetAllFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlowFilteringSummaryDataSource_Read_InvalidJSON exercises GetAllFlowFilteringSummaryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFlowFilteringSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFlowFilteringSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFlowFilteringSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
