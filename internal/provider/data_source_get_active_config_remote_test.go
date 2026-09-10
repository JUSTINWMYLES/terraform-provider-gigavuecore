package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetActiveConfigDataSource_Read_Happy exercises GetActiveConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetActiveConfigDataSource_Read_Happy(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetActiveConfigDataSource_Read_NilClient exercises GetActiveConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetActiveConfigDataSource_Read_NilClient(t *testing.T) {
	r := &GetActiveConfigDataSource{}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetActiveConfigDataSource_Read_BuildError exercises GetActiveConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetActiveConfigDataSource_Read_BuildError(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetActiveConfigDataSource_Read_SendError exercises GetActiveConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetActiveConfigDataSource_Read_SendError(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newTransportErrorClient(t)}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetActiveConfigDataSource_Read_NotFound exercises GetActiveConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetActiveConfigDataSource_Read_NotFound(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetActiveConfigDataSource_Read_APIError exercises GetActiveConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetActiveConfigDataSource_Read_APIError(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_active_config")
}

// TestGetActiveConfigDataSource_Read_APIErrorReadBody exercises GetActiveConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetActiveConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetActiveConfigDataSource_Read_InvalidJSON exercises GetActiveConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetActiveConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetActiveConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetActiveConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
