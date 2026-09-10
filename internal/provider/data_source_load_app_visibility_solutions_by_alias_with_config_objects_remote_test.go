package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_Happy exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_NilClient exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_BuildError exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_SendError exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newTransportErrorClient(t)}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_NotFound exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_APIError exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_APIError(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_APIErrorReadBody exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_InvalidJSON exercises LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
