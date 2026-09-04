package provider

import (
	"context"
	"testing"
)

// TestElbListResource_List_Happy exercises ElbListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestElbListResource_List_Happy(t *testing.T) {
	r := &ElbListResource{client: newMockClientStatus(t, 200, "{\"elbs\":[]}")}
	m := ElbListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestElbListResource_List_NilClient exercises ElbListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestElbListResource_List_NilClient(t *testing.T) {
	r := &ElbListResource{}
	m := ElbListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestElbListResource_List_BuildError exercises ElbListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestElbListResource_List_BuildError(t *testing.T) {
	r := &ElbListResource{client: newMalformedBaseURLClient(t)}
	m := ElbListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestElbListResource_List_SendError exercises ElbListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestElbListResource_List_SendError(t *testing.T) {
	r := &ElbListResource{client: newTransportErrorClient(t)}
	m := ElbListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestElbListResource_List_InvalidJSON exercises ElbListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestElbListResource_List_InvalidJSON(t *testing.T) {
	r := &ElbListResource{client: newMockClientStatus(t, 200, "{{")}
	m := ElbListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
