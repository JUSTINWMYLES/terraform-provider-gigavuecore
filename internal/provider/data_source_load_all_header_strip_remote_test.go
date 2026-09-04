package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllHeaderStripDataSource_Read_Happy exercises LoadAllHeaderStripDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllHeaderStripDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllHeaderStripDataSource{client: newMockClientStatus(t, 200, "{\"headerStripsDef\":[]}")}
	m := LoadAllHeaderStripDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllHeaderStripDataSource_Read_NilClient exercises LoadAllHeaderStripDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllHeaderStripDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllHeaderStripDataSource{}
	m := LoadAllHeaderStripDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllHeaderStripDataSource_Read_BuildError exercises LoadAllHeaderStripDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllHeaderStripDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllHeaderStripDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllHeaderStripDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeaderStripDataSource_Read_SendError exercises LoadAllHeaderStripDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllHeaderStripDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllHeaderStripDataSource{client: newTransportErrorClient(t)}
	m := LoadAllHeaderStripDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeaderStripDataSource_Read_InvalidJSON exercises LoadAllHeaderStripDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllHeaderStripDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllHeaderStripDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllHeaderStripDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
