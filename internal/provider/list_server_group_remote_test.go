package provider

import (
	"context"
	"testing"
)

// TestServerGroupListResource_List_Happy exercises ServerGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestServerGroupListResource_List_Happy(t *testing.T) {
	r := &ServerGroupListResource{client: newMockClientStatus(t, 200, "{\"icapServerGroups\":[]}")}
	m := ServerGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestServerGroupListResource_List_NilClient exercises ServerGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServerGroupListResource_List_NilClient(t *testing.T) {
	r := &ServerGroupListResource{}
	m := ServerGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestServerGroupListResource_List_BuildError exercises ServerGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestServerGroupListResource_List_BuildError(t *testing.T) {
	r := &ServerGroupListResource{client: newMalformedBaseURLClient(t)}
	m := ServerGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestServerGroupListResource_List_SendError exercises ServerGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestServerGroupListResource_List_SendError(t *testing.T) {
	r := &ServerGroupListResource{client: newTransportErrorClient(t)}
	m := ServerGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestServerGroupListResource_List_InvalidJSON exercises ServerGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestServerGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &ServerGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ServerGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
