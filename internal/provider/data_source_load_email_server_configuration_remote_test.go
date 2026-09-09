package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadEmailServerConfigurationDataSource_Read_Happy exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadEmailServerConfigurationDataSource_Read_Happy(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadEmailServerConfigurationDataSource_Read_NilClient exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadEmailServerConfigurationDataSource_Read_NilClient(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadEmailServerConfigurationDataSource_Read_BuildError exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadEmailServerConfigurationDataSource_Read_BuildError(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadEmailServerConfigurationDataSource_Read_SendError exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadEmailServerConfigurationDataSource_Read_SendError(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newTransportErrorClient(t)}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadEmailServerConfigurationDataSource_Read_NotFound exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadEmailServerConfigurationDataSource_Read_NotFound(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadEmailServerConfigurationDataSource_Read_APIError exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadEmailServerConfigurationDataSource_Read_APIError(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_email_server_configuration")
}

// TestLoadEmailServerConfigurationDataSource_Read_APIErrorReadBody exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadEmailServerConfigurationDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadEmailServerConfigurationDataSource_Read_InvalidJSON exercises LoadEmailServerConfigurationDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadEmailServerConfigurationDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadEmailServerConfigurationDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadEmailServerConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
