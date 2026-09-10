package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAuditLogEntriesDataSource_Read_Happy exercises LoadAuditLogEntriesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAuditLogEntriesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAuditLogEntriesDataSource{client: newMockClientStatus(t, 200, "{\"auditLogEntries\":[]}")}
	m := LoadAuditLogEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAuditLogEntriesDataSource_Read_NilClient exercises LoadAuditLogEntriesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAuditLogEntriesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAuditLogEntriesDataSource{}
	m := LoadAuditLogEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAuditLogEntriesDataSource_Read_BuildError exercises LoadAuditLogEntriesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAuditLogEntriesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAuditLogEntriesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAuditLogEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAuditLogEntriesDataSource_Read_SendError exercises LoadAuditLogEntriesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAuditLogEntriesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAuditLogEntriesDataSource{client: newTransportErrorClient(t)}
	m := LoadAuditLogEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAuditLogEntriesDataSource_Read_InvalidJSON exercises LoadAuditLogEntriesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAuditLogEntriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAuditLogEntriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAuditLogEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
