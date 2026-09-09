package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSessionAwareApfProfilesDataSource_Read_Happy exercises LoadAllSessionAwareApfProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSessionAwareApfProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSessionAwareApfProfilesDataSource{client: newMockClientStatus(t, 200, "{\"saApfProfiles\":[]}")}
	m := LoadAllSessionAwareApfProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSessionAwareApfProfilesDataSource_Read_NilClient exercises LoadAllSessionAwareApfProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSessionAwareApfProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSessionAwareApfProfilesDataSource{}
	m := LoadAllSessionAwareApfProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSessionAwareApfProfilesDataSource_Read_BuildError exercises LoadAllSessionAwareApfProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSessionAwareApfProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSessionAwareApfProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSessionAwareApfProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSessionAwareApfProfilesDataSource_Read_SendError exercises LoadAllSessionAwareApfProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSessionAwareApfProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSessionAwareApfProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSessionAwareApfProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSessionAwareApfProfilesDataSource_Read_InvalidJSON exercises LoadAllSessionAwareApfProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSessionAwareApfProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSessionAwareApfProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSessionAwareApfProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
