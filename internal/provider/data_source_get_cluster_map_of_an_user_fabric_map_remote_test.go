package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_Happy exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_NilClient exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_BuildError exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_SendError exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_NotFound exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_APIError exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_map_of_an_user_fabric_map")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_APIErrorReadBody exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterMapOfAnUserFabricMapDataSource_Read_InvalidJSON exercises GetClusterMapOfAnUserFabricMapDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterMapOfAnUserFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterMapOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterMapOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
