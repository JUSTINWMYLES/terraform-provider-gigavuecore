package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSshSupportedParametersDataSource_Read_Happy exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadSshSupportedParametersDataSource_Read_Happy(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSshSupportedParametersDataSource_Read_NilClient exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSshSupportedParametersDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSshSupportedParametersDataSource_Read_BuildError exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSshSupportedParametersDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSshSupportedParametersDataSource_Read_SendError exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSshSupportedParametersDataSource_Read_SendError(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newTransportErrorClient(t)}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSshSupportedParametersDataSource_Read_NotFound exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadSshSupportedParametersDataSource_Read_NotFound(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadSshSupportedParametersDataSource_Read_APIError exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSshSupportedParametersDataSource_Read_APIError(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_ssh_supported_parameters")
}

// TestLoadSshSupportedParametersDataSource_Read_APIErrorReadBody exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSshSupportedParametersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadSshSupportedParametersDataSource_Read_InvalidJSON exercises LoadSshSupportedParametersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadSshSupportedParametersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSshSupportedParametersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSshSupportedParametersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
