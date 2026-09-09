package provider

import (
	"context"
	"testing"
)

// TestExportTargetListResource_List_Happy exercises ExportTargetListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestExportTargetListResource_List_Happy(t *testing.T) {
	r := &ExportTargetListResource{client: newMockClientStatus(t, 200, "{\"externalExportTargets\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestExportTargetListResource_List_NilClient exercises ExportTargetListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportTargetListResource_List_NilClient(t *testing.T) {
	r := &ExportTargetListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestExportTargetListResource_List_BuildError exercises ExportTargetListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestExportTargetListResource_List_BuildError(t *testing.T) {
	r := &ExportTargetListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExportTargetListResource_List_SendError exercises ExportTargetListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestExportTargetListResource_List_SendError(t *testing.T) {
	r := &ExportTargetListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestExportTargetListResource_List_InvalidJSON exercises ExportTargetListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestExportTargetListResource_List_InvalidJSON(t *testing.T) {
	r := &ExportTargetListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
