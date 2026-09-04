package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllTacacsServersDataSource_Read_Happy exercises LoadAllTacacsServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllTacacsServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllTacacsServersDataSource{client: newMockClientStatus(t, 200, "{\"tacacsServers\":[]}")}
	m := LoadAllTacacsServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllTacacsServersDataSource_Read_NilClient exercises LoadAllTacacsServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllTacacsServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllTacacsServersDataSource{}
	m := LoadAllTacacsServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllTacacsServersDataSource_Read_BuildError exercises LoadAllTacacsServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllTacacsServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllTacacsServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllTacacsServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllTacacsServersDataSource_Read_SendError exercises LoadAllTacacsServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllTacacsServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllTacacsServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllTacacsServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllTacacsServersDataSource_Read_InvalidJSON exercises LoadAllTacacsServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllTacacsServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllTacacsServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllTacacsServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
