package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFmHaRoleDataSource_Read_Happy exercises GetFmHaRoleDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFmHaRoleDataSource_Read_Happy(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFmHaRoleDataSource_Read_NilClient exercises GetFmHaRoleDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFmHaRoleDataSource_Read_NilClient(t *testing.T) {
	r := &GetFmHaRoleDataSource{}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFmHaRoleDataSource_Read_BuildError exercises GetFmHaRoleDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFmHaRoleDataSource_Read_BuildError(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFmHaRoleDataSource_Read_SendError exercises GetFmHaRoleDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFmHaRoleDataSource_Read_SendError(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newTransportErrorClient(t)}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFmHaRoleDataSource_Read_NotFound exercises GetFmHaRoleDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFmHaRoleDataSource_Read_NotFound(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFmHaRoleDataSource_Read_APIError exercises GetFmHaRoleDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFmHaRoleDataSource_Read_APIError(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_fm_ha_role")
}

// TestGetFmHaRoleDataSource_Read_APIErrorReadBody exercises GetFmHaRoleDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFmHaRoleDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFmHaRoleDataSource_Read_InvalidJSON exercises GetFmHaRoleDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFmHaRoleDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFmHaRoleDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFmHaRoleDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
