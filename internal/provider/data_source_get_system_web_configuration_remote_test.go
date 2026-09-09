package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemWebConfigurationDataSource_Read_Happy exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemWebConfigurationDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemWebConfigurationDataSource_Read_NilClient exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemWebConfigurationDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemWebConfigurationDataSource_Read_BuildError exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemWebConfigurationDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemWebConfigurationDataSource_Read_SendError exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemWebConfigurationDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newTransportErrorClient(t)}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemWebConfigurationDataSource_Read_NotFound exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemWebConfigurationDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemWebConfigurationDataSource_Read_APIError exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemWebConfigurationDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_web_configuration")
}

// TestGetSystemWebConfigurationDataSource_Read_APIErrorReadBody exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemWebConfigurationDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemWebConfigurationDataSource_Read_InvalidJSON exercises GetSystemWebConfigurationDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemWebConfigurationDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemWebConfigurationDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemWebConfigurationDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
