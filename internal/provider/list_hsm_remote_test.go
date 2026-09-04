package provider

import (
	"context"
	"testing"
)

// TestHsmListResource_List_Happy exercises HsmListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestHsmListResource_List_Happy(t *testing.T) {
	r := &HsmListResource{client: newMockClientStatus(t, 200, "{\"hsms\":[]}")}
	_, diags := r.listRemote(context.Background())
	requireNoErrors(t, diags)
}

// TestHsmListResource_List_NilClient exercises HsmListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHsmListResource_List_NilClient(t *testing.T) {
	r := &HsmListResource{}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestHsmListResource_List_BuildError exercises HsmListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestHsmListResource_List_BuildError(t *testing.T) {
	r := &HsmListResource{client: newMalformedBaseURLClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHsmListResource_List_SendError exercises HsmListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestHsmListResource_List_SendError(t *testing.T) {
	r := &HsmListResource{client: newTransportErrorClient(t)}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHsmListResource_List_InvalidJSON exercises HsmListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestHsmListResource_List_InvalidJSON(t *testing.T) {
	r := &HsmListResource{client: newMockClientStatus(t, 200, "{{")}
	_, diags := r.listRemote(context.Background())
	hasErrorContaining(t, diags, "Could not decode list page")
}
