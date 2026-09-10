package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMetadataApplicationProfilesDataSource_Read_Happy exercises LoadAllMetadataApplicationProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllMetadataApplicationProfilesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllMetadataApplicationProfilesDataSource{client: newMockClientStatus(t, 200, "{\"applicationProfiles\":[]}")}
	m := LoadAllMetadataApplicationProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllMetadataApplicationProfilesDataSource_Read_NilClient exercises LoadAllMetadataApplicationProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllMetadataApplicationProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllMetadataApplicationProfilesDataSource{}
	m := LoadAllMetadataApplicationProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllMetadataApplicationProfilesDataSource_Read_BuildError exercises LoadAllMetadataApplicationProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllMetadataApplicationProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllMetadataApplicationProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllMetadataApplicationProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataApplicationProfilesDataSource_Read_SendError exercises LoadAllMetadataApplicationProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllMetadataApplicationProfilesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllMetadataApplicationProfilesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllMetadataApplicationProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllMetadataApplicationProfilesDataSource_Read_InvalidJSON exercises LoadAllMetadataApplicationProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllMetadataApplicationProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllMetadataApplicationProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllMetadataApplicationProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
