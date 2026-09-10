package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPortThrottleReportByAliasDataSource_Read_Happy exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPortThrottleReportByAliasDataSource_Read_Happy(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPortThrottleReportByAliasDataSource_Read_NilClient exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPortThrottleReportByAliasDataSource_Read_NilClient(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPortThrottleReportByAliasDataSource_Read_BuildError exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPortThrottleReportByAliasDataSource_Read_BuildError(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPortThrottleReportByAliasDataSource_Read_SendError exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPortThrottleReportByAliasDataSource_Read_SendError(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newTransportErrorClient(t)}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPortThrottleReportByAliasDataSource_Read_NotFound exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPortThrottleReportByAliasDataSource_Read_NotFound(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPortThrottleReportByAliasDataSource_Read_APIError exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPortThrottleReportByAliasDataSource_Read_APIError(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_port_throttle_report_by_alias")
}

// TestGetPortThrottleReportByAliasDataSource_Read_APIErrorReadBody exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPortThrottleReportByAliasDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPortThrottleReportByAliasDataSource_Read_InvalidJSON exercises GetPortThrottleReportByAliasDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPortThrottleReportByAliasDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPortThrottleReportByAliasDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPortThrottleReportByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
