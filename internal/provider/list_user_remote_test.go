package provider

import (
	"context"
	"testing"
)

// TestUserListResource_List_Happy exercises UserListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestUserListResource_List_Happy(t *testing.T) {
	r := &UserListResource{client: newMockClientStatus(t, 200, "{\"users\":[]}")}
	m := UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestUserListResource_List_NilClient exercises UserListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUserListResource_List_NilClient(t *testing.T) {
	r := &UserListResource{}
	m := UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestUserListResource_List_BuildError exercises UserListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestUserListResource_List_BuildError(t *testing.T) {
	r := &UserListResource{client: newMalformedBaseURLClient(t)}
	m := UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestUserListResource_List_SendError exercises UserListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestUserListResource_List_SendError(t *testing.T) {
	r := &UserListResource{client: newTransportErrorClient(t)}
	m := UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestUserListResource_List_InvalidJSON exercises UserListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestUserListResource_List_InvalidJSON(t *testing.T) {
	r := &UserListResource{client: newMockClientStatus(t, 200, "{{")}
	m := UserListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
