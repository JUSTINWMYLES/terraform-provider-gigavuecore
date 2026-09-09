package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTroubleshootSummaryDataSource_Read_Happy exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTroubleshootSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTroubleshootSummaryDataSource_Read_NilClient exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTroubleshootSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTroubleshootSummaryDataSource_Read_BuildError exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTroubleshootSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTroubleshootSummaryDataSource_Read_SendError exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTroubleshootSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTroubleshootSummaryDataSource_Read_NotFound exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTroubleshootSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTroubleshootSummaryDataSource_Read_APIError exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTroubleshootSummaryDataSource_Read_APIError(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_troubleshoot_summary")
}

// TestGetTroubleshootSummaryDataSource_Read_APIErrorReadBody exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTroubleshootSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTroubleshootSummaryDataSource_Read_InvalidJSON exercises GetTroubleshootSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTroubleshootSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTroubleshootSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTroubleshootSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
