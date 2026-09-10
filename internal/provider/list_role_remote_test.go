package provider

import (
	"context"
	"testing"
)

// TestRoleListResource_List_Happy exercises RoleListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestRoleListResource_List_Happy(t *testing.T) {
	r := &RoleListResource{client: newMockClientStatus(t, 200, "{\"roles\":[]}")}
	m := RoleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestRoleListResource_List_NilClient exercises RoleListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRoleListResource_List_NilClient(t *testing.T) {
	r := &RoleListResource{}
	m := RoleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestRoleListResource_List_BuildError exercises RoleListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestRoleListResource_List_BuildError(t *testing.T) {
	r := &RoleListResource{client: newMalformedBaseURLClient(t)}
	m := RoleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRoleListResource_List_SendError exercises RoleListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestRoleListResource_List_SendError(t *testing.T) {
	r := &RoleListResource{client: newTransportErrorClient(t)}
	m := RoleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRoleListResource_List_InvalidJSON exercises RoleListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestRoleListResource_List_InvalidJSON(t *testing.T) {
	r := &RoleListResource{client: newMockClientStatus(t, 200, "{{")}
	m := RoleListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
