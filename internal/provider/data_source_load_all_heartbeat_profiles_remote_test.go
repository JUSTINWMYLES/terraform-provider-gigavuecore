package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllHeartbeatProfilesDataSource_Read_Happy exercises LoadAllHeartbeatProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllHeartbeatProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllHeartbeatProfilesDataSource{client: newMockClientStatus(t, 200, "{\"heartbeatProfiles\":[]}")}
	m := LoadAllHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllHeartbeatProfilesDataSource_Read_NilClient exercises LoadAllHeartbeatProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllHeartbeatProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllHeartbeatProfilesDataSource{}
	m := LoadAllHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllHeartbeatProfilesDataSource_Read_BuildError exercises LoadAllHeartbeatProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllHeartbeatProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllHeartbeatProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeartbeatProfilesDataSource_Read_SendError exercises LoadAllHeartbeatProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllHeartbeatProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllHeartbeatProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeartbeatProfilesDataSource_Read_InvalidJSON exercises LoadAllHeartbeatProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllHeartbeatProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllHeartbeatProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
