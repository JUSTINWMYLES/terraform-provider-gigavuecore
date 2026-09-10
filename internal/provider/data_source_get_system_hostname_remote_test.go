package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemHostnameDataSource_Read_Happy exercises GetSystemHostnameDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemHostnameDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemHostnameDataSource_Read_NilClient exercises GetSystemHostnameDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemHostnameDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemHostnameDataSource{}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemHostnameDataSource_Read_BuildError exercises GetSystemHostnameDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemHostnameDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemHostnameDataSource_Read_SendError exercises GetSystemHostnameDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemHostnameDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newTransportErrorClient(t)}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemHostnameDataSource_Read_NotFound exercises GetSystemHostnameDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemHostnameDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemHostnameDataSource_Read_APIError exercises GetSystemHostnameDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemHostnameDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_hostname")
}

// TestGetSystemHostnameDataSource_Read_APIErrorReadBody exercises GetSystemHostnameDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemHostnameDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemHostnameDataSource_Read_InvalidJSON exercises GetSystemHostnameDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemHostnameDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemHostnameDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemHostnameDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
