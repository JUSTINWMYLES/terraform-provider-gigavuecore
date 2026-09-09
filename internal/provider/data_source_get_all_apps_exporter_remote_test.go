package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAppsExporterDataSource_Read_Happy exercises GetAllAppsExporterDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAppsExporterDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAppsExporterDataSource{client: newMockClientStatus(t, 200, "{\"appsExporters\":[]}")}
	m := GetAllAppsExporterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAppsExporterDataSource_Read_NilClient exercises GetAllAppsExporterDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAppsExporterDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAppsExporterDataSource{}
	m := GetAllAppsExporterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAppsExporterDataSource_Read_BuildError exercises GetAllAppsExporterDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAppsExporterDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAppsExporterDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAppsExporterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppsExporterDataSource_Read_SendError exercises GetAllAppsExporterDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAppsExporterDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAppsExporterDataSource{client: newTransportErrorClient(t)}
	m := GetAllAppsExporterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppsExporterDataSource_Read_InvalidJSON exercises GetAllAppsExporterDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAppsExporterDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAppsExporterDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAppsExporterDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
