package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineSslProfilesDataSource_Read_Happy exercises LoadAllInlineSslProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllInlineSslProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllInlineSslProfilesDataSource{client: newMockClientStatus(t, 200, "{\"inlineSslProfiles\":[]}")}
	m := LoadAllInlineSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllInlineSslProfilesDataSource_Read_NilClient exercises LoadAllInlineSslProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllInlineSslProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllInlineSslProfilesDataSource{}
	m := LoadAllInlineSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllInlineSslProfilesDataSource_Read_BuildError exercises LoadAllInlineSslProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllInlineSslProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllInlineSslProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllInlineSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineSslProfilesDataSource_Read_SendError exercises LoadAllInlineSslProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllInlineSslProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllInlineSslProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllInlineSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllInlineSslProfilesDataSource_Read_InvalidJSON exercises LoadAllInlineSslProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllInlineSslProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllInlineSslProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllInlineSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
