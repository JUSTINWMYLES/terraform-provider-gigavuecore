package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSysInfoDataSource_Read_Happy exercises GetSysInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSysInfoDataSource_Read_Happy(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSysInfoDataSource_Read_NilClient exercises GetSysInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSysInfoDataSource_Read_NilClient(t *testing.T) {
	r := &GetSysInfoDataSource{}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSysInfoDataSource_Read_BuildError exercises GetSysInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSysInfoDataSource_Read_BuildError(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSysInfoDataSource_Read_SendError exercises GetSysInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSysInfoDataSource_Read_SendError(t *testing.T) {
	r := &GetSysInfoDataSource{client: newTransportErrorClient(t)}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSysInfoDataSource_Read_NotFound exercises GetSysInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSysInfoDataSource_Read_NotFound(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSysInfoDataSource_Read_APIError exercises GetSysInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSysInfoDataSource_Read_APIError(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_sys_info")
}

// TestGetSysInfoDataSource_Read_APIErrorReadBody exercises GetSysInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSysInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSysInfoDataSource_Read_InvalidJSON exercises GetSysInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSysInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSysInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSysInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
