package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowOpsReportDataSource_Read_Happy exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowOpsReportDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowOpsReportDataSource_Read_NilClient exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowOpsReportDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowOpsReportDataSource{}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowOpsReportDataSource_Read_BuildError exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowOpsReportDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowOpsReportDataSource_Read_SendError exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowOpsReportDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newTransportErrorClient(t)}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowOpsReportDataSource_Read_NotFound exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowOpsReportDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowOpsReportDataSource_Read_APIError exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowOpsReportDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_ops_report")
}

// TestGetFlowOpsReportDataSource_Read_APIErrorReadBody exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowOpsReportDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowOpsReportDataSource_Read_InvalidJSON exercises GetFlowOpsReportDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowOpsReportDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowOpsReportDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowOpsReportDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
