package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetExpiringCountDataSource_Read_Happy exercises GetExpiringCountDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetExpiringCountDataSource_Read_Happy(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetExpiringCountDataSource_Read_NilClient exercises GetExpiringCountDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetExpiringCountDataSource_Read_NilClient(t *testing.T) {
	r := &GetExpiringCountDataSource{}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetExpiringCountDataSource_Read_BuildError exercises GetExpiringCountDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetExpiringCountDataSource_Read_BuildError(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMalformedBaseURLClient(t)}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetExpiringCountDataSource_Read_SendError exercises GetExpiringCountDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetExpiringCountDataSource_Read_SendError(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newTransportErrorClient(t)}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetExpiringCountDataSource_Read_NotFound exercises GetExpiringCountDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetExpiringCountDataSource_Read_NotFound(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetExpiringCountDataSource_Read_APIError exercises GetExpiringCountDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetExpiringCountDataSource_Read_APIError(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_expiring_count")
}

// TestGetExpiringCountDataSource_Read_APIErrorReadBody exercises GetExpiringCountDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetExpiringCountDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetExpiringCountDataSource_Read_InvalidJSON exercises GetExpiringCountDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetExpiringCountDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetExpiringCountDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetExpiringCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
