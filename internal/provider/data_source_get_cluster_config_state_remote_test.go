package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterConfigStateDataSource_Read_Happy exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterConfigStateDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterConfigStateDataSource_Read_NilClient exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterConfigStateDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterConfigStateDataSource{}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterConfigStateDataSource_Read_BuildError exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterConfigStateDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterConfigStateDataSource_Read_SendError exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterConfigStateDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newTransportErrorClient(t)}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterConfigStateDataSource_Read_NotFound exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterConfigStateDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterConfigStateDataSource_Read_APIError exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterConfigStateDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_config_state")
}

// TestGetClusterConfigStateDataSource_Read_APIErrorReadBody exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterConfigStateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterConfigStateDataSource_Read_InvalidJSON exercises GetClusterConfigStateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterConfigStateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterConfigStateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterConfigStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
