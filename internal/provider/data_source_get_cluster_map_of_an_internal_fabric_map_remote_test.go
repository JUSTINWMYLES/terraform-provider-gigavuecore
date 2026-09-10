package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_Happy exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_NilClient exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_BuildError exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_SendError exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_NotFound exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_APIError exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_map_of_an_internal_fabric_map")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_APIErrorReadBody exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterMapOfAnInternalFabricMapDataSource_Read_InvalidJSON exercises GetClusterMapOfAnInternalFabricMapDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterMapOfAnInternalFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterMapOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterMapOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
