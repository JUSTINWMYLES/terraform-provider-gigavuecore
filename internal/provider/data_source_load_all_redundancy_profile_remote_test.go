package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllRedundancyProfileDataSource_Read_Happy exercises LoadAllRedundancyProfileDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllRedundancyProfileDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllRedundancyProfileDataSource{client: newMockClientStatus(t, 200, "{\"redundancyProfiles\":[]}")}
	m := LoadAllRedundancyProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllRedundancyProfileDataSource_Read_NilClient exercises LoadAllRedundancyProfileDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllRedundancyProfileDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllRedundancyProfileDataSource{}
	m := LoadAllRedundancyProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllRedundancyProfileDataSource_Read_BuildError exercises LoadAllRedundancyProfileDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllRedundancyProfileDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllRedundancyProfileDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllRedundancyProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRedundancyProfileDataSource_Read_SendError exercises LoadAllRedundancyProfileDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllRedundancyProfileDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllRedundancyProfileDataSource{client: newTransportErrorClient(t)}
	m := LoadAllRedundancyProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRedundancyProfileDataSource_Read_InvalidJSON exercises LoadAllRedundancyProfileDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllRedundancyProfileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllRedundancyProfileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllRedundancyProfileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
