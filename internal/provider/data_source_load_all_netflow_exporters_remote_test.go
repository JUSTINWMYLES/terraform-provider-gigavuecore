package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNetflowExportersDataSource_Read_Happy exercises LoadAllNetflowExportersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllNetflowExportersDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllNetflowExportersDataSource{client: newMockClientStatus(t, 200, "{\"nfExporters\":[]}")}
	m := LoadAllNetflowExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllNetflowExportersDataSource_Read_NilClient exercises LoadAllNetflowExportersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllNetflowExportersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllNetflowExportersDataSource{}
	m := LoadAllNetflowExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllNetflowExportersDataSource_Read_BuildError exercises LoadAllNetflowExportersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllNetflowExportersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllNetflowExportersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllNetflowExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowExportersDataSource_Read_SendError exercises LoadAllNetflowExportersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllNetflowExportersDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllNetflowExportersDataSource{client: newTransportErrorClient(t)}
	m := LoadAllNetflowExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllNetflowExportersDataSource_Read_InvalidJSON exercises LoadAllNetflowExportersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllNetflowExportersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllNetflowExportersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllNetflowExportersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
