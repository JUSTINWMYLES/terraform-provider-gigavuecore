package provider

import (
	"context"
	"testing"
)

// TestInlineSslProfileListResource_List_Happy exercises InlineSslProfileListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestInlineSslProfileListResource_List_Happy(t *testing.T) {
	r := &InlineSslProfileListResource{client: newMockClientStatus(t, 200, "{\"inlineSslProfiles\":[]}")}
	m := InlineSslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestInlineSslProfileListResource_List_NilClient exercises InlineSslProfileListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslProfileListResource_List_NilClient(t *testing.T) {
	r := &InlineSslProfileListResource{}
	m := InlineSslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestInlineSslProfileListResource_List_BuildError exercises InlineSslProfileListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestInlineSslProfileListResource_List_BuildError(t *testing.T) {
	r := &InlineSslProfileListResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineSslProfileListResource_List_SendError exercises InlineSslProfileListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestInlineSslProfileListResource_List_SendError(t *testing.T) {
	r := &InlineSslProfileListResource{client: newTransportErrorClient(t)}
	m := InlineSslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineSslProfileListResource_List_InvalidJSON exercises InlineSslProfileListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestInlineSslProfileListResource_List_InvalidJSON(t *testing.T) {
	r := &InlineSslProfileListResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslProfileListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
