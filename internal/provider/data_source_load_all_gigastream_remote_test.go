package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGigastreamDataSource_Read_Happy exercises LoadAllGigastreamDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllGigastreamDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllGigastreamDataSource{client: newMockClientStatus(t, 200, "{\"gigastreams\":[]}")}
	m := LoadAllGigastreamDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllGigastreamDataSource_Read_NilClient exercises LoadAllGigastreamDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllGigastreamDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllGigastreamDataSource{}
	m := LoadAllGigastreamDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllGigastreamDataSource_Read_BuildError exercises LoadAllGigastreamDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllGigastreamDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllGigastreamDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllGigastreamDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGigastreamDataSource_Read_SendError exercises LoadAllGigastreamDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllGigastreamDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllGigastreamDataSource{client: newTransportErrorClient(t)}
	m := LoadAllGigastreamDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllGigastreamDataSource_Read_InvalidJSON exercises LoadAllGigastreamDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllGigastreamDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllGigastreamDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllGigastreamDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
