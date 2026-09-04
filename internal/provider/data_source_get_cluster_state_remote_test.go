package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterStateDataSource_Read_Happy exercises GetClusterStateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterStateDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterStateDataSource_Read_NilClient exercises GetClusterStateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterStateDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterStateDataSource{}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterStateDataSource_Read_BuildError exercises GetClusterStateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterStateDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterStateDataSource_Read_SendError exercises GetClusterStateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterStateDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterStateDataSource{client: newTransportErrorClient(t)}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterStateDataSource_Read_NotFound exercises GetClusterStateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterStateDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterStateDataSource_Read_APIError exercises GetClusterStateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterStateDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_state")
}

// TestGetClusterStateDataSource_Read_APIErrorReadBody exercises GetClusterStateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterStateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterStateDataSource_Read_InvalidJSON exercises GetClusterStateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterStateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterStateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterStateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
