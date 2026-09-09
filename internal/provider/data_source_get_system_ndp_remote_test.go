package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemNdpDataSource_Read_Happy exercises GetSystemNdpDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemNdpDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemNdpDataSource_Read_NilClient exercises GetSystemNdpDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemNdpDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemNdpDataSource{}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemNdpDataSource_Read_BuildError exercises GetSystemNdpDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemNdpDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemNdpDataSource_Read_SendError exercises GetSystemNdpDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemNdpDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newTransportErrorClient(t)}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemNdpDataSource_Read_NotFound exercises GetSystemNdpDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemNdpDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemNdpDataSource_Read_APIError exercises GetSystemNdpDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemNdpDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_ndp")
}

// TestGetSystemNdpDataSource_Read_APIErrorReadBody exercises GetSystemNdpDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemNdpDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemNdpDataSource_Read_InvalidJSON exercises GetSystemNdpDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemNdpDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemNdpDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemNdpDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
