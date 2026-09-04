package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentAllowanceDataSource_Read_Happy exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetCurrentAllowanceDataSource_Read_Happy(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCurrentAllowanceDataSource_Read_NilClient exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCurrentAllowanceDataSource_Read_NilClient(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCurrentAllowanceDataSource_Read_BuildError exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetCurrentAllowanceDataSource_Read_BuildError(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetCurrentAllowanceDataSource_Read_SendError exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetCurrentAllowanceDataSource_Read_SendError(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newTransportErrorClient(t)}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetCurrentAllowanceDataSource_Read_NotFound exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetCurrentAllowanceDataSource_Read_NotFound(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetCurrentAllowanceDataSource_Read_APIError exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetCurrentAllowanceDataSource_Read_APIError(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_current_allowance")
}

// TestGetCurrentAllowanceDataSource_Read_APIErrorReadBody exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetCurrentAllowanceDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetCurrentAllowanceDataSource_Read_InvalidJSON exercises GetCurrentAllowanceDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetCurrentAllowanceDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCurrentAllowanceDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCurrentAllowanceDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
