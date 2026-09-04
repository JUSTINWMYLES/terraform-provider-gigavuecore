package provider

import (
	"context"
	"testing"
)

// TestKeyMapListResource_List_Happy exercises KeyMapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestKeyMapListResource_List_Happy(t *testing.T) {
	r := &KeyMapListResource{client: newMockClientStatus(t, 200, "{\"keyMaps\":[]}")}
	m := KeyMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestKeyMapListResource_List_NilClient exercises KeyMapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestKeyMapListResource_List_NilClient(t *testing.T) {
	r := &KeyMapListResource{}
	m := KeyMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestKeyMapListResource_List_BuildError exercises KeyMapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestKeyMapListResource_List_BuildError(t *testing.T) {
	r := &KeyMapListResource{client: newMalformedBaseURLClient(t)}
	m := KeyMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestKeyMapListResource_List_SendError exercises KeyMapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestKeyMapListResource_List_SendError(t *testing.T) {
	r := &KeyMapListResource{client: newTransportErrorClient(t)}
	m := KeyMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestKeyMapListResource_List_InvalidJSON exercises KeyMapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestKeyMapListResource_List_InvalidJSON(t *testing.T) {
	r := &KeyMapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := KeyMapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
