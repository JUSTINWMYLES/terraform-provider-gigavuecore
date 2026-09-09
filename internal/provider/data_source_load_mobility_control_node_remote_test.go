package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityControlNodeDataSource_Read_Happy exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilityControlNodeDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilityControlNodeDataSource_Read_NilClient exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilityControlNodeDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilityControlNodeDataSource_Read_BuildError exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilityControlNodeDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilityControlNodeDataSource_Read_SendError exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilityControlNodeDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilityControlNodeDataSource_Read_NotFound exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilityControlNodeDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilityControlNodeDataSource_Read_APIError exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilityControlNodeDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_control_node")
}

// TestLoadMobilityControlNodeDataSource_Read_APIErrorReadBody exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilityControlNodeDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilityControlNodeDataSource_Read_InvalidJSON exercises LoadMobilityControlNodeDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilityControlNodeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilityControlNodeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilityControlNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
