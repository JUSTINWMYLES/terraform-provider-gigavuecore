package provider

import (
	"context"
	"testing"
)

// TestGetSystemDiagPsuListResource_List_Happy exercises GetSystemDiagPsuListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemDiagPsuListResource_List_Happy(t *testing.T) {
	r := &GetSystemDiagPsuListResource{client: newMockClientStatus(t, 200, "{\"systemPsuDiagDetails\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestGetSystemDiagPsuListResource_List_NilClient exercises GetSystemDiagPsuListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemDiagPsuListResource_List_NilClient(t *testing.T) {
	r := &GetSystemDiagPsuListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetSystemDiagPsuListResource_List_BuildError exercises GetSystemDiagPsuListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemDiagPsuListResource_List_BuildError(t *testing.T) {
	r := &GetSystemDiagPsuListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemDiagPsuListResource_List_SendError exercises GetSystemDiagPsuListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemDiagPsuListResource_List_SendError(t *testing.T) {
	r := &GetSystemDiagPsuListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemDiagPsuListResource_List_InvalidJSON exercises GetSystemDiagPsuListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemDiagPsuListResource_List_InvalidJSON(t *testing.T) {
	r := &GetSystemDiagPsuListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
