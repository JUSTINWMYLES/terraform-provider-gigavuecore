package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppInfoDataSource_Read_Happy exercises GetAppInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAppInfoDataSource_Read_Happy(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAppInfoDataSource_Read_NilClient exercises GetAppInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAppInfoDataSource_Read_NilClient(t *testing.T) {
	r := &GetAppInfoDataSource{}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAppInfoDataSource_Read_BuildError exercises GetAppInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAppInfoDataSource_Read_BuildError(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAppInfoDataSource_Read_SendError exercises GetAppInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAppInfoDataSource_Read_SendError(t *testing.T) {
	r := &GetAppInfoDataSource{client: newTransportErrorClient(t)}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAppInfoDataSource_Read_NotFound exercises GetAppInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAppInfoDataSource_Read_NotFound(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAppInfoDataSource_Read_APIError exercises GetAppInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAppInfoDataSource_Read_APIError(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_app_info")
}

// TestGetAppInfoDataSource_Read_APIErrorReadBody exercises GetAppInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAppInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAppInfoDataSource_Read_InvalidJSON exercises GetAppInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAppInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAppInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAppInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
