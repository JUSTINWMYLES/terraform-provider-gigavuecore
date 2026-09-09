package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowSipSummaryDataSource_Read_Happy exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFlowSipSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlowSipSummaryDataSource_Read_NilClient exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlowSipSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlowSipSummaryDataSource_Read_BuildError exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFlowSipSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFlowSipSummaryDataSource_Read_SendError exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFlowSipSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFlowSipSummaryDataSource_Read_NotFound exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFlowSipSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFlowSipSummaryDataSource_Read_APIError exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFlowSipSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_flow_sip_summary")
}

// TestGetFlowSipSummaryDataSource_Read_APIErrorReadBody exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFlowSipSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFlowSipSummaryDataSource_Read_InvalidJSON exercises GetFlowSipSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFlowSipSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlowSipSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlowSipSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
