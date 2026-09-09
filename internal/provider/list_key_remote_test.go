package provider

import (
	"context"
	"testing"
)

// TestKeyListResource_List_Happy exercises KeyListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestKeyListResource_List_Happy(t *testing.T) {
	r := &KeyListResource{client: newMockClientStatus(t, 200, "{\"keys\":[]}")}
	m := KeyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestKeyListResource_List_NilClient exercises KeyListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyListResource_List_NilClient(t *testing.T) {
	r := &KeyListResource{}
	m := KeyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestKeyListResource_List_BuildError exercises KeyListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestKeyListResource_List_BuildError(t *testing.T) {
	r := &KeyListResource{client: newMalformedBaseURLClient(t)}
	m := KeyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestKeyListResource_List_SendError exercises KeyListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestKeyListResource_List_SendError(t *testing.T) {
	r := &KeyListResource{client: newTransportErrorClient(t)}
	m := KeyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestKeyListResource_List_InvalidJSON exercises KeyListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestKeyListResource_List_InvalidJSON(t *testing.T) {
	r := &KeyListResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
