package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityControlNodeConfigsDataSource_Read_Happy exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityControlNodeConfigsDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_NilClient exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityControlNodeConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_BuildError exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityControlNodeConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_SendError exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityControlNodeConfigsDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_NotFound exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityControlNodeConfigsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_APIError exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityControlNodeConfigsDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_control_node_configs")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_APIErrorReadBody exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityControlNodeConfigsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityControlNodeConfigsDataSource_Read_InvalidJSON exercises LoadMobilityControlNodeConfigsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityControlNodeConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityControlNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityControlNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
