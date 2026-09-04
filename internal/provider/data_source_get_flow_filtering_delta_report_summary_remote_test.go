package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_Happy exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_NilClient exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_BuildError exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_SendError exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_NotFound exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_APIError exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_filtering_delta_report_summary")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_APIErrorReadBody exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowFilteringDeltaReportSummaryDataSource_Read_InvalidJSON exercises GetFlowFilteringDeltaReportSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowFilteringDeltaReportSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowFilteringDeltaReportSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowFilteringDeltaReportSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
