package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNameServersDataSource_Read_Happy exercises LoadAllNameServersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNameServersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNameServersDataSource{client: newMockClientStatus(t, 200, "{\"nameServers\":[]}")}
	m := LoadAllNameServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNameServersDataSource_Read_NilClient exercises LoadAllNameServersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNameServersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNameServersDataSource{}
	m := LoadAllNameServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNameServersDataSource_Read_BuildError exercises LoadAllNameServersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNameServersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNameServersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNameServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNameServersDataSource_Read_SendError exercises LoadAllNameServersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNameServersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNameServersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNameServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNameServersDataSource_Read_InvalidJSON exercises LoadAllNameServersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNameServersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNameServersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNameServersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
