package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListConnectionsDataSource_Read_Happy exercises ListConnectionsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestListConnectionsDataSource_Read_Happy(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListConnectionsDataSource_Read_NilClient exercises ListConnectionsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListConnectionsDataSource_Read_NilClient(t *testing.T) {
	r := &ListConnectionsDataSource{}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListConnectionsDataSource_Read_BuildError exercises ListConnectionsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListConnectionsDataSource_Read_BuildError(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMalformedBaseURLClient(t)}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListConnectionsDataSource_Read_SendError exercises ListConnectionsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListConnectionsDataSource_Read_SendError(t *testing.T) {
	r := &ListConnectionsDataSource{client: newTransportErrorClient(t)}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListConnectionsDataSource_Read_NotFound exercises ListConnectionsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestListConnectionsDataSource_Read_NotFound(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMockClientStatus(t, 404, "")}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestListConnectionsDataSource_Read_APIError exercises ListConnectionsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListConnectionsDataSource_Read_APIError(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_list_connections")
}

// TestListConnectionsDataSource_Read_APIErrorReadBody exercises ListConnectionsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListConnectionsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListConnectionsDataSource_Read_InvalidJSON exercises ListConnectionsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListConnectionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListConnectionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListConnectionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
