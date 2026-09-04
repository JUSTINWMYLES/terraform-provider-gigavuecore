package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllVportsDataSource_Read_Happy exercises LoadAllVportsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllVportsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllVportsDataSource{client: newMockClientStatus(t, 200, "{\"vports\":[]}")}
	m := LoadAllVportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllVportsDataSource_Read_NilClient exercises LoadAllVportsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllVportsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllVportsDataSource{}
	m := LoadAllVportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllVportsDataSource_Read_BuildError exercises LoadAllVportsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllVportsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllVportsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllVportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllVportsDataSource_Read_SendError exercises LoadAllVportsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllVportsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllVportsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllVportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllVportsDataSource_Read_InvalidJSON exercises LoadAllVportsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllVportsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllVportsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllVportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
