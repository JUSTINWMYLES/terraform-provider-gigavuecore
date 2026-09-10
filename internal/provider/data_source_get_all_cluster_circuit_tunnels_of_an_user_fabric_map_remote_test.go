package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_Happy exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_NilClient exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_BuildError exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_SendError exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_InvalidJSON exercises GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
