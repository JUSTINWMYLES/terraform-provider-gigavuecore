package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSystemTimeDataSource_Read_Happy exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetNodeSystemTimeDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeSystemTimeDataSource_Read_NilClient exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeSystemTimeDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeSystemTimeDataSource_Read_BuildError exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetNodeSystemTimeDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetNodeSystemTimeDataSource_Read_SendError exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetNodeSystemTimeDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newTransportErrorClient(t)}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetNodeSystemTimeDataSource_Read_NotFound exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetNodeSystemTimeDataSource_Read_NotFound(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetNodeSystemTimeDataSource_Read_APIError exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetNodeSystemTimeDataSource_Read_APIError(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_node_system_time")
}

// TestGetNodeSystemTimeDataSource_Read_APIErrorReadBody exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetNodeSystemTimeDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetNodeSystemTimeDataSource_Read_InvalidJSON exercises GetNodeSystemTimeDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetNodeSystemTimeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeSystemTimeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeSystemTimeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
