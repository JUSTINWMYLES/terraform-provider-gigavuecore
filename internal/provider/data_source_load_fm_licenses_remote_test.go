package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmLicensesDataSource_Read_Happy exercises LoadFmLicensesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadFmLicensesDataSource_Read_Happy(t *testing.T) {
	r := &LoadFmLicensesDataSource{client: newMockClientStatus(t, 200, "{\"licenses\":[]}")}
	m := LoadFmLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFmLicensesDataSource_Read_NilClient exercises LoadFmLicensesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFmLicensesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFmLicensesDataSource{}
	m := LoadFmLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFmLicensesDataSource_Read_BuildError exercises LoadFmLicensesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadFmLicensesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFmLicensesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFmLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmLicensesDataSource_Read_SendError exercises LoadFmLicensesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadFmLicensesDataSource_Read_SendError(t *testing.T) {
	r := &LoadFmLicensesDataSource{client: newTransportErrorClient(t)}
	m := LoadFmLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFmLicensesDataSource_Read_InvalidJSON exercises LoadFmLicensesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadFmLicensesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFmLicensesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFmLicensesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
