package provider

import (
	"context"
	"testing"
)

// TestSffpProfileListResource_List_Happy exercises SffpProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestSffpProfileListResource_List_Happy(t *testing.T) {
	r := &SffpProfileListResource{client: newMockClientStatus(t, 200, "{\"sffpProfiles\":[]}")}
	m := SffpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestSffpProfileListResource_List_NilClient exercises SffpProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSffpProfileListResource_List_NilClient(t *testing.T) {
	r := &SffpProfileListResource{}
	m := SffpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestSffpProfileListResource_List_BuildError exercises SffpProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestSffpProfileListResource_List_BuildError(t *testing.T) {
	r := &SffpProfileListResource{client: newMalformedBaseURLClient(t)}
	m := SffpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSffpProfileListResource_List_SendError exercises SffpProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestSffpProfileListResource_List_SendError(t *testing.T) {
	r := &SffpProfileListResource{client: newTransportErrorClient(t)}
	m := SffpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestSffpProfileListResource_List_InvalidJSON exercises SffpProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestSffpProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &SffpProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := SffpProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
