package provider

import (
	"context"
	"testing"
)

// TestInlineSslAppListResource_List_Happy exercises InlineSslAppListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestInlineSslAppListResource_List_Happy(t *testing.T) {
	r := &InlineSslAppListResource{client: newMockClientStatus(t, 200, "{\"gigaFlexInlineSslApps\":[]}")}
	m := InlineSslAppListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestInlineSslAppListResource_List_NilClient exercises InlineSslAppListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInlineSslAppListResource_List_NilClient(t *testing.T) {
	r := &InlineSslAppListResource{}
	m := InlineSslAppListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestInlineSslAppListResource_List_BuildError exercises InlineSslAppListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestInlineSslAppListResource_List_BuildError(t *testing.T) {
	r := &InlineSslAppListResource{client: newMalformedBaseURLClient(t)}
	m := InlineSslAppListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineSslAppListResource_List_SendError exercises InlineSslAppListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestInlineSslAppListResource_List_SendError(t *testing.T) {
	r := &InlineSslAppListResource{client: newTransportErrorClient(t)}
	m := InlineSslAppListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestInlineSslAppListResource_List_InvalidJSON exercises InlineSslAppListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestInlineSslAppListResource_List_InvalidJSON(t *testing.T) {
	r := &InlineSslAppListResource{client: newMockClientStatus(t, 200, "{{")}
	m := InlineSslAppListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
