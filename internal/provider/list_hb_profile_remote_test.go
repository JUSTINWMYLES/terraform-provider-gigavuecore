package provider

import (
	"context"
	"testing"
)

// TestHbProfileListResource_List_Happy exercises HbProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestHbProfileListResource_List_Happy(t *testing.T) {
	r := &HbProfileListResource{client: newMockClientStatus(t, 200, "{\"heartbeatProfiles\":[]}")}
	m := HbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestHbProfileListResource_List_NilClient exercises HbProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHbProfileListResource_List_NilClient(t *testing.T) {
	r := &HbProfileListResource{}
	m := HbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestHbProfileListResource_List_BuildError exercises HbProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestHbProfileListResource_List_BuildError(t *testing.T) {
	r := &HbProfileListResource{client: newMalformedBaseURLClient(t)}
	m := HbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHbProfileListResource_List_SendError exercises HbProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestHbProfileListResource_List_SendError(t *testing.T) {
	r := &HbProfileListResource{client: newTransportErrorClient(t)}
	m := HbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestHbProfileListResource_List_InvalidJSON exercises HbProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestHbProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &HbProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := HbProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
