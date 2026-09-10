package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_Happy exercises GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_NilClient exercises GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{}
	m := GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_BuildError exercises GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_SendError exercises GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_InvalidJSON exercises GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
