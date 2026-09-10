package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_Happy exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_NilClient exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_BuildError exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_SendError exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_NotFound exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_NotFound(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_APIError exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_APIError(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_APIErrorReadBody exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_InvalidJSON exercises GetClusterCircuitTunnelOfAnInternalFabricMapDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetClusterCircuitTunnelOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
