package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTunnelLogicalGroupStatsDataSource_Read_Happy exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetTunnelLogicalGroupStatsDataSource_Read_Happy(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_NilClient exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTunnelLogicalGroupStatsDataSource_Read_NilClient(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_BuildError exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetTunnelLogicalGroupStatsDataSource_Read_BuildError(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_SendError exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetTunnelLogicalGroupStatsDataSource_Read_SendError(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newTransportErrorClient(t)}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_NotFound exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetTunnelLogicalGroupStatsDataSource_Read_NotFound(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_APIError exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetTunnelLogicalGroupStatsDataSource_Read_APIError(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_tunnel_logical_group_stats")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_APIErrorReadBody exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetTunnelLogicalGroupStatsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetTunnelLogicalGroupStatsDataSource_Read_InvalidJSON exercises GetTunnelLogicalGroupStatsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetTunnelLogicalGroupStatsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTunnelLogicalGroupStatsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTunnelLogicalGroupStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
