package provider

import (
	"context"
	"testing"
)

// TestMapGroupListResource_List_Happy exercises MapGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestMapGroupListResource_List_Happy(t *testing.T) {
	r := &MapGroupListResource{client: newMockClientStatus(t, 200, "{\"mapGroups\":[]}")}
	m := MapGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestMapGroupListResource_List_NilClient exercises MapGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMapGroupListResource_List_NilClient(t *testing.T) {
	r := &MapGroupListResource{}
	m := MapGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestMapGroupListResource_List_BuildError exercises MapGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestMapGroupListResource_List_BuildError(t *testing.T) {
	r := &MapGroupListResource{client: newMalformedBaseURLClient(t)}
	m := MapGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapGroupListResource_List_SendError exercises MapGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestMapGroupListResource_List_SendError(t *testing.T) {
	r := &MapGroupListResource{client: newTransportErrorClient(t)}
	m := MapGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestMapGroupListResource_List_InvalidJSON exercises MapGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestMapGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &MapGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := MapGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
