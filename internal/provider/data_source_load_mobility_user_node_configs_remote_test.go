package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityUserNodeConfigsDataSource_Read_Happy exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityUserNodeConfigsDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_NilClient exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityUserNodeConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_BuildError exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityUserNodeConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_SendError exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityUserNodeConfigsDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_NotFound exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityUserNodeConfigsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_APIError exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityUserNodeConfigsDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_user_node_configs")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_APIErrorReadBody exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityUserNodeConfigsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityUserNodeConfigsDataSource_Read_InvalidJSON exercises LoadMobilityUserNodeConfigsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityUserNodeConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityUserNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityUserNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
