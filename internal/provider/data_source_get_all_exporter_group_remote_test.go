package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllExporterGroupDataSource_Read_Happy exercises GetAllExporterGroupDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllExporterGroupDataSource_Read_Happy(t *testing.T) {
	r := &GetAllExporterGroupDataSource{client: newMockClientStatus(t, 200, "{\"appsExporterGroups\":[]}")}
	m := GetAllExporterGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllExporterGroupDataSource_Read_NilClient exercises GetAllExporterGroupDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllExporterGroupDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllExporterGroupDataSource{}
	m := GetAllExporterGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllExporterGroupDataSource_Read_BuildError exercises GetAllExporterGroupDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllExporterGroupDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllExporterGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllExporterGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllExporterGroupDataSource_Read_SendError exercises GetAllExporterGroupDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllExporterGroupDataSource_Read_SendError(t *testing.T) {
	r := &GetAllExporterGroupDataSource{client: newTransportErrorClient(t)}
	m := GetAllExporterGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllExporterGroupDataSource_Read_InvalidJSON exercises GetAllExporterGroupDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllExporterGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllExporterGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllExporterGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
