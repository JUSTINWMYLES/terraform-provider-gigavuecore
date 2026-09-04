package provider

import (
	"context"
	"testing"
)

// TestGsopListResource_List_Happy exercises GsopListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGsopListResource_List_Happy(t *testing.T) {
	r := &GsopListResource{client: newMockClientStatus(t, 200, "{\"gsops\":[]}")}
	m := GsopListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGsopListResource_List_NilClient exercises GsopListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsopListResource_List_NilClient(t *testing.T) {
	r := &GsopListResource{}
	m := GsopListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGsopListResource_List_BuildError exercises GsopListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGsopListResource_List_BuildError(t *testing.T) {
	r := &GsopListResource{client: newMalformedBaseURLClient(t)}
	m := GsopListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGsopListResource_List_SendError exercises GsopListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGsopListResource_List_SendError(t *testing.T) {
	r := &GsopListResource{client: newTransportErrorClient(t)}
	m := GsopListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGsopListResource_List_InvalidJSON exercises GsopListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGsopListResource_List_InvalidJSON(t *testing.T) {
	r := &GsopListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsopListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
