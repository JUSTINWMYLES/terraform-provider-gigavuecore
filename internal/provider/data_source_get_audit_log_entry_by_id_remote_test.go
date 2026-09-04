package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAuditLogEntryByIdDataSource_Read_Happy exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAuditLogEntryByIdDataSource_Read_Happy(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAuditLogEntryByIdDataSource_Read_NilClient exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAuditLogEntryByIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAuditLogEntryByIdDataSource_Read_BuildError exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAuditLogEntryByIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAuditLogEntryByIdDataSource_Read_SendError exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAuditLogEntryByIdDataSource_Read_SendError(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newTransportErrorClient(t)}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAuditLogEntryByIdDataSource_Read_NotFound exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAuditLogEntryByIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAuditLogEntryByIdDataSource_Read_APIError exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAuditLogEntryByIdDataSource_Read_APIError(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_audit_log_entry_by_id")
}

// TestGetAuditLogEntryByIdDataSource_Read_APIErrorReadBody exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAuditLogEntryByIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAuditLogEntryByIdDataSource_Read_InvalidJSON exercises GetAuditLogEntryByIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAuditLogEntryByIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAuditLogEntryByIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAuditLogEntryByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
