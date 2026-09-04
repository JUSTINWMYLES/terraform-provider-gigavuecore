package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGsGroupsDataSource_Read_Happy exercises LoadAllGsGroupsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllGsGroupsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllGsGroupsDataSource{client: newMockClientStatus(t, 200, "{\"gsGroups\":[]}")}
	m := LoadAllGsGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllGsGroupsDataSource_Read_NilClient exercises LoadAllGsGroupsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllGsGroupsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllGsGroupsDataSource{}
	m := LoadAllGsGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllGsGroupsDataSource_Read_BuildError exercises LoadAllGsGroupsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllGsGroupsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllGsGroupsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllGsGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGsGroupsDataSource_Read_SendError exercises LoadAllGsGroupsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllGsGroupsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllGsGroupsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllGsGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGsGroupsDataSource_Read_InvalidJSON exercises LoadAllGsGroupsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllGsGroupsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllGsGroupsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllGsGroupsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
