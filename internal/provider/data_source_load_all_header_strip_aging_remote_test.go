package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllHeaderStripAgingDataSource_Read_Happy exercises LoadAllHeaderStripAgingDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllHeaderStripAgingDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllHeaderStripAgingDataSource{client: newMockClientStatus(t, 200, "{\"headerStripsAgingDef\":[]}")}
	m := LoadAllHeaderStripAgingDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllHeaderStripAgingDataSource_Read_NilClient exercises LoadAllHeaderStripAgingDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllHeaderStripAgingDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllHeaderStripAgingDataSource{}
	m := LoadAllHeaderStripAgingDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllHeaderStripAgingDataSource_Read_BuildError exercises LoadAllHeaderStripAgingDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllHeaderStripAgingDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllHeaderStripAgingDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllHeaderStripAgingDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeaderStripAgingDataSource_Read_SendError exercises LoadAllHeaderStripAgingDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllHeaderStripAgingDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllHeaderStripAgingDataSource{client: newTransportErrorClient(t)}
	m := LoadAllHeaderStripAgingDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeaderStripAgingDataSource_Read_InvalidJSON exercises LoadAllHeaderStripAgingDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllHeaderStripAgingDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllHeaderStripAgingDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllHeaderStripAgingDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
