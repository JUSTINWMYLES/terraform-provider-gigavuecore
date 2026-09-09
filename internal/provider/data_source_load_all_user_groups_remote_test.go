package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUserGroupsDataSource_Read_Happy exercises LoadAllUserGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllUserGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllUserGroupsDataSource{client: newMockClientStatus(t, 200, "{\"groups\":[]}")}
	m := LoadAllUserGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllUserGroupsDataSource_Read_NilClient exercises LoadAllUserGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllUserGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllUserGroupsDataSource{}
	m := LoadAllUserGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllUserGroupsDataSource_Read_BuildError exercises LoadAllUserGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllUserGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllUserGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllUserGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUserGroupsDataSource_Read_SendError exercises LoadAllUserGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllUserGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllUserGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllUserGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUserGroupsDataSource_Read_InvalidJSON exercises LoadAllUserGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllUserGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllUserGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllUserGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
