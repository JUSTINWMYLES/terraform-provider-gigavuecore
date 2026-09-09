package provider

import (
	"context"
	"testing"
)

// TestListBulkReplicateConfigRestoreLogsListResource_List_Happy exercises ListBulkReplicateConfigRestoreLogsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListBulkReplicateConfigRestoreLogsListResource_List_Happy(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsListResource{client: newMockClientStatus(t, 200, "{\"bulkReplicateConfigRestoreLogs\":[]}")}
	m := ListBulkReplicateConfigRestoreLogsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestListBulkReplicateConfigRestoreLogsListResource_List_NilClient exercises ListBulkReplicateConfigRestoreLogsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListBulkReplicateConfigRestoreLogsListResource_List_NilClient(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsListResource{}
	m := ListBulkReplicateConfigRestoreLogsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestListBulkReplicateConfigRestoreLogsListResource_List_BuildError exercises ListBulkReplicateConfigRestoreLogsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListBulkReplicateConfigRestoreLogsListResource_List_BuildError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsListResource{client: newMalformedBaseURLClient(t)}
	m := ListBulkReplicateConfigRestoreLogsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListBulkReplicateConfigRestoreLogsListResource_List_SendError exercises ListBulkReplicateConfigRestoreLogsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListBulkReplicateConfigRestoreLogsListResource_List_SendError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsListResource{client: newTransportErrorClient(t)}
	m := ListBulkReplicateConfigRestoreLogsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListBulkReplicateConfigRestoreLogsListResource_List_InvalidJSON exercises ListBulkReplicateConfigRestoreLogsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListBulkReplicateConfigRestoreLogsListResource_List_InvalidJSON(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ListBulkReplicateConfigRestoreLogsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
