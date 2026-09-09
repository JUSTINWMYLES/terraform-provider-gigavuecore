package provider

import (
	"context"
	"testing"
)

// TestListFmStoredClustersConfigBackupSnapshotsListResource_List_Happy exercises ListFmStoredClustersConfigBackupSnapshotsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestListFmStoredClustersConfigBackupSnapshotsListResource_List_Happy(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsListResource{client: newMockClientStatus(t, 200, "{\"clustersConfigBackups\":[]}")}
	m := ListFmStoredClustersConfigBackupSnapshotsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestListFmStoredClustersConfigBackupSnapshotsListResource_List_NilClient exercises ListFmStoredClustersConfigBackupSnapshotsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListFmStoredClustersConfigBackupSnapshotsListResource_List_NilClient(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsListResource{}
	m := ListFmStoredClustersConfigBackupSnapshotsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestListFmStoredClustersConfigBackupSnapshotsListResource_List_BuildError exercises ListFmStoredClustersConfigBackupSnapshotsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestListFmStoredClustersConfigBackupSnapshotsListResource_List_BuildError(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsListResource{client: newMalformedBaseURLClient(t)}
	m := ListFmStoredClustersConfigBackupSnapshotsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListFmStoredClustersConfigBackupSnapshotsListResource_List_SendError exercises ListFmStoredClustersConfigBackupSnapshotsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestListFmStoredClustersConfigBackupSnapshotsListResource_List_SendError(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsListResource{client: newTransportErrorClient(t)}
	m := ListFmStoredClustersConfigBackupSnapshotsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestListFmStoredClustersConfigBackupSnapshotsListResource_List_InvalidJSON exercises ListFmStoredClustersConfigBackupSnapshotsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestListFmStoredClustersConfigBackupSnapshotsListResource_List_InvalidJSON(t *testing.T) {
	r := &ListFmStoredClustersConfigBackupSnapshotsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ListFmStoredClustersConfigBackupSnapshotsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
