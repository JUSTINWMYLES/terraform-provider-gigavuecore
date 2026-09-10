package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllMapMigrationResultsByAliasDataSource_Read_Happy exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_Happy(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_NilClient exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_BuildError exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_SendError exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_SendError(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newTransportErrorClient(t)}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_NotFound exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_APIError exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_APIError(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_map_migration_results_by_alias")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_APIErrorReadBody exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllMapMigrationResultsByAliasDataSource_Read_InvalidJSON exercises GetAllMapMigrationResultsByAliasDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllMapMigrationResultsByAliasDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllMapMigrationResultsByAliasDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllMapMigrationResultsByAliasDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
