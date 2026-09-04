package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_Happy exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_NilClient exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_BuildError exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_SendError exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_NotFound exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_APIError exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_circuit_tunnel_of_an_user_fabric_map")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_APIErrorReadBody exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_InvalidJSON exercises GetClusterCircuitTunnelOfAnUserFabricMapDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterCircuitTunnelOfAnUserFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterCircuitTunnelOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
