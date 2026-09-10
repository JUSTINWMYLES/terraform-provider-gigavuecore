package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapMigrationResultsDataSource_Read_Happy exercises GetMapMigrationResultsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMapMigrationResultsDataSource_Read_Happy(t *testing.T) {
	r := &GetMapMigrationResultsDataSource{client: newMockClientStatus(t, 200, "{\"mapMigrationResults\":[]}")}
	m := GetMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMapMigrationResultsDataSource_Read_NilClient exercises GetMapMigrationResultsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMapMigrationResultsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMapMigrationResultsDataSource{}
	m := GetMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMapMigrationResultsDataSource_Read_BuildError exercises GetMapMigrationResultsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMapMigrationResultsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMapMigrationResultsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMapMigrationResultsDataSource_Read_SendError exercises GetMapMigrationResultsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMapMigrationResultsDataSource_Read_SendError(t *testing.T) {
	r := &GetMapMigrationResultsDataSource{client: newTransportErrorClient(t)}
	m := GetMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMapMigrationResultsDataSource_Read_InvalidJSON exercises GetMapMigrationResultsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMapMigrationResultsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMapMigrationResultsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMapMigrationResultsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
