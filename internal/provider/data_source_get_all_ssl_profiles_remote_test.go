package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSslProfilesDataSource_Read_Happy exercises GetAllSslProfilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllSslProfilesDataSource_Read_Happy(t *testing.T) {
	r := &GetAllSslProfilesDataSource{client: newMockClientStatus(t, 200, "{\"appsSslProfiles\":[]}")}
	m := GetAllSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllSslProfilesDataSource_Read_NilClient exercises GetAllSslProfilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllSslProfilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllSslProfilesDataSource{}
	m := GetAllSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllSslProfilesDataSource_Read_BuildError exercises GetAllSslProfilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllSslProfilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllSslProfilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSslProfilesDataSource_Read_SendError exercises GetAllSslProfilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllSslProfilesDataSource_Read_SendError(t *testing.T) {
	r := &GetAllSslProfilesDataSource{client: newTransportErrorClient(t)}
	m := GetAllSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSslProfilesDataSource_Read_InvalidJSON exercises GetAllSslProfilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllSslProfilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllSslProfilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllSslProfilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
