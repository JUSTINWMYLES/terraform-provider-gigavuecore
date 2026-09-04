package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNetflowExporterFilterDataSource_Read_Happy exercises GetNetflowExporterFilterDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNetflowExporterFilterDataSource_Read_Happy(t *testing.T) {
	r := &GetNetflowExporterFilterDataSource{client: newMockClientStatus(t, 200, "{\"rules\":[]}")}
	m := GetNetflowExporterFilterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNetflowExporterFilterDataSource_Read_NilClient exercises GetNetflowExporterFilterDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNetflowExporterFilterDataSource_Read_NilClient(t *testing.T) {
	r := &GetNetflowExporterFilterDataSource{}
	m := GetNetflowExporterFilterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNetflowExporterFilterDataSource_Read_BuildError exercises GetNetflowExporterFilterDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNetflowExporterFilterDataSource_Read_BuildError(t *testing.T) {
	r := &GetNetflowExporterFilterDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNetflowExporterFilterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNetflowExporterFilterDataSource_Read_SendError exercises GetNetflowExporterFilterDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNetflowExporterFilterDataSource_Read_SendError(t *testing.T) {
	r := &GetNetflowExporterFilterDataSource{client: newTransportErrorClient(t)}
	m := GetNetflowExporterFilterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNetflowExporterFilterDataSource_Read_InvalidJSON exercises GetNetflowExporterFilterDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNetflowExporterFilterDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNetflowExporterFilterDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNetflowExporterFilterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
