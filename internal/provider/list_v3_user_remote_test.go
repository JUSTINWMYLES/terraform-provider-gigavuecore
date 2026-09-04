package provider

import (
	"context"
	"testing"
)

// TestV3UserListResource_List_Happy exercises V3UserListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestV3UserListResource_List_Happy(t *testing.T) {
	r := &V3UserListResource{client: newMockClientStatus(t, 200, "{\"v3Users\":[]}")}
	m := V3UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestV3UserListResource_List_NilClient exercises V3UserListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestV3UserListResource_List_NilClient(t *testing.T) {
	r := &V3UserListResource{}
	m := V3UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestV3UserListResource_List_BuildError exercises V3UserListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestV3UserListResource_List_BuildError(t *testing.T) {
	r := &V3UserListResource{client: newMalformedBaseURLClient(t)}
	m := V3UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestV3UserListResource_List_SendError exercises V3UserListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestV3UserListResource_List_SendError(t *testing.T) {
	r := &V3UserListResource{client: newTransportErrorClient(t)}
	m := V3UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestV3UserListResource_List_InvalidJSON exercises V3UserListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestV3UserListResource_List_InvalidJSON(t *testing.T) {
	r := &V3UserListResource{client: newMockClientStatus(t, 200, "{{")}
	m := V3UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
