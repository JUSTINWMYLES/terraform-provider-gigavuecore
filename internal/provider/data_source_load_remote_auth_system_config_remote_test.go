package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadRemoteAuthSystemConfigDataSource_Read_Happy exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadRemoteAuthSystemConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_NilClient exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadRemoteAuthSystemConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_BuildError exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadRemoteAuthSystemConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_SendError exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadRemoteAuthSystemConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_NotFound exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadRemoteAuthSystemConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_APIError exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadRemoteAuthSystemConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_remote_auth_system_config")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_APIErrorReadBody exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadRemoteAuthSystemConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadRemoteAuthSystemConfigDataSource_Read_InvalidJSON exercises LoadRemoteAuthSystemConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadRemoteAuthSystemConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadRemoteAuthSystemConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadRemoteAuthSystemConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
