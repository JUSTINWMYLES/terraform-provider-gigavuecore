package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemArpRefreshIntervalDataSource_Read_Happy exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemArpRefreshIntervalDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemArpRefreshIntervalDataSource_Read_NilClient exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemArpRefreshIntervalDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_BuildError exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemArpRefreshIntervalDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_SendError exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemArpRefreshIntervalDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newTransportErrorClient(t)}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_NotFound exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemArpRefreshIntervalDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_APIError exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemArpRefreshIntervalDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_arp_refresh_interval")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_APIErrorReadBody exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemArpRefreshIntervalDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemArpRefreshIntervalDataSource_Read_InvalidJSON exercises GetSystemArpRefreshIntervalDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemArpRefreshIntervalDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemArpRefreshIntervalDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemArpRefreshIntervalDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
