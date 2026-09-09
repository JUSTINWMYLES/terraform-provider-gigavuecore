package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetL2CircuitResourceConfigDataSource_Read_Happy exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetL2CircuitResourceConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetL2CircuitResourceConfigDataSource_Read_NilClient exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetL2CircuitResourceConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetL2CircuitResourceConfigDataSource_Read_BuildError exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetL2CircuitResourceConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetL2CircuitResourceConfigDataSource_Read_SendError exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetL2CircuitResourceConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newTransportErrorClient(t)}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetL2CircuitResourceConfigDataSource_Read_NotFound exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetL2CircuitResourceConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetL2CircuitResourceConfigDataSource_Read_APIError exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetL2CircuitResourceConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_l2_circuit_resource_config")
}

// TestGetL2CircuitResourceConfigDataSource_Read_APIErrorReadBody exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetL2CircuitResourceConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetL2CircuitResourceConfigDataSource_Read_InvalidJSON exercises GetL2CircuitResourceConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetL2CircuitResourceConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetL2CircuitResourceConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetL2CircuitResourceConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
