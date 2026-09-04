package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGigaPortsDataSource_Read_Happy exercises LoadGigaPortsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadGigaPortsDataSource_Read_Happy(t *testing.T) {
	r := &LoadGigaPortsDataSource{client: newMockClientStatus(t, 200, "{\"ports\":[]}")}
	m := LoadGigaPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGigaPortsDataSource_Read_NilClient exercises LoadGigaPortsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGigaPortsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGigaPortsDataSource{}
	m := LoadGigaPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGigaPortsDataSource_Read_BuildError exercises LoadGigaPortsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadGigaPortsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGigaPortsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGigaPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGigaPortsDataSource_Read_SendError exercises LoadGigaPortsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadGigaPortsDataSource_Read_SendError(t *testing.T) {
	r := &LoadGigaPortsDataSource{client: newTransportErrorClient(t)}
	m := LoadGigaPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGigaPortsDataSource_Read_InvalidJSON exercises LoadGigaPortsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadGigaPortsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGigaPortsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGigaPortsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
