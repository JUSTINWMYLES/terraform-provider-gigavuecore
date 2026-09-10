package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSysdumpDataSource_Read_Happy exercises LoadSysdumpDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadSysdumpDataSource_Read_Happy(t *testing.T) {
	r := &LoadSysdumpDataSource{client: newMockClientStatus(t, 200, "{\"sysdumpFiles\":[]}")}
	m := LoadSysdumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSysdumpDataSource_Read_NilClient exercises LoadSysdumpDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSysdumpDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSysdumpDataSource{}
	m := LoadSysdumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSysdumpDataSource_Read_BuildError exercises LoadSysdumpDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadSysdumpDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSysdumpDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSysdumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSysdumpDataSource_Read_SendError exercises LoadSysdumpDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadSysdumpDataSource_Read_SendError(t *testing.T) {
	r := &LoadSysdumpDataSource{client: newTransportErrorClient(t)}
	m := LoadSysdumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadSysdumpDataSource_Read_InvalidJSON exercises LoadSysdumpDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadSysdumpDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSysdumpDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSysdumpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
