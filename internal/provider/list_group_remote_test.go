package provider

import (
	"context"
	"testing"
)

// TestGroupListResource_List_Happy exercises GroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGroupListResource_List_Happy(t *testing.T) {
	r := &GroupListResource{client: newMockClientStatus(t, 200, "{\"groups\":[]}")}
	m := GroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGroupListResource_List_NilClient exercises GroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGroupListResource_List_NilClient(t *testing.T) {
	r := &GroupListResource{}
	m := GroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGroupListResource_List_BuildError exercises GroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGroupListResource_List_BuildError(t *testing.T) {
	r := &GroupListResource{client: newMalformedBaseURLClient(t)}
	m := GroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGroupListResource_List_SendError exercises GroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGroupListResource_List_SendError(t *testing.T) {
	r := &GroupListResource{client: newTransportErrorClient(t)}
	m := GroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGroupListResource_List_InvalidJSON exercises GroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &GroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
