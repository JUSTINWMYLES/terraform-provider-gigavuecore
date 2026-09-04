package provider

import (
	"context"
	"testing"
)

// TestLoadAuditLogEntriesListResource_List_Happy exercises LoadAuditLogEntriesListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAuditLogEntriesListResource_List_Happy(t *testing.T) {
	r := &LoadAuditLogEntriesListResource{client: newMockClientStatus(t, 200, "{\"auditLogEntries\":[]}")}
	m := LoadAuditLogEntriesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLoadAuditLogEntriesListResource_List_NilClient exercises LoadAuditLogEntriesListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAuditLogEntriesListResource_List_NilClient(t *testing.T) {
	r := &LoadAuditLogEntriesListResource{}
	m := LoadAuditLogEntriesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLoadAuditLogEntriesListResource_List_BuildError exercises LoadAuditLogEntriesListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAuditLogEntriesListResource_List_BuildError(t *testing.T) {
	r := &LoadAuditLogEntriesListResource{client: newMalformedBaseURLClient(t)}
	m := LoadAuditLogEntriesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAuditLogEntriesListResource_List_SendError exercises LoadAuditLogEntriesListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAuditLogEntriesListResource_List_SendError(t *testing.T) {
	r := &LoadAuditLogEntriesListResource{client: newTransportErrorClient(t)}
	m := LoadAuditLogEntriesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLoadAuditLogEntriesListResource_List_InvalidJSON exercises LoadAuditLogEntriesListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAuditLogEntriesListResource_List_InvalidJSON(t *testing.T) {
	r := &LoadAuditLogEntriesListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAuditLogEntriesListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
