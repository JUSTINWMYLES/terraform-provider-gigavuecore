package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapProfilesDataSource_Read_Happy exercises LoadAllIcapProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllIcapProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllIcapProfilesDataSource{client: newMockClientStatus(t, 200, "{\"icapProfiles\":[]}")}
	m := LoadAllIcapProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllIcapProfilesDataSource_Read_NilClient exercises LoadAllIcapProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllIcapProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllIcapProfilesDataSource{}
	m := LoadAllIcapProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllIcapProfilesDataSource_Read_BuildError exercises LoadAllIcapProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllIcapProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllIcapProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllIcapProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapProfilesDataSource_Read_SendError exercises LoadAllIcapProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllIcapProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllIcapProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllIcapProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllIcapProfilesDataSource_Read_InvalidJSON exercises LoadAllIcapProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllIcapProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllIcapProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllIcapProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
