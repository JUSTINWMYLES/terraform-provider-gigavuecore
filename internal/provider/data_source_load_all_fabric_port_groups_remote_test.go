package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFabricPortGroupsDataSource_Read_Happy exercises LoadAllFabricPortGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFabricPortGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFabricPortGroupsDataSource{client: newMockClientStatus(t, 200, "{\"gigaFabricPortGroups\":[]}")}
	m := LoadAllFabricPortGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFabricPortGroupsDataSource_Read_NilClient exercises LoadAllFabricPortGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFabricPortGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFabricPortGroupsDataSource{}
	m := LoadAllFabricPortGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFabricPortGroupsDataSource_Read_BuildError exercises LoadAllFabricPortGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFabricPortGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFabricPortGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFabricPortGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFabricPortGroupsDataSource_Read_SendError exercises LoadAllFabricPortGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFabricPortGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFabricPortGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFabricPortGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFabricPortGroupsDataSource_Read_InvalidJSON exercises LoadAllFabricPortGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFabricPortGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFabricPortGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFabricPortGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
