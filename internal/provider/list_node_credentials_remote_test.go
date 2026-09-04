package provider

import (
	"context"
	"testing"
)

// TestNodeCredentialsListResource_List_Happy exercises NodeCredentialsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestNodeCredentialsListResource_List_Happy(t *testing.T) {
	r := &NodeCredentialsListResource{client: newMockClientStatus(t, 200, "{\"devCredsList\":[]}")}
	m := NodeCredentialsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestNodeCredentialsListResource_List_NilClient exercises NodeCredentialsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNodeCredentialsListResource_List_NilClient(t *testing.T) {
	r := &NodeCredentialsListResource{}
	m := NodeCredentialsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestNodeCredentialsListResource_List_BuildError exercises NodeCredentialsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestNodeCredentialsListResource_List_BuildError(t *testing.T) {
	r := &NodeCredentialsListResource{client: newMalformedBaseURLClient(t)}
	m := NodeCredentialsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNodeCredentialsListResource_List_SendError exercises NodeCredentialsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestNodeCredentialsListResource_List_SendError(t *testing.T) {
	r := &NodeCredentialsListResource{client: newTransportErrorClient(t)}
	m := NodeCredentialsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestNodeCredentialsListResource_List_InvalidJSON exercises NodeCredentialsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestNodeCredentialsListResource_List_InvalidJSON(t *testing.T) {
	r := &NodeCredentialsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := NodeCredentialsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
