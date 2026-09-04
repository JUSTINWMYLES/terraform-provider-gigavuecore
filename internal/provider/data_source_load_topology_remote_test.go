package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyDataSource_Read_Happy exercises LoadTopologyDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadTopologyDataSource_Read_Happy(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadTopologyDataSource_Read_NilClient exercises LoadTopologyDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadTopologyDataSource_Read_NilClient(t *testing.T) {
	r := &LoadTopologyDataSource{}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadTopologyDataSource_Read_BuildError exercises LoadTopologyDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadTopologyDataSource_Read_BuildError(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadTopologyDataSource_Read_SendError exercises LoadTopologyDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadTopologyDataSource_Read_SendError(t *testing.T) {
	r := &LoadTopologyDataSource{client: newTransportErrorClient(t)}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadTopologyDataSource_Read_NotFound exercises LoadTopologyDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadTopologyDataSource_Read_NotFound(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadTopologyDataSource_Read_APIError exercises LoadTopologyDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadTopologyDataSource_Read_APIError(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_topology")
}

// TestLoadTopologyDataSource_Read_APIErrorReadBody exercises LoadTopologyDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadTopologyDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadTopologyDataSource_Read_InvalidJSON exercises LoadTopologyDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadTopologyDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadTopologyDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
