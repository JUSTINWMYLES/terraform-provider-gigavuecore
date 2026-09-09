package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemWebProxyDataSource_Read_Happy exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemWebProxyDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemWebProxyDataSource_Read_NilClient exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemWebProxyDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemWebProxyDataSource{}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemWebProxyDataSource_Read_BuildError exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemWebProxyDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemWebProxyDataSource_Read_SendError exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemWebProxyDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newTransportErrorClient(t)}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemWebProxyDataSource_Read_NotFound exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemWebProxyDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemWebProxyDataSource_Read_APIError exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemWebProxyDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_web_proxy")
}

// TestGetSystemWebProxyDataSource_Read_APIErrorReadBody exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemWebProxyDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemWebProxyDataSource_Read_InvalidJSON exercises GetSystemWebProxyDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemWebProxyDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemWebProxyDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemWebProxyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
