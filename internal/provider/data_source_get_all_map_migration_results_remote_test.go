package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllMapMigrationResultsDataSource_Read_Happy exercises GetAllMapMigrationResultsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllMapMigrationResultsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllMapMigrationResultsDataSource{client: newMockClientStatus(t, 200, "{\"mapMigrationResults\":[]}")}
	m := GetAllMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllMapMigrationResultsDataSource_Read_NilClient exercises GetAllMapMigrationResultsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllMapMigrationResultsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllMapMigrationResultsDataSource{}
	m := GetAllMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllMapMigrationResultsDataSource_Read_BuildError exercises GetAllMapMigrationResultsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllMapMigrationResultsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllMapMigrationResultsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllMapMigrationResultsDataSource_Read_SendError exercises GetAllMapMigrationResultsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllMapMigrationResultsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllMapMigrationResultsDataSource{client: newTransportErrorClient(t)}
	m := GetAllMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllMapMigrationResultsDataSource_Read_InvalidJSON exercises GetAllMapMigrationResultsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllMapMigrationResultsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllMapMigrationResultsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
