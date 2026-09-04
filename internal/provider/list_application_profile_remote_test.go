package provider

import (
	"context"
	"testing"
)

// TestApplicationProfileListResource_List_Happy exercises ApplicationProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestApplicationProfileListResource_List_Happy(t *testing.T) {
	r := &ApplicationProfileListResource{client: newMockClientStatus(t, 200, "{\"applicationProfiles\":[]}")}
	m := ApplicationProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestApplicationProfileListResource_List_NilClient exercises ApplicationProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplicationProfileListResource_List_NilClient(t *testing.T) {
	r := &ApplicationProfileListResource{}
	m := ApplicationProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestApplicationProfileListResource_List_BuildError exercises ApplicationProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestApplicationProfileListResource_List_BuildError(t *testing.T) {
	r := &ApplicationProfileListResource{client: newMalformedBaseURLClient(t)}
	m := ApplicationProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestApplicationProfileListResource_List_SendError exercises ApplicationProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestApplicationProfileListResource_List_SendError(t *testing.T) {
	r := &ApplicationProfileListResource{client: newTransportErrorClient(t)}
	m := ApplicationProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestApplicationProfileListResource_List_InvalidJSON exercises ApplicationProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestApplicationProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &ApplicationProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ApplicationProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
