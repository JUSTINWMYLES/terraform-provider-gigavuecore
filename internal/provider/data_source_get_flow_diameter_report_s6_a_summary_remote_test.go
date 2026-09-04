package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_Happy exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_NilClient exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_BuildError exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_SendError exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_NotFound exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_APIError exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_diameter_report_s6_a_summary")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_APIErrorReadBody exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowDiameterReportS6ASummaryDataSource_Read_InvalidJSON exercises GetFlowDiameterReportS6ASummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowDiameterReportS6ASummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowDiameterReportS6ASummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowDiameterReportS6ASummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
