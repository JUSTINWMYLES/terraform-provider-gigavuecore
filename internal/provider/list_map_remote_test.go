package provider

import (
	"context"
	"testing"
)

// TestMapListResource_List_Happy exercises MapListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMapListResource_List_Happy(t *testing.T) {
	r := &MapListResource{client: newMockClientStatus(t, 200, "{\"maps\":[]}")}
	m := MapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestMapListResource_List_NilClient exercises MapListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapListResource_List_NilClient(t *testing.T) {
	r := &MapListResource{}
	m := MapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMapListResource_List_BuildError exercises MapListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMapListResource_List_BuildError(t *testing.T) {
	r := &MapListResource{client: newMalformedBaseURLClient(t)}
	m := MapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapListResource_List_SendError exercises MapListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMapListResource_List_SendError(t *testing.T) {
	r := &MapListResource{client: newTransportErrorClient(t)}
	m := MapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapListResource_List_InvalidJSON exercises MapListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMapListResource_List_InvalidJSON(t *testing.T) {
	r := &MapListResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
