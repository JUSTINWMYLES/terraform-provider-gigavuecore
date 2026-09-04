package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSysStatsDataSource_Read_Happy exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetNodeSysStatsDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeSysStatsDataSource_Read_NilClient exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeSysStatsDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeSysStatsDataSource{}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeSysStatsDataSource_Read_BuildError exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetNodeSysStatsDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetNodeSysStatsDataSource_Read_SendError exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetNodeSysStatsDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newTransportErrorClient(t)}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetNodeSysStatsDataSource_Read_NotFound exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetNodeSysStatsDataSource_Read_NotFound(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetNodeSysStatsDataSource_Read_APIError exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetNodeSysStatsDataSource_Read_APIError(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_node_sys_stats")
}

// TestGetNodeSysStatsDataSource_Read_APIErrorReadBody exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetNodeSysStatsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetNodeSysStatsDataSource_Read_InvalidJSON exercises GetNodeSysStatsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetNodeSysStatsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeSysStatsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeSysStatsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
