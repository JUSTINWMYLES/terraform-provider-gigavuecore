package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllLocalUsersDataSource_Read_Happy exercises LoadAllLocalUsersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllLocalUsersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllLocalUsersDataSource{client: newMockClientStatus(t, 200, "{\"localUsers\":[]}")}
	m := LoadAllLocalUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllLocalUsersDataSource_Read_NilClient exercises LoadAllLocalUsersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllLocalUsersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllLocalUsersDataSource{}
	m := LoadAllLocalUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllLocalUsersDataSource_Read_BuildError exercises LoadAllLocalUsersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllLocalUsersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllLocalUsersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllLocalUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllLocalUsersDataSource_Read_SendError exercises LoadAllLocalUsersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllLocalUsersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllLocalUsersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllLocalUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllLocalUsersDataSource_Read_InvalidJSON exercises LoadAllLocalUsersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllLocalUsersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllLocalUsersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllLocalUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
