package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllRolesDataSource_Read_Happy exercises LoadAllRolesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllRolesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllRolesDataSource{client: newMockClientStatus(t, 200, "{\"roles\":[]}")}
	m := LoadAllRolesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllRolesDataSource_Read_NilClient exercises LoadAllRolesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllRolesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllRolesDataSource{}
	m := LoadAllRolesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllRolesDataSource_Read_BuildError exercises LoadAllRolesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllRolesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllRolesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllRolesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRolesDataSource_Read_SendError exercises LoadAllRolesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllRolesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllRolesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllRolesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllRolesDataSource_Read_InvalidJSON exercises LoadAllRolesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllRolesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllRolesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllRolesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
