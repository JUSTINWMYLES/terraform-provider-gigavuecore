package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_Happy exercises GetAllClusterMapsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterMapsOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterMapsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_NilClient exercises GetAllClusterMapsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterMapsOfAnInternalFabricMapDataSource{}
	m := GetAllClusterMapsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_BuildError exercises GetAllClusterMapsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterMapsOfAnInternalFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterMapsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_SendError exercises GetAllClusterMapsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterMapsOfAnInternalFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterMapsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_InvalidJSON exercises GetAllClusterMapsOfAnInternalFabricMapDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterMapsOfAnInternalFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterMapsOfAnInternalFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
