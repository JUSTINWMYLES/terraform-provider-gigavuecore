package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllV3UsersDataSource_Read_Happy exercises LoadAllV3UsersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllV3UsersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllV3UsersDataSource{client: newMockClientStatus(t, 200, "{\"v3Users\":[]}")}
	m := LoadAllV3UsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllV3UsersDataSource_Read_NilClient exercises LoadAllV3UsersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllV3UsersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllV3UsersDataSource{}
	m := LoadAllV3UsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllV3UsersDataSource_Read_BuildError exercises LoadAllV3UsersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllV3UsersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllV3UsersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllV3UsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllV3UsersDataSource_Read_SendError exercises LoadAllV3UsersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllV3UsersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllV3UsersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllV3UsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllV3UsersDataSource_Read_InvalidJSON exercises LoadAllV3UsersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllV3UsersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllV3UsersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllV3UsersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
