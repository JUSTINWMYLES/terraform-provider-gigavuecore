package provider

import (
	"context"
	"testing"
)

// TestGigastreamListResource_List_Happy exercises GigastreamListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGigastreamListResource_List_Happy(t *testing.T) {
	r := &GigastreamListResource{client: newMockClientStatus(t, 200, "{\"gigastreams\":[]}")}
	m := GigastreamListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGigastreamListResource_List_NilClient exercises GigastreamListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigastreamListResource_List_NilClient(t *testing.T) {
	r := &GigastreamListResource{}
	m := GigastreamListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGigastreamListResource_List_BuildError exercises GigastreamListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGigastreamListResource_List_BuildError(t *testing.T) {
	r := &GigastreamListResource{client: newMalformedBaseURLClient(t)}
	m := GigastreamListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGigastreamListResource_List_SendError exercises GigastreamListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGigastreamListResource_List_SendError(t *testing.T) {
	r := &GigastreamListResource{client: newTransportErrorClient(t)}
	m := GigastreamListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGigastreamListResource_List_InvalidJSON exercises GigastreamListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGigastreamListResource_List_InvalidJSON(t *testing.T) {
	r := &GigastreamListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GigastreamListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
