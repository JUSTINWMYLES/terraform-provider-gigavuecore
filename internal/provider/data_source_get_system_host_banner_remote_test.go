package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemHostBannerDataSource_Read_Happy exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemHostBannerDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemHostBannerDataSource_Read_NilClient exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemHostBannerDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemHostBannerDataSource{}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemHostBannerDataSource_Read_BuildError exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemHostBannerDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemHostBannerDataSource_Read_SendError exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemHostBannerDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newTransportErrorClient(t)}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemHostBannerDataSource_Read_NotFound exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemHostBannerDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemHostBannerDataSource_Read_APIError exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemHostBannerDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_host_banner")
}

// TestGetSystemHostBannerDataSource_Read_APIErrorReadBody exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemHostBannerDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemHostBannerDataSource_Read_InvalidJSON exercises GetSystemHostBannerDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemHostBannerDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemHostBannerDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemHostBannerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
