package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizConfigDataSource_Read_Happy exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadTopologyVizConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadTopologyVizConfigDataSource_Read_NilClient exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadTopologyVizConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadTopologyVizConfigDataSource_Read_BuildError exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadTopologyVizConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadTopologyVizConfigDataSource_Read_SendError exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadTopologyVizConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadTopologyVizConfigDataSource_Read_NotFound exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadTopologyVizConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadTopologyVizConfigDataSource_Read_APIError exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadTopologyVizConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_topology_viz_config")
}

// TestLoadTopologyVizConfigDataSource_Read_APIErrorReadBody exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadTopologyVizConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadTopologyVizConfigDataSource_Read_InvalidJSON exercises LoadTopologyVizConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadTopologyVizConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadTopologyVizConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadTopologyVizConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
