package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNegativeHeartbeatProfilesDataSource_Read_Happy exercises LoadAllNegativeHeartbeatProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNegativeHeartbeatProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNegativeHeartbeatProfilesDataSource{client: newMockClientStatus(t, 200, "{\"negativeHeartbeatProfiles\":[]}")}
	m := LoadAllNegativeHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNegativeHeartbeatProfilesDataSource_Read_NilClient exercises LoadAllNegativeHeartbeatProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNegativeHeartbeatProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNegativeHeartbeatProfilesDataSource{}
	m := LoadAllNegativeHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNegativeHeartbeatProfilesDataSource_Read_BuildError exercises LoadAllNegativeHeartbeatProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNegativeHeartbeatProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNegativeHeartbeatProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNegativeHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNegativeHeartbeatProfilesDataSource_Read_SendError exercises LoadAllNegativeHeartbeatProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNegativeHeartbeatProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNegativeHeartbeatProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNegativeHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNegativeHeartbeatProfilesDataSource_Read_InvalidJSON exercises LoadAllNegativeHeartbeatProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNegativeHeartbeatProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNegativeHeartbeatProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNegativeHeartbeatProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
