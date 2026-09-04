package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUsersDataSource_Read_Happy exercises LoadAllUsersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllUsersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllUsersDataSource{client: newMockClientStatus(t, 200, "{\"users\":[]}")}
	m := LoadAllUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllUsersDataSource_Read_NilClient exercises LoadAllUsersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllUsersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllUsersDataSource{}
	m := LoadAllUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllUsersDataSource_Read_BuildError exercises LoadAllUsersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllUsersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllUsersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUsersDataSource_Read_SendError exercises LoadAllUsersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllUsersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllUsersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllUsersDataSource_Read_InvalidJSON exercises LoadAllUsersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllUsersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllUsersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllUsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
