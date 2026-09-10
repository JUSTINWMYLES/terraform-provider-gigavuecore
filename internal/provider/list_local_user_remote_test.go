package provider

import (
	"context"
	"testing"
)

// TestLocalUserListResource_List_Happy exercises LocalUserListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLocalUserListResource_List_Happy(t *testing.T) {
	r := &LocalUserListResource{client: newMockClientStatus(t, 200, "{\"localUsers\":[]}")}
	m := LocalUserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestLocalUserListResource_List_NilClient exercises LocalUserListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLocalUserListResource_List_NilClient(t *testing.T) {
	r := &LocalUserListResource{}
	m := LocalUserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestLocalUserListResource_List_BuildError exercises LocalUserListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLocalUserListResource_List_BuildError(t *testing.T) {
	r := &LocalUserListResource{client: newMalformedBaseURLClient(t)}
	m := LocalUserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLocalUserListResource_List_SendError exercises LocalUserListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLocalUserListResource_List_SendError(t *testing.T) {
	r := &LocalUserListResource{client: newTransportErrorClient(t)}
	m := LocalUserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestLocalUserListResource_List_InvalidJSON exercises LocalUserListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLocalUserListResource_List_InvalidJSON(t *testing.T) {
	r := &LocalUserListResource{client: newMockClientStatus(t, 200, "{{")}
	m := LocalUserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
