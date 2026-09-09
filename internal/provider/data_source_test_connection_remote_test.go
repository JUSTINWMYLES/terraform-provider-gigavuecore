package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestTestConnectionDataSource_Read_Happy exercises TestConnectionDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestTestConnectionDataSource_Read_Happy(t *testing.T) {
	r := &TestConnectionDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTestConnectionDataSource_Read_NilClient exercises TestConnectionDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTestConnectionDataSource_Read_NilClient(t *testing.T) {
	r := &TestConnectionDataSource{}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTestConnectionDataSource_Read_BuildError exercises TestConnectionDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTestConnectionDataSource_Read_BuildError(t *testing.T) {
	r := &TestConnectionDataSource{client: newMalformedBaseURLClient(t)}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTestConnectionDataSource_Read_SendError exercises TestConnectionDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTestConnectionDataSource_Read_SendError(t *testing.T) {
	r := &TestConnectionDataSource{client: newTransportErrorClient(t)}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTestConnectionDataSource_Read_NotFound exercises TestConnectionDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestTestConnectionDataSource_Read_NotFound(t *testing.T) {
	r := &TestConnectionDataSource{client: newMockClientStatus(t, 404, "")}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestTestConnectionDataSource_Read_APIError exercises TestConnectionDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTestConnectionDataSource_Read_APIError(t *testing.T) {
	r := &TestConnectionDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_test_connection")
}

// TestTestConnectionDataSource_Read_APIErrorReadBody exercises TestConnectionDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTestConnectionDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &TestConnectionDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTestConnectionDataSource_Read_InvalidJSON exercises TestConnectionDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTestConnectionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &TestConnectionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := TestConnectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
