package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeVersionDataSource_Read_Happy exercises GetNodeVersionDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetNodeVersionDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeVersionDataSource_Read_NilClient exercises GetNodeVersionDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeVersionDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeVersionDataSource{}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeVersionDataSource_Read_BuildError exercises GetNodeVersionDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetNodeVersionDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetNodeVersionDataSource_Read_SendError exercises GetNodeVersionDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetNodeVersionDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newTransportErrorClient(t)}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetNodeVersionDataSource_Read_NotFound exercises GetNodeVersionDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetNodeVersionDataSource_Read_NotFound(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetNodeVersionDataSource_Read_APIError exercises GetNodeVersionDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetNodeVersionDataSource_Read_APIError(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_node_version")
}

// TestGetNodeVersionDataSource_Read_APIErrorReadBody exercises GetNodeVersionDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetNodeVersionDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetNodeVersionDataSource_Read_InvalidJSON exercises GetNodeVersionDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetNodeVersionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeVersionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
