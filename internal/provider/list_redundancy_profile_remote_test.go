package provider

import (
	"context"
	"testing"
)

// TestRedundancyProfileListResource_List_Happy exercises RedundancyProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestRedundancyProfileListResource_List_Happy(t *testing.T) {
	r := &RedundancyProfileListResource{client: newMockClientStatus(t, 200, "{\"redundancyProfiles\":[]}")}
	m := RedundancyProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestRedundancyProfileListResource_List_NilClient exercises RedundancyProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedundancyProfileListResource_List_NilClient(t *testing.T) {
	r := &RedundancyProfileListResource{}
	m := RedundancyProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestRedundancyProfileListResource_List_BuildError exercises RedundancyProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestRedundancyProfileListResource_List_BuildError(t *testing.T) {
	r := &RedundancyProfileListResource{client: newMalformedBaseURLClient(t)}
	m := RedundancyProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRedundancyProfileListResource_List_SendError exercises RedundancyProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestRedundancyProfileListResource_List_SendError(t *testing.T) {
	r := &RedundancyProfileListResource{client: newTransportErrorClient(t)}
	m := RedundancyProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestRedundancyProfileListResource_List_InvalidJSON exercises RedundancyProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestRedundancyProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &RedundancyProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := RedundancyProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
