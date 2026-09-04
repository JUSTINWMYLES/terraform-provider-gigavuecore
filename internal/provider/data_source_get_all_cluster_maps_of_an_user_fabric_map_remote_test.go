package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_Happy exercises GetAllClusterMapsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{\"clusterMaps\":[]}")}
	m := GetAllClusterMapsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_NilClient exercises GetAllClusterMapsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapDataSource{}
	m := GetAllClusterMapsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_BuildError exercises GetAllClusterMapsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterMapsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_SendError exercises GetAllClusterMapsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterMapsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_InvalidJSON exercises GetAllClusterMapsOfAnUserFabricMapDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllClusterMapsOfAnUserFabricMapDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterMapsOfAnUserFabricMapDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterMapsOfAnUserFabricMapDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
