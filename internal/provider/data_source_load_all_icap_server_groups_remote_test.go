package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapServerGroupsDataSource_Read_Happy exercises LoadAllIcapServerGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllIcapServerGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllIcapServerGroupsDataSource{client: newMockClientStatus(t, 200, "{\"icapServerGroups\":[]}")}
	m := LoadAllIcapServerGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllIcapServerGroupsDataSource_Read_NilClient exercises LoadAllIcapServerGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllIcapServerGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllIcapServerGroupsDataSource{}
	m := LoadAllIcapServerGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllIcapServerGroupsDataSource_Read_BuildError exercises LoadAllIcapServerGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllIcapServerGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllIcapServerGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllIcapServerGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapServerGroupsDataSource_Read_SendError exercises LoadAllIcapServerGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllIcapServerGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllIcapServerGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllIcapServerGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapServerGroupsDataSource_Read_InvalidJSON exercises LoadAllIcapServerGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllIcapServerGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllIcapServerGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllIcapServerGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
