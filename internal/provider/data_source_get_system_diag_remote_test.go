package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDiagDataSource_Read_Happy exercises GetSystemDiagDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemDiagDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemDiagDataSource_Read_NilClient exercises GetSystemDiagDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemDiagDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemDiagDataSource{}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemDiagDataSource_Read_BuildError exercises GetSystemDiagDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemDiagDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemDiagDataSource_Read_SendError exercises GetSystemDiagDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemDiagDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newTransportErrorClient(t)}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemDiagDataSource_Read_NotFound exercises GetSystemDiagDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemDiagDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemDiagDataSource_Read_APIError exercises GetSystemDiagDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemDiagDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_system_diag")
}

// TestGetSystemDiagDataSource_Read_APIErrorReadBody exercises GetSystemDiagDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemDiagDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemDiagDataSource_Read_InvalidJSON exercises GetSystemDiagDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemDiagDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemDiagDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemDiagDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
