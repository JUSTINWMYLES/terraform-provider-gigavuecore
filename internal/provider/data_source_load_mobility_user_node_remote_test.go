package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityUserNodeDataSource_Read_Happy exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityUserNodeDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityUserNodeDataSource_Read_NilClient exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityUserNodeDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityUserNodeDataSource_Read_BuildError exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityUserNodeDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityUserNodeDataSource_Read_SendError exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityUserNodeDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityUserNodeDataSource_Read_NotFound exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityUserNodeDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityUserNodeDataSource_Read_APIError exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityUserNodeDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_user_node")
}

// TestLoadMobilityUserNodeDataSource_Read_APIErrorReadBody exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityUserNodeDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityUserNodeDataSource_Read_InvalidJSON exercises LoadMobilityUserNodeDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityUserNodeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityUserNodeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityUserNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
