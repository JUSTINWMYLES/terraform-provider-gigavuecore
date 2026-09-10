package provider

import (
	"context"
	"testing"
)

// TestArchiveServerListResource_List_Happy exercises ArchiveServerListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestArchiveServerListResource_List_Happy(t *testing.T) {
	r := &ArchiveServerListResource{client: newMockClientStatus(t, 200, "{\"fmBackupArchiveServers\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestArchiveServerListResource_List_NilClient exercises ArchiveServerListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestArchiveServerListResource_List_NilClient(t *testing.T) {
	r := &ArchiveServerListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestArchiveServerListResource_List_BuildError exercises ArchiveServerListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestArchiveServerListResource_List_BuildError(t *testing.T) {
	r := &ArchiveServerListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestArchiveServerListResource_List_SendError exercises ArchiveServerListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestArchiveServerListResource_List_SendError(t *testing.T) {
	r := &ArchiveServerListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestArchiveServerListResource_List_InvalidJSON exercises ArchiveServerListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestArchiveServerListResource_List_InvalidJSON(t *testing.T) {
	r := &ArchiveServerListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
