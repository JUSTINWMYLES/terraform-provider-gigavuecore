package provider

import (
	"context"
	"testing"
)

// TestGsGroupListResource_List_Happy exercises GsGroupListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGsGroupListResource_List_Happy(t *testing.T) {
	r := &GsGroupListResource{client: newMockClientStatus(t, 200, "{\"gsGroups\":[]}")}
	m := GsGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGsGroupListResource_List_NilClient exercises GsGroupListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsGroupListResource_List_NilClient(t *testing.T) {
	r := &GsGroupListResource{}
	m := GsGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGsGroupListResource_List_BuildError exercises GsGroupListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGsGroupListResource_List_BuildError(t *testing.T) {
	r := &GsGroupListResource{client: newMalformedBaseURLClient(t)}
	m := GsGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGsGroupListResource_List_SendError exercises GsGroupListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGsGroupListResource_List_SendError(t *testing.T) {
	r := &GsGroupListResource{client: newTransportErrorClient(t)}
	m := GsGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGsGroupListResource_List_InvalidJSON exercises GsGroupListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGsGroupListResource_List_InvalidJSON(t *testing.T) {
	r := &GsGroupListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsGroupListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
